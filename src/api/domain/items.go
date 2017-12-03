package domain

import (
	"api/database"
	"github.com/guregu/null"
	"time"
)

type Item struct {
	ID     string `db:"id" json:"id"`
	FeedID string `db:"feed_id" json:"feed_id"`

	Link    string `db:"link" json:"link"`
	Content string `db:"content" json:"content"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt null.Time `db:"updated_at" json:"updated_at"`
}
