package repos

import (
	"libs/base"
	"libs/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type CarRepository struct {
	*base.Repository[models.CarModel]
}

func NewCarRepository(collection *mongo.Collection) *CarRepository {
	return &CarRepository{
		Repository: base.NewRepository[models.CarModel](collection),
	}
}
