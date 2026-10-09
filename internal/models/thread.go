package models

import (
	"time"

	"github.com/google/uuid"
)

type Thread struct {
	ID uuid.UUID

	MessageID uuid.UUID

	AuthorID uuid.UUID

	Content string

	CreatedAt time.Time
	UpdatedAt time.Time
}
