package main

import (
	"log"
	"migration/migrations"
	"migration/services"
)

type Migration struct {
	migrationID string
	description string
	Up          func() error
}

var all_migrations = []Migration{
	{
		migrationID: "20240921",
		description: "20240921",
		Up:          migrations.MigrateUsers,
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
