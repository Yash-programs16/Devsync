package models

import (
	"time"

	"github.com/google/uuid"
)

type Reaction struct {
	ID uuid.UUID

	MessageID uuid.UUID

	UserID uuid.UUID

	Emoji string

	CreatedAt time.Time
}
