package repos

import (
	"libs/base"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	*base.Repository
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		Repository: base.NewRepository(collection),
	}
}
