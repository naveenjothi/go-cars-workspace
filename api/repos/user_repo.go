package repos

import (
	"libs/base"
	"libs/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	*base.Repository[models.UserModel]
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		Repository: base.NewRepository[models.UserModel](collection),
	}
}
