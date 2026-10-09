package models

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	ID uuid.UUID `db:"id"`

	WorkspaceID uuid.UUID `db:"workspace_id"`

	Name string `db:"name"`

	Description *string `db:"description"`

	Type ChannelType `db:"type"`

	CreatedBy uuid.UUID `db:"created_by"`

	Position int32 `db:"position"`

	IsPrivate bool `db:"is_private"`

	DeletedAt *time.Time `db:"deleted_at"`

	CreatedAt time.Time `db:"created_at"`

	UpdatedAt time.Time `db:"updated_at"`
}
