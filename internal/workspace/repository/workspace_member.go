package repository

import (
	"context"
	"errors"

	"github.com/Yash-programs16/devsync/internal/database"
	"github.com/Yash-programs16/devsync/internal/models"
	"github.com/google/uuid"
)

type WorkspaceMemberRepository struct {
	db *database.Database
}

func NewWorkspaceMemberRepository(
	db *database.Database,
) *WorkspaceMemberRepository {
	return &WorkspaceMemberRepository{
		db: db,
	}
}

func (r *WorkspaceMemberRepository) Create(
	ctx context.Context,
	member *models.WorkspaceMember,
) error {

	const query = `
		INSERT INTO workspace_members (
			id,
			workspace_id,
			user_id
		)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Pool.Exec(
		ctx,
		query,
		member.ID,
		member.WorkspaceID,
		member.UserID,
	)

	return err
}

func (r *WorkspaceMemberRepository) IsMember(
	ctx context.Context,
	workspaceID uuid.UUID,
	userID uuid.UUID,
) (bool, error) {

	const query = `
		SELECT EXISTS (
			SELECT 1
			FROM workspace_members
			WHERE workspace_id = $1
			  AND user_id = $2
		)
	`

	var exists bool

	err := r.db.Pool.QueryRow(
		ctx,
		query,
		workspaceID,
		userID,
	).Scan(&exists)

	return exists, err
}

func (r *WorkspaceMemberRepository) ListByWorkspaceID(
	ctx context.Context,
	workspaceID uuid.UUID,
) ([]models.WorkspaceMember, error) {

	const query = `
		SELECT
			id,
			workspace_id,
			user_id,
			joined_at
		FROM workspace_members
		WHERE workspace_id = $1
		ORDER BY joined_at ASC
	`

	rows, err := r.db.Pool.Query(
		ctx,
		query,
		workspaceID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	members := make([]models.WorkspaceMember, 0)

	for rows.Next() {
		var member models.WorkspaceMember

		if err := rows.Scan(
			&member.ID,
			&member.WorkspaceID,
			&member.UserID,
			&member.JoinedAt,
		); err != nil {
			return nil, err
		}

		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}

func (r *WorkspaceMemberRepository) GetByWorkspaceAndUser(
	ctx context.Context,
	workspaceID uuid.UUID,
	userID uuid.UUID,
) (*models.WorkspaceMember, error) {

	const query = `
		SELECT
			id,
			workspace_id,
			user_id,
			joined_at
		FROM workspace_members
		WHERE workspace_id = $1
		  AND user_id = $2
	`

	var member models.WorkspaceMember

	err := r.db.Pool.QueryRow(
		ctx,
		query,
		workspaceID,
		userID,
	).Scan(
		&member.ID,
		&member.WorkspaceID,
		&member.UserID,
		&member.JoinedAt,
	)

	if err != nil {
		return nil, err
	}

	return &member, nil
}

func (r *WorkspaceMemberRepository) Delete(
	ctx context.Context,
	workspaceID uuid.UUID,
	userID uuid.UUID,
) error {

	const query = `
		DELETE FROM workspace_members
		WHERE workspace_id = $1
		  AND user_id = $2
	`

	result, err := r.db.Pool.Exec(
		ctx,
		query,
		workspaceID,
		userID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("workspace member not found")
	}

	return nil
}
