package token_test

import (
	"testing"

	"github.com/Yash-programs16/devsync/internal/token"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPaseto(t *testing.T) {

	key := "your_32_character_key_here"

	maker, err := token.NewPasetoMaker(key)

	require.NoError(t, err)

	accessToken, err := maker.CreateAccessToken(
		uuid.New(),
		"yash",
	)

	require.NoError(t, err)

	payload, err := maker.VerifyToken(
		accessToken,
	)

	require.NoError(t, err)

	require.Equal(
		t,
		"yash",
		payload.Username,
	)

}
