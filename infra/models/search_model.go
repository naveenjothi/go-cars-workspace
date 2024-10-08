package models

import (
	"libs/models"
)

type CitySearchInputModel struct {
	Query     string
	CountryId string
	Paging    *models.PagingModel
}

func NewCitySearchInputModel() *CitySearchInputModel {
	return &CitySearchInputModel{
		Paging: models.NewPagingModel(),
	}
}
