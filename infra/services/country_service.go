package services

import (
	"infra/models"
	"infra/repos"
	"libs/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type CountryService struct {
	repository *repos.CountryRepository
}

func NewCountryService(collection *mongo.Collection) *CountryService {
	return &CountryService{
		repository: repos.NewCountryRepository(collection),
	}
}

func (s *CountryService) FindOneCountryByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := models.NewCountryModel()
	resp, err := s.repository.FindById(id)
	if err != nil {
		return utils.HandleMongoError(ctx, err, id)
	}

	if err := resp.Decode(user); err != nil {
		return utils.HandleMongoError(ctx, err, id)
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}
