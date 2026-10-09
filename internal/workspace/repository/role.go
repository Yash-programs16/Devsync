package repository

import (
	"context"

	"github.com/Yash-programs16/devsync/internal/database"
	"github.com/Yash-programs16/devsync/internal/models"
	"github.com/google/uuid"
)

type RoleRepository struct {
	db *database.Database
}

func NewRoleRepository(
	db *database.Database,
) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) Create(
	ctx context.Context,
	role *models.Role,
) error {

	const query = `
		INSERT INTO roles (
			id,
			workspace_id,
			name
		)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Pool.Exec(
		ctx,
		query,
		role.ID,
		role.WorkspaceID,
		role.Name,
	)

	return err
}

func (r *RoleRepository) GetByName(
	ctx context.Context,
	workspaceID uuid.UUID,
	name string,
) (*models.Role, error) {

	const query = `
		SELECT
			id,
			workspace_id,
			name,
			created_at
		FROM roles
		WHERE workspace_id = $1
		  AND name = $2
	`

	var role models.Role

	err := r.db.Pool.QueryRow(
		ctx,
		query,
		workspaceID,
		name,
	).Scan(
		&role.ID,
		&role.WorkspaceID,
		&role.Name,
		&role.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &role, nil
}
