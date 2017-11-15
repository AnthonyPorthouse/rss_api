package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func MapBaseRoutes(e *echo.Echo) {
	e.GET("/", getHome).Name = "home"
	e.GET("/status", getStatus).Name = "status"
}

func getHome(c echo.Context) error {
	data := struct {
		Echo *echo.Echo
	}{
		c.Echo(),
	}

	return c.Render(http.StatusOK, "home.html", data)
}

func getStatus(c echo.Context) error {
	return c.String(http.StatusOK, "Everything's OK!")
}
