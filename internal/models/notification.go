package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID uuid.UUID

	UserID uuid.UUID

	Type string

	ReferenceID uuid.UUID

	Read bool

	CreatedAt time.Time
}
