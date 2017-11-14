package domain

import (
	"api/database"
	"log"
	"time"
)

type Feed struct {
	ID        string    `db:"id" json:"id"`
	URL       string    `db:"url" json:"url"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Feeds []Feed

func GetAllFeeds(limit int, after string) Feeds {
	db := database.GetDB()

	feeds := make(Feeds, 0)
	err := db.Select(&feeds, `
		SELECT *
		FROM feeds
		WHERE id > ?

		ORDER BY created_at DESC
		LIMIT ?
	`, after, limit)

	if err != nil {
		log.Println(err)
	}

	return feeds
}
