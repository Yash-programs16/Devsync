package repository

import (
	"context"
	"errors"

	"github.com/Yash-programs16/devsync/internal/database"
	"github.com/Yash-programs16/devsync/internal/models"
	"github.com/google/uuid"
)

type WorkspaceRepository struct {
	db *database.Database
}

func NewWorkspaceRepository(db *database.Database) *WorkspaceRepository {
	return &WorkspaceRepository{
		db: db,
	}
}

func (r *WorkspaceRepository) Create(
	ctx context.Context,
	workspace *models.Workspace,
) error {

	const query = `
		INSERT INTO workspaces (
			id,
			name,
			slug,
			description,
			owner_id
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Pool.Exec(
		ctx,
		query,
		workspace.ID,
		workspace.Name,
		workspace.Slug,
		workspace.Description,
		workspace.OwnerID,
	)

	return err
}

func (r *WorkspaceRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Workspace, error) {

	const query = `
		SELECT
			id,
			name,
			slug,
			description,
			owner_id,
			created_at,
			updated_at
		FROM workspaces
		WHERE id = $1
	`

	var workspace models.Workspace

	err := r.db.Pool.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&workspace.ID,
		&workspace.Name,
		&workspace.Slug,
		&workspace.Description,
		&workspace.OwnerID,
		&workspace.CreatedAt,
		&workspace.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (r *WorkspaceRepository) GetBySlug(
	ctx context.Context,
	slug string,
) (*models.Workspace, error) {

	const query = `
		SELECT
			id,
			name,
			slug,
			description,
			owner_id,
			created_at,
			updated_at
		FROM workspaces
		WHERE slug = $1
	`

	var workspace models.Workspace

	err := r.db.Pool.QueryRow(
		ctx,
		query,
		slug,
	).Scan(
		&workspace.ID,
		&workspace.Name,
		&workspace.Slug,
		&workspace.Description,
		&workspace.OwnerID,
		&workspace.CreatedAt,
		&workspace.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (r *WorkspaceRepository) ListByUserID(
	ctx context.Context,
	userID uuid.UUID,
) ([]models.Workspace, error) {

	const query = `
		SELECT
			w.id,
			w.name,
			w.slug,
			w.description,
			w.owner_id,
			w.created_at,
			w.updated_at
		FROM workspaces w
		INNER JOIN workspace_members wm
			ON wm.workspace_id = w.id
		WHERE wm.user_id = $1
		ORDER BY w.created_at DESC
	`

	rows, err := r.db.Pool.Query(
		ctx,
		query,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workspaces := make([]models.Workspace, 0)

	for rows.Next() {
		var workspace models.Workspace

		if err := rows.Scan(
			&workspace.ID,
			&workspace.Name,
			&workspace.Slug,
			&workspace.Description,
			&workspace.OwnerID,
			&workspace.CreatedAt,
			&workspace.UpdatedAt,
		); err != nil {
			return nil, err
		}

		workspaces = append(workspaces, workspace)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return workspaces, nil
}

func (r *WorkspaceRepository) Update(
	ctx context.Context,
	workspace *models.Workspace,
) error {

	const query = `
		UPDATE workspaces
		SET
			name = $2,
			slug = $3,
			description = $4,
			updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.Pool.Exec(
		ctx,
		query,
		workspace.ID,
		workspace.Name,
		workspace.Slug,
		workspace.Description,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("workspace not found")
	}

	return nil
}

func (r *WorkspaceRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {

	const query = `
		DELETE FROM workspaces
		WHERE id = $1
	`

	result, err := r.db.Pool.Exec(
		ctx,
		query,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("workspace not found")
	}

	return nil
}

func (r *WorkspaceRepository) CreateWithOwner(
	ctx context.Context,
	workspace *models.Workspace,
	member *models.WorkspaceMember,
	role *models.Role,
) error {

	tx, err := r.db.Pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	// 1. Create workspace
	const workspaceQuery = `
		INSERT INTO workspaces (
			id,
			name,
			slug,
			description,
			owner_id
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = tx.Exec(
		ctx,
		workspaceQuery,
		workspace.ID,
		workspace.Name,
		workspace.Slug,
		workspace.Description,
		workspace.OwnerID,
	)
	if err != nil {
		return err
	}

	// 2. Add owner as workspace member
	const memberQuery = `
		INSERT INTO workspace_members (
			id,
			workspace_id,
			user_id
		)
		VALUES ($1, $2, $3)
	`

	_, err = tx.Exec(
		ctx,
		memberQuery,
		member.ID,
		member.WorkspaceID,
		member.UserID,
	)
	if err != nil {
		return err
	}

	// 3. Create owner role
	const roleQuery = `
		INSERT INTO roles (
			id,
			workspace_id,
			name
		)
		VALUES ($1, $2, $3)
	`

	_, err = tx.Exec(
		ctx,
		roleQuery,
		role.ID,
		role.WorkspaceID,
		role.Name,
	)
	if err != nil {
		return err
	}

	// 4. Assign owner role
	const memberRoleQuery = `
		INSERT INTO member_roles (
			member_id,
			role_id
		)
		VALUES ($1, $2)
	`

	_, err = tx.Exec(
		ctx,
		memberRoleQuery,
		member.ID,
		role.ID,
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
