package token

import (
	"encoding/hex"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	symmetricKey []byte
}

func NewPasetoMaker(key string) (*PasetoMaker, error) {

	decodedKey, err := hex.DecodeString(key)

	if err != nil {
		return nil, errors.New("invalid key format")
	}

	if len(decodedKey) != 32 {
		return nil, errors.New("key must be exactly 32 bytes")
	}

	return &PasetoMaker{
		symmetricKey: decodedKey,
	}, nil
}

func (maker *PasetoMaker) CreateAccessToken(
	userID uuid.UUID,
	username string,
) (string, error) {

	payload := Payload{
		UserID:    userID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(15 * time.Minute),
	}

	token := paseto.JSONToken{
		Subject:    userID.String(),
		IssuedAt:   payload.IssuedAt,
		Expiration: payload.ExpiredAt,
	}

	token.Set("username", username)

	v2 := paseto.NewV2()

	encrypted, err := v2.Encrypt(
		maker.symmetricKey,
		token,
		nil,
	)

	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func (maker *PasetoMaker) CreateRefreshToken(
	userID uuid.UUID,
) (string, error) {

	token := paseto.JSONToken{
		Subject:  userID.String(),
		IssuedAt: time.Now(),
		Expiration: time.Now().Add(
			7 * 24 * time.Hour,
		),
	}

	v2 := paseto.NewV2()

	encrypted, err := v2.Encrypt(
		maker.symmetricKey,
		token,
		nil,
	)

	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func (maker *PasetoMaker) VerifyToken(
	token string,
) (*Payload, error) {

	v2 := paseto.NewV2()

	var parsedToken paseto.JSONToken

	err := v2.Decrypt(
		token,
		maker.symmetricKey,
		&parsedToken,
		nil,
	)

	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(parsedToken.Subject)

	if err != nil {
		return nil, err
	}

	username := parsedToken.Get("username")

	payload := &Payload{

		UserID: userID,

		Username: username,

		IssuedAt: parsedToken.IssuedAt,

		ExpiredAt: parsedToken.Expiration,
	}

	return payload, nil
}
