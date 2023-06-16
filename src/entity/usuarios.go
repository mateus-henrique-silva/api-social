package entity

import (
	"time"

	"github.com/google/uuid"
)

type Usuario struct {
	Id        uuid.UUID `bson:"id"`
	Name      string    `bson:"name"`
	CreatedAt time.Time `bson:"created"`
}
