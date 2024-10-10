package services

import (
	"api/repos"
	"libs/models"
	"libs/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	repository *repos.UserRepository
}

func NewUserService(collection *mongo.Collection) *UserService {
	return &UserService{
		repository: repos.NewUserRepository(collection),
	}
}

func (s *UserService) CreateUser(ctx *fiber.Ctx) error {
	dto := models.NewUserModel()
	if err := utils.GetBodyPayload(ctx, dto); err != nil {
		return err
	}
	dto.InitiliseDefaultValue()
	insertResult, err := s.repository.InsertOne(dto)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Profile name already exists"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unable to create user"})
	}
	dto.Id = insertResult.InsertedID.(primitive.ObjectID)
	return ctx.Status(fiber.StatusCreated).JSON(dto)
}

func (s *UserService) FindOneUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	resp, err := s.repository.FindById(id)
	if err != nil {
		return utils.HandleMongoError(ctx, err, id)
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (s *UserService) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	dto := models.NewUserModel()
	if err := utils.GetBodyPayload(ctx, dto); err != nil {
		return err
	}
	exists, err := s.repository.FindById(id)
	if err != nil {
		return utils.HandleMongoError(ctx, err, id)
	}
	changes := utils.IdentifyChanges(exists, dto)
	if len(changes) == 0 {
		return ctx.Status(fiber.StatusOK).JSON(exists)
	}
	changes["updatedAt"] = time.Now()
	updatedData, err := s.repository.UpdateOne(id, changes)
	if err != nil {
		return utils.HandleMongoError(ctx, err, id)
	}
	return ctx.Status(fiber.StatusOK).JSON(updatedData)
}
