package server

import (
	"fmt"
	"log"

	"cleanArch/config"
	"cleanArch/internal/controller"
	"cleanArch/internal/controller/router"

	"github.com/labstack/echo"
)

func NewHttpServer(appController controller.AppController) {
	e := echo.New()
	router.RegisterRoutes(e, appController)

	fmt.Println("Server listen at http://localhost" + ":" + config.Configuration.Server.Address)

	if err := e.Start(":" + config.Configuration.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
