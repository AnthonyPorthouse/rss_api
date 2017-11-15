package domain

import (
	"api/database"
	"errors"
	"fmt"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

type Feed struct {
	ID        string    `db:"id" json:"id"`
	URL       string    `db:"url" json:"url"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt null.Time `db:"updated_at" json:"updated_at"`
}

type Feeds []Feed

func GetFeed(id string) (Feed, error) {
	db := database.GetDB()

	var feed Feed
	db.Get(&feed, `
		SELECT *
		FROM feeds
		WHERE id = ?
	`, id)

	if feed.ID == "" {
		return feed, errors.New(
			fmt.Sprintf("No feed with id %s found", id),
		)
	}

	return feed, nil
}

func GetAllFeeds(limit int, after string) (Feeds, error) {
	db := database.GetDB()

	feeds := make(Feeds, 0, limit)
	err := db.Select(&feeds, `
		SELECT *
		FROM feeds
		WHERE id > ?

		ORDER BY created_at DESC
		LIMIT ?
	`, after, limit)

	if err != nil {
		return feeds, err
	}

	return feeds, nil
}

func CreateFeed(feed Feed) (Feed, error) {
	db := database.GetDB()

	feed.ID = uuid.NewV4().String()

	db.MustExec(`
		INSERT INTO feeds (id, url, created_at)
		VALUES (
			?,
			?,
			NOW()
		)
	`, feed.ID, feed.URL)

	return GetFeed(feed.ID)
}
