package models

import (
	"time"

	"github.com/google/uuid"
)

type CodeSnippet struct {
	ID uuid.UUID

	MessageID uuid.UUID

	Language string

	Code string

	CreatedAt time.Time
}
