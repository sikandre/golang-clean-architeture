package router

import (
	"cleanArch/internal/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RegisterRoutes(server *echo.Echo, appController controller.AppController) {
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	server.GET("/users", func(context echo.Context) error { return appController.User.GetUsers(context) })
}
