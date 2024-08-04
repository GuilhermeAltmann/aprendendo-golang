package routers

import (
	"testezinho/src/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
	root := e.Group("api/v1")
	controllers.AddCarsRoutes(root)
}
