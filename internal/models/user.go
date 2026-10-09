package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID

	Username string
	Email    string

	PasswordHash string

	DisplayName string

	AvatarURL *string

	Status UserStatus

	CreatedAt     time.Time
	UpdatedAt     time.Time
	EmailVerified bool
}
