package repos

import (
	"libs/base"
	"migration/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type MigrationRepo struct {
	*base.Repository[models.MigrationModel]
}

func NewMigrationRepo(collection *mongo.Collection) *MigrationRepo {
	return &MigrationRepo{
		Repository: base.NewRepository[models.MigrationModel](collection),
	}
}
