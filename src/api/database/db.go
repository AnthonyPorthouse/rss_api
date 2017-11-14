package database

import (
	"api/config"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func GetDB() *sqlx.DB {
	if db == nil {
		initDB()
	}

	return db
}

func initDB() {
	config := config.GetConfig()

	c := mysql.Config{
		User:   config.DbUsername,
		Passwd: config.DbPassword,
		DBName: config.DbName,
		Net:    "tcp",
		Addr: fmt.Sprintf(
			"%s:%s",
			config.DbHost,
			config.DbPort,
		),
		MultiStatements: true,
		ParseTime:       true,
		Strict:          true,
	}

	attempt := 0
	var err error

	for db == nil && attempt <= 10 {
		attempt += 1

		log.Println("Attempting to connect to DB:", attempt)

		db, err = sqlx.Connect("mysql", c.FormatDSN())

		if err != nil && attempt == 10 {
			log.Panicln("Failed to connect to DB:", err)
		}

		if err != nil {
			log.Println("Sleeping 10 seconds before retrying")
			time.Sleep(10 * time.Second)
		}

		if db != nil {
			log.Println("Connected to DB")
		}
	}
}
