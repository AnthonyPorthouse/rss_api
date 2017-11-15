package controllers

import (
	"api/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func MapFeedRoutes(e *echo.Echo) {
	feeds := e.Group("/feeds")

	feeds.GET("", getAllFeeds).Name = "feeds.index"
	feeds.POST("", addFeed).Name = "feeds.create"
	feeds.GET("/:id", getFeed).Name = "feeds.read"
}

func getFeed(c echo.Context) error {
	feed := domain.GetFeed(c.Param("id"))

	if feed == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, feed)
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

func addFeed(c echo.Context) (err error) {
	var feed *domain.Feed

	if err = c.Bind(feed); err != nil {
		return
	}

	feed = domain.CreateFeed(*feed)

	return c.JSON(http.StatusCreated, feed)
}
