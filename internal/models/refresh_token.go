package models

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID uuid.UUID

	UserID uuid.UUID

	TokenHash string

	ExpiresAt time.Time

	CreatedAt time.Time
	Revoked   bool
}
