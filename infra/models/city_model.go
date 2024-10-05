package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CityModel struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	StateId     string             `json:"stateId" bson:"stateId"`
	State       string             `json:"state" bson:"state"`
	CountryId   string             `json:"countryId" bson:"countryId"`
	Country     string             `json:"country" bson:"country"`
	CountryFlag string             `json:"countryFlag" bson:"countryFlag"`
}

func NewCityModel() *CityModel {
	return &CityModel{}
}
