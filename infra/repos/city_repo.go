package repos

import (
	"infra/models"
	"libs/base"

	"go.mongodb.org/mongo-driver/mongo"
)

type CityRepository struct {
	*base.Repository[models.CityModel]
}

func NewCityRepository(collection *mongo.Collection) *CityRepository {
	return &CityRepository{
		Repository: base.NewRepository[models.CityModel](collection),
	}
}
