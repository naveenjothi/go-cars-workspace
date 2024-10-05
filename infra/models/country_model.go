package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CountryModel struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Demonym    string             `json:"demonym" bson:"demonym"`
	Alpha2Code string             `json:"alpha2Code" bson:"alpha2Code"`
	Alpha3Code string             `json:"alpha3Code" bson:"alpha3Code"`
	Flag       string             `json:"flag" bson:"flag"`
}

func NewCountryModel() *CountryModel {
	return &CountryModel{}
}
