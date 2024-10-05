package migrations

import (
	"api/services"
	"encoding/csv"
	"fmt"
	"libs/base"
	"libs/constants"
	"libs/database"
	"libs/models"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
)

var cars_collection_name = "cars"

func MigrateCars(client *mongo.Client) error {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %s", err)
	}
	filePath := filepath.Join(wd, "../datasets", "cars_ds_final.csv")
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()
	collection := database.GetCollection(client, constants.API_DB_NAME, cars_collection_name)

	car_service := services.NewCarService(collection)

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read csv: %w", err)
	}

	for i, record := range records[1:] {

		cylinders, _ := strconv.Atoi(record[6])

		gears, _ := strconv.Atoi(record[25])
		seatingCapacity, _ := strconv.Atoi(record[45])

		car := &models.CarModel{
			Make:            record[1],
			Variant:         record[3],
			Model:           record[2],
			FuelType:        record[14],
			Cylinder:        int8(cylinders),
			Mileage:         record[20],
			DriveTrain:      record[8],
			Displacement:    record[5],
			TankCapacity:    int32(51),
			BodyType:        record[18],
			Gears:           int8(gears),
			Power:           record[39],
			Torque:          record[40],
			Transmission:    record[47],
			EmissionNorm:    record[10],
			FrontBrake:      record[27],
			BackBrake:       record[28],
			SeatingCapacity: int8(seatingCapacity),
			BaseModel:       &base.BaseModel{},
		}
		car.BaseModel.InitiliseDefaultValue()

		car_service.CreateCar(car)
		fmt.Println("Processing car", i)
	}
	return nil
}
