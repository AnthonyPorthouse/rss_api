package controllers

import (
	"api/domain"
	"fmt"
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

	feeds := domain.GetAllFeeds(limit, after)

	fmt.Printf("%+v\n", feeds)

	return c.JSON(http.StatusOK, feeds)
}
