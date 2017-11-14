package controllers

import (
	"api/domain"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func MapFeedRoutes(e *echo.Echo) {
	group := e.Group("/feeds")

	group.GET("", getAllFeeds).Name = "feeds.index"
}

func getAllFeeds(c echo.Context) error {
	feeds := domain.GetAllFeeds()

	fmt.Printf("%+v\n", feeds)

	return c.JSON(http.StatusOK, feeds)
}
