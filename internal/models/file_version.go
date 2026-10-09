package models

import (
	"time"

	"github.com/google/uuid"
)

type FileVersion struct {
	ID uuid.UUID

	FileID uuid.UUID

	Version int

	StorageKey string

	UploadedBy uuid.UUID

	CreatedAt time.Time
}
