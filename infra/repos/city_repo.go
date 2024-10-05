package repos

import (
	"libs/base"

	"go.mongodb.org/mongo-driver/mongo"
)

type CityRepository struct {
	*base.Repository
}

func NewCityRepository(collection *mongo.Collection) *CityRepository {
	return &CityRepository{
		Repository: base.NewRepository(collection),
	}
}
