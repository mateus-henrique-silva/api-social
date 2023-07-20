package entity

import "time"

type Youtube struct {
	VideoID    string    `bson:"videoID" json:"videoID,omitempty"`
	Status     string    `bson:"status" json:"status,omitempty"`
	AuthorName string    `bson:"authorName" json:"authorName,omitempty"`
	Title      string    `bson:"title" json:"title,omitempty"`
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt,omitempty"`
	EditedAt   time.Time `bson:"editedAt" json:"editedAt,omitempty"`
	Summary    string    `bson:"summary" json:"summary,omitempty"`
	CategoryID string    `bson:"categoryID" json:"categoryID,omitempty"`
}
