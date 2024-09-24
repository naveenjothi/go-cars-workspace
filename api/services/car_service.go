package services

import (
	"api/repos"
	"libs/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type CarService struct {
	repository *repos.CarRepository
}

func NewCarService(collection *mongo.Collection) *CarService {
	return &CarService{
		repository: repos.NewCarRepository(collection),
	}
}

func (s *CarService) CreateCar(dto *models.CarModel) error {
	_, err := s.repository.InsertOne(dto)
	if err != nil {
		return err
	}
	return err
}
