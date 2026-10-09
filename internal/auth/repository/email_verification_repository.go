package repository

import (
	"context"

	"github.com/Yash-programs16/devsync/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmailVerificationRepository struct {
	db *pgxpool.Pool
}

func NewEmailVerificationRepository(
	db *pgxpool.Pool,
) *EmailVerificationRepository {

	return &EmailVerificationRepository{
		db: db,
	}
}
func (r *EmailVerificationRepository) Create(
	ctx context.Context,
	verification *models.EmailVerification,
) error {

	query := `
	INSERT INTO email_verifications
	(
		id,
		user_id,
		email,
		otp_hash,
		expires_at,
		verified,
		created_at
	)
	VALUES
	(
		$1,$2,$3,$4,$5,$6,$7
	)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		verification.ID,
		verification.UserID,
		verification.Email,
		verification.OTPHash,
		verification.ExpiresAt,
		verification.Verified,
		verification.CreatedAt,
	)

	return err
}
func (r *EmailVerificationRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*models.EmailVerification, error) {

	verification := &models.EmailVerification{}

	query := `
	SELECT
		id,
		user_id,
		email,
		otp_hash,
		expires_at,
		verified,
		created_at
	FROM email_verifications
	WHERE email=$1
	ORDER BY created_at DESC
	LIMIT 1
	`

	err := r.db.QueryRow(
		ctx,
		query,
		email,
	).Scan(
		&verification.ID,
		&verification.UserID,
		&verification.Email,
		&verification.OTPHash,
		&verification.ExpiresAt,
		&verification.Verified,
		&verification.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return verification, nil
}
func (r *EmailVerificationRepository) MarkVerified(
	ctx context.Context,
	id uuid.UUID,
) error {

	query := `
	UPDATE email_verifications
	SET verified=true
	WHERE id=$1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	return err
}
func (r *EmailVerificationRepository) DeleteExpired(
	ctx context.Context,
) error {

	query := `
	DELETE FROM email_verifications
	WHERE expires_at < NOW()
	`

	_, err := r.db.Exec(
		ctx,
		query,
	)

	return err
}
