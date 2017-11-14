package main

import (
	"api/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())

	controllers.MapBaseRoutes(e)
	controllers.MapFeedRoutes(e)

	e.Logger.Fatal(e.Start(":1234"))
}
