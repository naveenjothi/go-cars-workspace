package services

import (
	"infra/models"
	"infra/repos"
	"libs/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type CityService struct {
	repository *repos.CityRepository
}

func NewCityService(collection *mongo.Collection) *CityService {
	return &CityService{
		repository: repos.NewCityRepository(collection),
	}
}

func (s *CityService) FindOneCityByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := models.NewCityModel()
	resp, err := s.repository.FindById(id)
	if err != nil {
		return utils.HandleMongoError(ctx, err, id)
	}

	if err := resp.Decode(user); err != nil {
		return utils.HandleMongoError(ctx, err, id)
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}
