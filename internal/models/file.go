package models

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID uuid.UUID

	WorkspaceID uuid.UUID

	UploaderID uuid.UUID

	Name string

	StorageKey string

	Size int64

	MimeType string

	CreatedAt time.Time
}
