package services

import (
	"testezinho/src/models"
	"testezinho/src/shared"

	"gorm.io/gorm"
)

func GetAllCars() ([]models.Car, error) {
	var cars []models.Car
	result := shared.DB.Find(&cars)

	return cars, result.Error
}

func CreateNewCar(car models.Car) (models.Car, *gorm.DB) {
	result := shared.DB.Create(&car)
	return car, result
}
