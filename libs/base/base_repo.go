package base

import (
	"context"
	"fmt"
	"libs/utils"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository[T any] struct {
	collection *mongo.Collection
}

func NewRepository[T any](collection *mongo.Collection) *Repository[T] {
	return &Repository[T]{collection}
}

func (r *Repository[T]) FindById(id string) (*T, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	result := r.collection.FindOne(context.Background(), filter)

	var res *T

	if err := result.Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Repository[T]) FindOne(filter interface{}) *mongo.SingleResult {
	return r.collection.FindOne(context.Background(), filter)
}

func (r *Repository[T]) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.Background(), document)
}

func (r *Repository[T]) UpdateOne(id string, document interface{}) (*interface{}, error) {
	// Convert the ID string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		return nil, err
	}

	// Construct the filter to find the document
	filter := bson.M{"_id": objectID}

	// Construct the update document using $set to update specific fields
	update := bson.M{"$set": document}

	// Define the result variable to store the updated document
	result := new(interface{})

	// Find the document and apply the update with options to return the updated document
	err = r.collection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After), // Return the updated document
	).Decode(result)

	if err != nil {
		log.Printf("Error updating document: %v", err)
		return nil, err
	}

	return result, nil
}

func (r *Repository[T]) DeleteOne(filter interface{}) (*mongo.DeleteResult, error) {
	return r.collection.DeleteOne(context.Background(), filter)
}

func (r *Repository[T]) AtlasSearch(filter bson.D, page, pageSize int) (*BaseDto[T], error) {
	skip := int64(page * pageSize)
	limit := int64(pageSize)

	cursor, err := r.collection.Find(context.TODO(), filter, options.Find().SetSkip(skip).SetLimit(limit))
	if err != nil {
		return nil, fmt.Errorf("error finding documents: %v", err)
	}
	defer cursor.Close(context.TODO())

	var items []T
	if err = cursor.All(context.TODO(), &items); err != nil {
		return nil, fmt.Errorf("error retrieving documents: %v", err)
	}

	matchStage := bson.D{{Key: "$match", Value: filter}}
	countCursor, err := r.collection.Aggregate(context.TODO(), mongo.Pipeline{
		matchStage,
		utils.CountStage,
	})
	if err != nil {
		return nil, fmt.Errorf("error aggregating documents: %v", err)
	}
	defer countCursor.Close(context.TODO())

	var countResult []bson.M
	if err = countCursor.All(context.TODO(), &countResult); err != nil {
		return nil, fmt.Errorf("error retrieving count: %v", err)
	}

	var count int32
	if len(countResult) > 0 {
		if totalDocs, ok := countResult[0]["total_documents"].(int32); ok {
			count = totalDocs
		} else {
			return nil, fmt.Errorf("error asserting total_documents field")
		}
	}

	return &BaseDto[T]{
		Items: items,
		Count: count,
	}, nil
}
