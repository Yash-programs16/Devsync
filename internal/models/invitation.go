package models

import (
	"time"

	"github.com/google/uuid"
)

type Invitation struct {
	ID uuid.UUID

	WorkspaceID uuid.UUID

	InvitedBy uuid.UUID

	Email string

	Token string

	ExpiresAt time.Time

	Accepted bool

	CreatedAt time.Time
}
