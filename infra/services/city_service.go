package services

import (
	"infra/models"
	"infra/repos"
	"libs/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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
	resp, err := s.repository.FindById(id)
	if err != nil {
		return utils.HandleMongoError(ctx, err, id)
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (s *CityService) SearchCities(ctx *fiber.Ctx) error {
	dto := models.NewCitySearchInputModel()
	if err := utils.GetBodyPayload(ctx, dto); err != nil {
		return err
	}
	filter := bson.D{}
	resp, err := s.repository.AtlasSearch(filter, int(dto.Paging.Offset), int(dto.Paging.Limit))
	if err != nil {
		return utils.HandleMongoError(ctx, err, "")
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
