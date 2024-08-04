package main

import (
	"testezinho/src/routers"
	"testezinho/src/shared"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	shared.CreateDatabaseConnection()
	routers.SetupRouter(e)
	e.Logger.Fatal(e.Start(":8080"))
}
