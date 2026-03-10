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

type OrganizationRepositoryImpl struct {
	db *sql.DB
}

// NewOrganizationRepository creates a new instance of OrganizationRepositoryImpl
func NewOrganizationRepository(db *sql.DB) repository.OrganizationRepository {
	return &OrganizationRepositoryImpl{db: db}
}

func (r *OrganizationRepositoryImpl) CreateOrganization(ctx context.Context, org *entity.Organization) error {

	query := `
		INSERT INTO organizations (id, name, slug, description, plan_id, owner_id, is_active, is_deleted, created_by, updated_by, created_at, updated_at, deleted_by, deleted_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`

	_, err := r.db.ExecContext(ctx, query,
		org.ID, org.Name, org.Slug, org.Description, org.PlanID,
		org.OwnerID, org.IsActive, org.IsDeleted, org.CreatedBy, org.UpdatedBy, org.CreatedAt, org.UpdatedAt, org.DeletedBy, org.DeletedAt)
	if err != nil {
		return fmt.Errorf("failed to create organization: %w", err)
	}
	return err
}

func (r *OrganizationRepositoryImpl) GetOrganizationByID(ctx context.Context, id uuid.UUID) (*entity.Organization, error) {
	query := `
		SELECT id, name, slug, description, plan_id, owner_id, is_active, is_deleted, 
		       created_by, updated_by, created_at, updated_at, deleted_by, deleted_at
		FROM organizations 
		WHERE id = $1 AND is_deleted = false
	`
	org := &entity.Organization{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&org.ID, &org.Name, &org.Slug, &org.Description, &org.PlanID,
		&org.OwnerID, &org.IsActive, &org.IsDeleted,
		&org.CreatedBy, &org.UpdatedBy, &org.CreatedAt, &org.UpdatedAt, &org.DeletedBy, &org.DeletedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("organization not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get organization: %w", err)
	}

	return org, err
}

func (r *OrganizationRepositoryImpl) UpdateOrganization(ctx context.Context, org *entity.Organization) error {
	query := `
		UPDATE organizations 
		SET name = $2, slug = $3, description = $4, plan_id = $5, 
		    owner_id = $6, is_active = $7, updated_at = $8
		WHERE id = $1 AND is_deleted = false
	`

	result, err := r.db.ExecContext(ctx, query,
		org.ID, org.Name, org.Slug, org.Description, org.PlanID,
		org.OwnerID, org.IsActive, org.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to update organization: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("organization not found or already deleted")
	}

	return nil
}

func (r *OrganizationRepositoryImpl) DeleteOrganization(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) error {
	query := `
		UPDATE organizations 
		SET is_deleted = true, deleted_at = NOW(), updated_at = NOW(), deleted_by = $2
		WHERE id = $1 AND is_deleted = false
	`

	result, err := r.db.ExecContext(ctx, query, id, deletedBy)
	if err != nil {
		return fmt.Errorf("failed to delete organization: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("organization not found or already deleted")
	}

	return nil
}

func (r *OrganizationRepositoryImpl) RestoreOrganization(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE organizations 
		SET is_deleted = false, deleted_at = NULL, updated_at = NOW()
		WHERE id = $1 AND is_deleted = true
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to restore organization: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("organization not found or not deleted")
	}

	return nil
}

