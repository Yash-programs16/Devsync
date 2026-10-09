package models

import (
	"time"

	"github.com/google/uuid"
)

type Workspace struct {
	ID uuid.UUID `json:"id"`

	Name string `json:"name"`
	Slug string `json:"slug"`

	Description *string `json:"description,omitempty"`

	OwnerID uuid.UUID `json:"owner_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
