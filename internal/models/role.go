package models

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID uuid.UUID

	WorkspaceID uuid.UUID

	Name string

	CreatedAt time.Time
}
