package repos

import (
	"libs/base"

	"go.mongodb.org/mongo-driver/mongo"
)

type CarRepository struct {
	*base.Repository
}

func NewCarRepository(collection *mongo.Collection) *CarRepository {
	return &CarRepository{
		Repository: base.NewRepository(collection),
	}
}
