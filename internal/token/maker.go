package token

import "github.com/google/uuid"

type Maker interface {
	CreateAccessToken(
		userID uuid.UUID,
		username string,
	) (string, error)

	CreateRefreshToken(
		userID uuid.UUID,
	) (string, error)

	VerifyToken(
		token string,
	) (*Payload, error)
}
