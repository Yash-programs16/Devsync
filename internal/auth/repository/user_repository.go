package repository

import (
	"context"

	"github.com/Yash-programs16/devsync/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {

	return &UserRepository{
		db: db,
	}

}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	user *models.User,
) error {

	query := `
	INSERT INTO users
	(
		id,
		username,
		email,
		password_hash,
		display_name,
		avatar_url,
		status
	)
	VALUES
	(
		$1,$2,$3,$4,$5,$6,$7
	)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		user.ID,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.DisplayName,
		user.AvatarURL,
		user.Status,
	)

	return err
}
func (r *UserRepository) GetUserByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {

	user := &models.User{}

	query := `
SELECT
        id,
        username,
        email,
        password_hash,
        display_name,
        avatar_url,
        status,
        email_verified,
        created_at,
        updated_at
FROM users
WHERE email=$1
`

	err := r.db.QueryRow(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.DisplayName,
		&user.AvatarURL,
		&user.Status,
		&user.EmailVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.User, error) {

	user := &models.User{}

	query := `
	SELECT
		id,
		username,
		email,
		password_hash,
		display_name,
		avatar_url,
		status,
		created_at,
		updated_at
	FROM users
	WHERE id=$1
	`

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.DisplayName,
		&user.AvatarURL,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *UserRepository) GetUserByUsername(
	ctx context.Context,
	username string,
) (*models.User, error) {

	user := &models.User{}

	query := `
	SELECT
		id,
		username,
		email,
		password_hash,
		display_name,
		avatar_url,
		status,
		created_at,
		updated_at
	FROM users
	WHERE username=$1
	`

	err := r.db.QueryRow(
		ctx,
		query,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.DisplayName,
		&user.AvatarURL,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *UserRepository) UpdateLastLogin(
	ctx context.Context,
	id uuid.UUID,
) error {

	query := `
	UPDATE users
	SET last_login_at = NOW()
	WHERE id=$1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	return err
}
func (r *UserRepository) MarkEmailVerified(
	ctx context.Context,
	userID uuid.UUID,
) error {

	query := `
	UPDATE users
	SET email_verified=true
	WHERE id=$1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		userID,
	)

	return err
}
