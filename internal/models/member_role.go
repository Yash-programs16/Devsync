package models

import "github.com/google/uuid"

type MemberRole struct {
	ID uuid.UUID

	MemberID uuid.UUID

	RoleID uuid.UUID
}
