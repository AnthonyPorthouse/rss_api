package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	initDB()

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/status", func(c echo.Context) error {

		db := getDB()

		err := db.Ping()
		if err != nil {
			return c.String(http.StatusInternalServerError, "Database is Unavailable")
		}

		return c.String(http.StatusOK, "Everything's OK!")
	})
	e.Logger.Fatal(e.Start(":1234"))
}
