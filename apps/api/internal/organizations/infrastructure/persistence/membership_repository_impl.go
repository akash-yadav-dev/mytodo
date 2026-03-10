package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mytodo/apps/api/internal/organizations/domain/entity"
	"mytodo/apps/api/internal/organizations/domain/repository"

	"github.com/google/uuid"
)

type MembershipRepositoryImpl struct {
	db *sql.DB
}

// NewMembershipRepository creates a new instance of MembershipRepositoryImpl
func NewMembershipRepository(db *sql.DB) repository.MembershipRepository {
	return &MembershipRepositoryImpl{db: db}
}

func (r *MembershipRepositoryImpl) AddMember(ctx context.Context, member *entity.OrganizationMember) error {
	query := `
		INSERT INTO organization_members (id, organization_id, user_id, role, joined_at, updated_at, invited_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(ctx, query,
		member.ID, member.OrganizationID, member.UserID, member.Role,
		member.JoinedAt, member.UpdatedAt, member.InvitedBy)
	if err != nil {
		return fmt.Errorf("failed to add member: %w", err)
	}
	return nil
}

func (r *MembershipRepositoryImpl) RemoveMember(ctx context.Context, orgID, userID uuid.UUID) error {
	query := `DELETE FROM organization_members WHERE organization_id = $1 AND user_id = $2`

	result, err := r.db.ExecContext(ctx, query, orgID, userID)
	if err != nil {
		return fmt.Errorf("failed to remove member: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("member not found")
	}

	return nil
}

func (r *MembershipRepositoryImpl) UpdateMemberRole(ctx context.Context, orgID, userID uuid.UUID, role entity.Role) error {
	query := `
		UPDATE organization_members 
		SET role = $3, updated_at = NOW()
		WHERE organization_id = $1 AND user_id = $2
	`

	result, err := r.db.ExecContext(ctx, query, orgID, userID, role)
	if err != nil {
		return fmt.Errorf("failed to update member role: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("member not found")
	}

	return nil
}

func (r *MembershipRepositoryImpl) GetMember(ctx context.Context, orgID, userID uuid.UUID) (*entity.OrganizationMember, error) {
	query := `
		SELECT id, organization_id, user_id, role, joined_at, updated_at, invited_by
		FROM organization_members
		WHERE organization_id = $1 AND user_id = $2
	`

	member := &entity.OrganizationMember{}
	err := r.db.QueryRowContext(ctx, query, orgID, userID).Scan(
		&member.ID, &member.OrganizationID, &member.UserID, &member.Role,
		&member.JoinedAt, &member.UpdatedAt, &member.InvitedBy)

	if err == sql.ErrNoRows {
		return nil, errors.New("member not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get member: %w", err)
	}

	return member, nil
}

func (r *MembershipRepositoryImpl) ListMembers(ctx context.Context, orgID uuid.UUID, page, limit int) ([]*entity.OrganizationMember, int, error) {
	offset := (page - 1) * limit

	// Get total count
	countQuery := `SELECT COUNT(*) FROM organization_members WHERE organization_id = $1`
	var total int
	err := r.db.QueryRowContext(ctx, countQuery, orgID).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count members: %w", err)
	}

	// Get paginated results
	query := `
		SELECT id, organization_id, user_id, role, joined_at, updated_at, invited_by
		FROM organization_members
		WHERE organization_id = $1
		ORDER BY joined_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, orgID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list members: %w", err)
	}
	defer rows.Close()

	members := make([]*entity.OrganizationMember, 0)
	for rows.Next() {
		member := &entity.OrganizationMember{}
		err := rows.Scan(
			&member.ID, &member.OrganizationID, &member.UserID, &member.Role,
			&member.JoinedAt, &member.UpdatedAt, &member.InvitedBy)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan member: %w", err)
		}
		members = append(members, member)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating members: %w", err)
	}

	return members, total, nil
}

func (r *MembershipRepositoryImpl) IsMember(ctx context.Context, orgID, userID uuid.UUID) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM organization_members
			WHERE organization_id = $1 AND user_id = $2
		)
	`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, orgID, userID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check membership: %w", err)
	}

	return exists, nil
}

func (r *MembershipRepositoryImpl) GetMemberRole(ctx context.Context, orgID, userID uuid.UUID) (entity.Role, error) {
	query := `SELECT role FROM organization_members WHERE organization_id = $1 AND user_id = $2`

	var role entity.Role
	err := r.db.QueryRowContext(ctx, query, orgID, userID).Scan(&role)

	if err == sql.ErrNoRows {
		return "", errors.New("member not found")
	}
	if err != nil {
		return "", fmt.Errorf("failed to get member role: %w", err)
	}

	return role, nil
}
