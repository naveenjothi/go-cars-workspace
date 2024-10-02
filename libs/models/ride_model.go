package models

import (
	"libs/base"
	"time"
)

type RideModel struct {
	From            Address         `json:"from" bson:"from"`
	To              Address         `json:"to" bson:"to"`
	AllowedPeoples  int8            `json:"allowedPeoples" bson:"allowedPeoples"`
	StartsAt        time.Time       `json:"startsAt" bson:"startsAt"`
	RiderId         string          `json:"riderId" bson:"riderId"`
	Stops           []RideStop      `json:"stops" bson:"stops"`
	CarId           string          `json:"carId" bson:"carId"`
	RiderPreference RiderPreference `json:"riderPreference" bson:"riderPreference"`
	*base.BaseModel `bson:",inline"`
}

type RiderPreference struct {
	WomenPassengersOnly bool `json:"womenPassengersOnly" bson:"womenPassengersOnly"`
	NoSmoking           bool `json:"noSmoking" bson:"noSmoking"`
	NoPets              bool `json:"noPets" bson:"noPets"`
}

type RideStop struct {
	Location    Address   `json:"location" bson:"location"`
	Description string    `json:"description" bson:"description"`
	StopTime    time.Time `json:"stopTime" bson:"stopTime"`
}

func NewRideModel() *RideModel {
	return &RideModel{
		BaseModel: base.NewBaseModel(),
	}
}
