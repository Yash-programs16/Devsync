package repository

import (
	"context"

	"github.com/Yash-programs16/devsync/internal/database"
	"github.com/google/uuid"
)

type MemberRoleRepository struct {
	db *database.Database
}

func NewMemberRoleRepository(
	db *database.Database,
) *MemberRoleRepository {
	return &MemberRoleRepository{
		db: db,
	}
}

func (r *MemberRoleRepository) AssignRole(
	ctx context.Context,
	memberID uuid.UUID,
	roleID uuid.UUID,
) error {

	const query = `
		INSERT INTO member_roles (
			member_id,
			role_id
		)
		VALUES ($1, $2)
	`

	_, err := r.db.Pool.Exec(
		ctx,
		query,
		memberID,
		roleID,
	)

	return err
}
