package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate"
	migrateMysql "github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
)

func MigrateDB() {
	driver, err := migrateMysql.WithInstance(GetDB().DB, &migrateMysql.Config{})
	if err != nil {
		log.Panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///go/src/api/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Migrating DB")

	err = m.Up()

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Migrated DB")
}
