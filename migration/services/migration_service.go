package services

import (
	"fmt"
	"libs/database"
	"log"
	"migration/repos"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MigrationService struct {
	repository *repos.MigrationRepo
}

func NewMigrationService() *MigrationService {
	var migration_collection_name = "migrations"
	client, err := database.InitializeMongoClient("API")
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB client: %s", err)
	}
	collection := database.GetCollection(client, migration_collection_name)
	return &MigrationService{
		repository: repos.NewMigrationRepo(collection),
	}
}

func (s *MigrationService) Up(migrationID, description string, cb func() error) {
	var result bson.M
	err := s.repository.FindOne(bson.M{"version": migrationID}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Apply migration logic here
		fmt.Println("Applying migration", migrationID)
		if cbErr := cb(); cbErr != nil {
			log.Fatal(cbErr)
		}
		_, err := s.repository.InsertOne(bson.M{
			"version":     migrationID,
			"description": description,
			"applied_at":  time.Now(),
			"status":      "Success",
		})
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Migration", migrationID, "has already been applied.")
	}
}
