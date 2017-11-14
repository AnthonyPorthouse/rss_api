package domain

import (
	"api/database"
	"time"
)

type Feed struct {
	ID        string    `db:"id" json:"id"`
	URL       string    `db:"url" json:"url"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Feeds []*Feed

func GetAllFeeds() *Feeds {
	db := database.GetDB()

	var feeds *Feeds
	db.Select(&feeds, "SELECT * FROM feeds")

	return feeds
}
