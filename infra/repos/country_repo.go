package repos

import (
	"infra/models"
	"libs/base"

	"go.mongodb.org/mongo-driver/mongo"
)

type CountryRepository struct {
	*base.Repository[models.CountryModel]
}

func NewCountryRepository(collection *mongo.Collection) *CountryRepository {
	return &CountryRepository{
		Repository: base.NewRepository[models.CountryModel](collection),
	}
}
