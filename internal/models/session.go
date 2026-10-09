package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID uuid.UUID

	UserID uuid.UUID

	RefreshTokenHash string

	UserAgent string

	IPAddress string

	ExpiresAt time.Time

	Revoked bool

	CreatedAt time.Time
}
