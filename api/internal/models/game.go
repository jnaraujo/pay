package models

import (
	"time"

	"github.com/google/uuid"
)

type Game struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	OwnerId   uuid.UUID `json:"owner_id"`
}
