package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID  `json:"id" validate:"uuid"`
	Name      string     `json:"name" validate:"required,lte=100"`
	CreatedAt time.Time  `json:"created_at"`
	GameId    *uuid.UUID `json:"game_id"`
	Balance   int        `json:"balance"`
}
