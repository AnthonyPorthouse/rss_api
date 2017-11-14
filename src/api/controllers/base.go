package controllers

import (
	"api/database"
	"net/http"

	"github.com/labstack/echo"
)

func MapBaseRoutes(e *echo.Echo) {
	e.GET("/", getHome).Name = "home"
	e.GET("/status", getStatus).Name = "status"
}

func getHome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getStatus(c echo.Context) error {
	db := database.GetDB()

	err := db.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Database is Unavailable")
	}

	return c.String(http.StatusOK, "Everything's OK!")
}
