package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func MapFeedRoutes(e *echo.Echo) {
	group := e.Group("/feeds")

	group.GET("/", getAllFeeds).Name = "feeds.index"
}

func getAllFeeds(c echo.Context) error {
	return c.String(http.StatusOK, "")
}
