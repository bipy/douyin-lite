package main

import (
	"douyin-lite/pkg/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	routes.PublicRoutes(app)
	routes.GeneralRoutes(app)

	app.Logger.Fatal(app.Start(":80"))
}
