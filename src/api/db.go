package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func getDB() *sqlx.DB {
	return db
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
		MultiStatements: true,
	}

	attempt := 0
	var err error

	for db == nil && attempt <= 10 {
		attempt += 1

		log.Println("Attempting to connect to DB:", attempt)

		db, err = sqlx.Connect("mysql", config.FormatDSN())

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

	migrateDB()
}
