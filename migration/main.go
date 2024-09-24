package main

import (
	"log"
	"migration/migrations"
	"migration/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type Migration struct {
	migrationID string
	description string
	Up          func(client *mongo.Client) error
}

var all_migrations = []Migration{
	{
		migrationID: "20241921",
		description: "Cars Data migration",
		Up:          migrations.MigrateCars,
	},
}

func main() {
	migration_service := services.NewMigrationService()
	for _, migration := range all_migrations {
		log.Printf("Processing migration: %s", migration.migrationID)
		migration_service.Up(migration.migrationID, migration.description, migration.Up)
	}
	log.Println("Migrations Completed Successfully!")
}
