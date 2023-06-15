package entity

import (
	"time"

	"github.com/google/uuid"
)

type Usuario struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created"`
}
