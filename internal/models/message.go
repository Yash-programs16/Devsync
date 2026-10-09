package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID uuid.UUID

	ChannelID uuid.UUID

	AuthorID uuid.UUID

	Content string

	Edited bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
