package repository

import (
	"context"
	"time"

	"github.com/Yash-programs16/devsync/internal/database"
	"github.com/Yash-programs16/devsync/internal/models"
	"github.com/google/uuid"
)

type InvitationRepository struct {
	db *database.Database
}

func NewInvitationRepository(
	db *database.Database,
) *InvitationRepository {
	return &InvitationRepository{
		db: db,
	}
}

func (r *InvitationRepository) Create(
	ctx context.Context,
	invitation *models.Invitation,
	tokenHash string,
) error {
	const query = `
		INSERT INTO workspace_invites (
			id,
			workspace_id,
			inviter_id,
			email,
			token_hash,
			expires_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Pool.Exec(
		ctx,
		query,
		invitation.ID,
		invitation.WorkspaceID,
		invitation.InvitedBy,
		invitation.Email,
		tokenHash,
		invitation.ExpiresAt,
	)

	return err
}

func (r *InvitationRepository) GetByTokenHash(
	ctx context.Context,
	tokenHash string,
) (*models.Invitation, error) {
	const query = `
		SELECT
			id,
			workspace_id,
			inviter_id,
			email,
			expires_at,
			accepted_at,
			created_at
		FROM workspace_invites
		WHERE token_hash = $1
	`

	var invitation models.Invitation
	var acceptedAt *time.Time

	err := r.db.Pool.QueryRow(
		ctx,
		query,
		tokenHash,
	).Scan(
		&invitation.ID,
		&invitation.WorkspaceID,
		&invitation.InvitedBy,
		&invitation.Email,
		&invitation.ExpiresAt,
		&acceptedAt,
		&invitation.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	invitation.Accepted = acceptedAt != nil

	return &invitation, nil
}

func (r *InvitationRepository) MarkAccepted(
	ctx context.Context,
	id uuid.UUID,
) error {
	const query = `
		UPDATE workspace_invites
		SET accepted_at = NOW()
		WHERE id = $1
		  AND accepted_at IS NULL
	`

	_, err := r.db.Pool.Exec(
		ctx,
		query,
		id,
	)

	return err
}
