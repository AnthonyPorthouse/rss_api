package main

import (
	"github.com/labstack/echo"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
	"github.com/labstack/gommon/log"
)

var db *sqlx.DB

func main() {

	initDB()

	e := echo.New()
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
		User: os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
		Net: "tcp",
		Addr: fmt.Sprintf(
			"%s:%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
		),
	}

	var err error
	db, err = sqlx.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
}
