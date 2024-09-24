package models

import "libs/base"

type CarModel struct {
	Make         string `json:"make" bson:"make"`
	Variant      string `json:"variant" bson:"variant"`
	Model        string `json:"model" bson:"model"`
	FuelType     string `json:"fuelType" bson:"fuelType"`
	Transmission string `json:"transmission" bson:"transmission"`
	Mileage      string `json:"mileage" bson:"mileage"`

	// Additional car specifications
	DriveTrain      string `json:"driveTrain" bson:"driveTrain"`     // E.g., FWD, RWD, AWD
	Displacement    string `json:"displacement" bson:"displacement"` // E.g., 1998 cc
	Cylinder        int8   `json:"cylinder" bson:"cylinder"`         // Number of engine cylinders
	EmissionNorm    string `json:"emissionNorm" bson:"emissionNorm"` // E.g., BS6, Euro 6
	TankCapacity    int32  `json:"tankCapacity" bson:"tankCapacity"` // Fuel tank capacity in liters
	BodyType        string `json:"bodyType" bson:"bodyType"`         // E.g., Sedan, SUV, Hatchback
	Gears           int8   `json:"gears" bson:"gears"`               // Number of gears (e.g., 5-speed)
	FrontBrake      string `json:"frontBrake" bson:"frontBrake"`
	BackBrake       string `json:"backBrake" bson:"backBrake"`             // E.g., Disc, Drum
	Power           string `json:"power" bson:"power"`                     // E.g., 100 kW or 150 HP
	Torque          string `json:"torque" bson:"torque"`                   // E.g., 300 Nm
	SeatingCapacity int8   `json:"seatingCapacity" bson:"seatingCapacity"` // Number of seats (e.g., 5-seater)
	*base.BaseModel `bson:",inline"`
}

func NewCarModel() *CarModel {
	return &CarModel{
		BaseModel: base.NewBaseModel(),
	}
}
