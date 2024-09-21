package repos

import (
	"libs/base"

	"go.mongodb.org/mongo-driver/mongo"
)

type MigrationRepo struct {
	*base.Repository
}

func NewMigrationRepo(collection *mongo.Collection) *MigrationRepo {
	return &MigrationRepo{
		Repository: base.NewRepository(collection),
	}
}
