package repository

import (
	"context"

	"github.com/Yash-programs16/devsync/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RefreshTokenRepository struct {
	db *pgxpool.Pool
}

func NewRefreshTokenRepository(
	db *pgxpool.Pool,
) *RefreshTokenRepository {

	return &RefreshTokenRepository{
		db: db,
	}
}
func (r *RefreshTokenRepository) CreateRefreshToken(
	ctx context.Context,
	token *models.RefreshToken,
) error {

	query := `
	INSERT INTO refresh_tokens
	(
		id,
		user_id,
		token_hash,
		expires_at,
		created_at,
		revoked
	)
	VALUES
	(
		$1,$2,$3,$4,$5,$6
	)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		token.ID,
		token.UserID,
		token.TokenHash,
		token.ExpiresAt,
		token.CreatedAt,
		token.Revoked,
	)

	return err
}
func (r *RefreshTokenRepository) GetByHash(
	ctx context.Context,
	hash string,
) (*models.RefreshToken, error) {

	token := &models.RefreshToken{}

	query := `
	SELECT
		id,
		user_id,
		token_hash,
		expires_at,
		created_at,
		revoked
	FROM refresh_tokens
	WHERE token_hash=$1
	`

	err := r.db.QueryRow(
		ctx,
		query,
		hash,
	).Scan(
		&token.ID,
		&token.UserID,
		&token.TokenHash,
		&token.ExpiresAt,
		&token.CreatedAt,
		&token.Revoked,
	)

	if err != nil {
		return nil, err
	}

	return token, nil
}
func (r *RefreshTokenRepository) RevokeToken(
	ctx context.Context,
	hash string,
) error {

	query := `
UPDATE refresh_tokens
SET revoked=true
WHERE token_hash=$1
`

	_, err := r.db.Exec(
		ctx,
		query,
		hash,
	)

	return err
}
func (r *RefreshTokenRepository) DeleteExpiredTokens(
	ctx context.Context,
) error {

	query := `
	DELETE FROM refresh_tokens
	WHERE expires_at < NOW()
	`

	_, err := r.db.Exec(
		ctx,
		query,
	)

	return err
}
