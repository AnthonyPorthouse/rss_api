package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/middleware"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var db *sqlx.DB

func main() {

	initDB()

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/status", func(c echo.Context) error {

		err := db.Ping()
		if err != nil {
			return c.String(http.StatusInternalServerError, "Database is Unavailable")
		}

		return c.String(http.StatusOK, "Everything's OK!")
	})
	e.Logger.Fatal(e.Start(":1234"))
}

func initDB() {
	config := mysql.Config{
		User:   Config.DbUsername,
		Passwd: Config.DbPassword,
		DBName: Config.DbName,
		Net:    "tcp",
		Addr: fmt.Sprintf(
			"%s:%s",
			Config.DbHost,
			Config.DbPort,
		),
	}

	var err error
	db, err = sqlx.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
}
