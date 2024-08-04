package controllers

import (
	"net/http"
	"testezinho/src/services"

	"github.com/labstack/echo/v4"
)

type Car struct {
	Name  string
	Price float64
}

var cars []Car

func createCars() {
	cars = append(cars, Car{Name: "Ferrari", Price: 100})
	cars = append(cars, Car{Name: "Mercedes", Price: 1000})
	cars = append(cars, Car{Name: "Porche", Price: 1020})
}

func createCar(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	return c.JSON(200, cars)
}

func getCars(c echo.Context) error {
	result, err := services.GetAllCars()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Erro!")
	}

	return c.JSON(http.StatusOK, result)
}

func AddCarsRoutes(e *echo.Group) {
	root := "/cars"
	e.GET(root, getCars)
	e.POST(root, createCar)
}
