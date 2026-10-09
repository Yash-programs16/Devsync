package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	UserID uuid.UUID

	Username string

	IssuedAt time.Time

	ExpiredAt time.Time
}
