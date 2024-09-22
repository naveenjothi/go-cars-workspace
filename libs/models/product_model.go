package models

import (
	"libs/base"
	"time"
)

type CarModel struct {
	Make                string    `json:"make" bson:"make"`
	Model               string    `json:"model" bson:"model"`
	Year                int       `json:"year" bson:"year"`
	Category            string    `json:"category" bson:"category"`
	FuelType            string    `json:"fuelType" bson:"fuelType"`
	Transmission        string    `json:"transmission" bson:"transmission"`
	Mileage             int       `json:"mileage" bson:"mileage"`
	Seats               int       `json:"seats" bson:"seats"`
	Color               string    `json:"color" bson:"color"`
	DailyRentalRate     float64   `json:"dailyRentalRate" bson:"dailyRentalRate"`
	WeeklyRentalRate    float64   `json:"weeklyRentalRate,omitempty" bson:"weeklyRentalRate,omitempty"`
	Availability        bool      `json:"availability" bson:"availability"`
	Location            string    `json:"location" bson:"location"`
	OwnerID             string    `json:"ownerId" bson:"ownerId"`
	IsInsured           bool      `json:"isInsured" bson:"isInsured"`
	InsuranceExpiryDate time.Time `json:"insuranceExpiryDate,omitempty" bson:"insuranceExpiryDate,omitempty"`
	IsMaintained        bool      `json:"isMaintained" bson:"isMaintained"`
	MaintenanceDetails  string    `json:"maintenanceDetails,omitempty" bson:"maintenanceDetails,omitempty"`
	LicensePlate        string    `json:"licensePlate" bson:"licensePlate"`
	VIN                 string    `json:"vin" bson:"vin"`
	RentalConditions    string    `json:"rentalConditions,omitempty" bson:"rentalConditions,omitempty"`
	Images              []string  `json:"images" bson:"images"`
	*base.BaseModel     `bson:",inline"`
}
