package controllers

import (
	"api/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func MapFeedRoutes(e *echo.Echo) {
	group := e.Group("/feeds")

	group.GET("", getAllFeeds).Name = "feeds.index"
}

func getAllFeeds(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	after := c.QueryParam("after")

	if limit <= 0 {
		limit = 10
	}
	if limit > 50 {
		limit = 50
	}

	feeds := domain.GetAllFeeds(limit, after)

	return c.JSON(http.StatusOK, feeds)
}
