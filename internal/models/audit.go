package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID uuid.UUID

	WorkspaceID uuid.UUID

	UserID uuid.UUID

	Action string

	Metadata json.RawMessage

	CreatedAt time.Time
}
