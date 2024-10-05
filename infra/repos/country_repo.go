package repos

import (
	"libs/base"

	"go.mongodb.org/mongo-driver/mongo"
)

type CountryRepository struct {
	*base.Repository
}

func NewCountryRepository(collection *mongo.Collection) *CountryRepository {
	return &CountryRepository{
		Repository: base.NewRepository(collection),
	}
}
