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

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{collection}
}

func (r *Repository) FindById(id string) (*mongo.SingleResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	result := r.collection.FindOne(context.Background(), filter)

	return result, nil
}

func (r *Repository) FindOne(filter interface{}) *mongo.SingleResult {
	return r.collection.FindOne(context.Background(), filter)
}

func (r *Repository) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.Background(), document)
}

func (r *Repository) UpdateOne(id string, document interface{}) (*interface{}, error) {
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

func (r *Repository) DeleteOne(filter interface{}) (*mongo.DeleteResult, error) {
	return r.collection.DeleteOne(context.Background(), filter)
}

func (r *Repository) AtlasSearch(filter bson.D, page, pageSize int) (*BaseDto, error) {
	skip := int64(page * pageSize)
	limit := int64(pageSize)

	cursor, err := r.collection.Find(context.TODO(), filter, options.Find().SetSkip(skip).SetLimit(limit))
	if err != nil {
		return nil, fmt.Errorf("error finding documents: %v", err)
	}
	defer cursor.Close(context.TODO())

	var items []bson.D
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

	return &BaseDto{
		Items: items,
		Count: count,
	}, nil
}
