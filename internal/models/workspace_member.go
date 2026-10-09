package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkspaceMember struct {
	ID uuid.UUID

	UserID uuid.UUID

	WorkspaceID uuid.UUID

	JoinedAt time.Time
}