func (r *OrganizationRepositoryImpl) ListOrganizations(ctx context.Context, page, limit int) ([]*entity.Organization, int, error) {
	offset := (page - 1) * limit

	// Get total count
	countQuery := `SELECT COUNT(*) FROM organizations WHERE is_deleted = false`
	var total int
	err := r.db.QueryRowContext(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count organizations: %w", err)
	}

	// Get paginated results
	query := `
		SELECT id, name, slug, description, plan_id, owner_id, is_active, is_deleted, 
		       created_by, updated_by, created_at, updated_at, deleted_by, deleted_at
		FROM organizations 
		WHERE is_deleted = false
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list organizations: %w", err)
	}
	defer rows.Close()

	organizations := make([]*entity.Organization, 0)
	for rows.Next() {
		org := &entity.Organization{}
		err := rows.Scan(
			&org.ID, &org.Name, &org.Slug, &org.Description, &org.PlanID,
			&org.OwnerID, &org.IsActive, &org.IsDeleted,
			&org.CreatedBy, &org.UpdatedBy, &org.CreatedAt, &org.UpdatedAt, &org.DeletedBy, &org.DeletedAt)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan organization: %w", err)
		}
		organizations = append(organizations, org)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating organizations: %w", err)
	}

	return organizations, total, nil
}

func (r *OrganizationRepositoryImpl) SearchOrganizations(ctx context.Context, query string, limit int) ([]*entity.Organization, error) {
	searchQuery := `
		SELECT id, name, slug, description, plan_id, owner_id, is_active, is_deleted, 
		       created_by, updated_by, created_at, updated_at, deleted_by, deleted_at
		FROM organizations 
		WHERE is_deleted = false 
		  AND (name ILIKE $1 OR description ILIKE $1 OR slug ILIKE $1)
		ORDER BY name
		LIMIT $2
	`

	searchPattern := "%" + query + "%"
	rows, err := r.db.QueryContext(ctx, searchQuery, searchPattern, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to search organizations: %w", err)
	}
	defer rows.Close()

	organizations := make([]*entity.Organization, 0)
	for rows.Next() {
		org := &entity.Organization{}
		err := rows.Scan(
			&org.ID, &org.Name, &org.Slug, &org.Description, &org.PlanID,
			&org.OwnerID, &org.IsActive, &org.IsDeleted,
			&org.CreatedBy, &org.UpdatedBy, &org.CreatedAt, &org.UpdatedAt, &org.DeletedBy, &org.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan organization: %w", err)
		}
		organizations = append(organizations, org)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating organizations: %w", err)
	}

	return organizations, nil
}

func (r *OrganizationRepositoryImpl) TransferOrganizationOwnership(ctx context.Context, orgID, newOwnerID uuid.UUID) error {
	query := `
		UPDATE organizations 
		SET owner_id = $2, updated_at = NOW()
		WHERE id = $1 AND is_deleted = false
	`

	result, err := r.db.ExecContext(ctx, query, orgID, newOwnerID)
	if err != nil {
		return fmt.Errorf("failed to transfer ownership: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("organization not found or already deleted")
	}

	return nil
}

func (r *OrganizationRepositoryImpl) GetOrganizationsByOwner(ctx context.Context, ownerID uuid.UUID) ([]*entity.Organization, error) {
	query := `
		SELECT id, name, slug, description, plan_id, owner_id, is_active, is_deleted, 
		       created_by, updated_by, created_at, updated_at, deleted_by, deleted_at
		FROM organizations 
		WHERE owner_id = $1 AND is_deleted = false
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, ownerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get organizations by owner: %w", err)
	}
	defer rows.Close()

	organizations := make([]*entity.Organization, 0)
	for rows.Next() {
		org := &entity.Organization{}
		err := rows.Scan(
			&org.ID, &org.Name, &org.Slug, &org.Description, &org.PlanID,
			&org.OwnerID, &org.IsActive, &org.IsDeleted,
			&org.CreatedBy, &org.UpdatedBy, &org.CreatedAt, &org.UpdatedAt, &org.DeletedBy, &org.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan organization: %w", err)
		}
		organizations = append(organizations, org)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating organizations: %w", err)
	}

	return organizations, nil
}

func (r *OrganizationRepositoryImpl) GetMemberOrganizations(ctx context.Context, userID uuid.UUID, page, limit int) ([]*entity.Organization, int, error) {
	offset := (page - 1) * limit

	// Get total count
	countQuery := `
		SELECT COUNT(DISTINCT o.id)
		FROM organizations o
		INNER JOIN organization_members om ON o.id = om.organization_id
		WHERE om.user_id = $1 AND o.is_deleted = false
	`
	var total int
	err := r.db.QueryRowContext(ctx, countQuery, userID).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count member organizations: %w", err)
	}

	// Get paginated results
	query := `
		SELECT DISTINCT o.id, o.name, o.slug, o.description, o.plan_id, o.owner_id, 
		       o.is_active, o.is_deleted, o.created_by, o.updated_by, 
		       o.created_at, o.updated_at, o.deleted_by, o.deleted_at
		FROM organizations o
		INNER JOIN organization_members om ON o.id = om.organization_id
		WHERE om.user_id = $1 AND o.is_deleted = false
		ORDER BY o.created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get member organizations: %w", err)
	}
	defer rows.Close()

	organizations := make([]*entity.Organization, 0)
	for rows.Next() {
		org := &entity.Organization{}
		err := rows.Scan(
			&org.ID, &org.Name, &org.Slug, &org.Description, &org.PlanID,
			&org.OwnerID, &org.IsActive, &org.IsDeleted,
			&org.CreatedBy, &org.UpdatedBy, &org.CreatedAt, &org.UpdatedAt, &org.DeletedBy, &org.DeletedAt)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan organization: %w", err)
		}
		organizations = append(organizations, org)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating organizations: %w", err)
	}

	return organizations, total, nil
}

func (r *OrganizationRepositoryImpl) IsOwnedBy(ctx context.Context, orgID, userID uuid.UUID) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM organizations 
			WHERE id = $1 AND owner_id = $2 AND is_deleted = false
		)
	`

	var isOwner bool
	err := r.db.QueryRowContext(ctx, query, orgID, userID).Scan(&isOwner)
	if err != nil {
		return false, fmt.Errorf("failed to check ownership: %w", err)
	}

	return isOwner, nil
}
