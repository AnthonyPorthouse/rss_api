package main

import (
	"api/controllers"
	"api/database"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	database.MigrateDB()

	e := echo.New()
	e.HideBanner = true
	e.Pre(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	e.Use(middleware.Logger())

	controllers.MapBaseRoutes(e)
	controllers.MapFeedRoutes(e)

	e.Logger.Fatal(e.Start(":1234"))
}
