package repository

import (
	"context"
	"mytodo/apps/api/internal/organizations/domain/entity"

	"github.com/google/uuid"
)

type OrganizationRepository interface {
	// Define methods for organization data access, e.g.:
	CreateOrganization(ctx context.Context, org *entity.Organization) error
	GetOrganizationByID(ctx context.Context, id uuid.UUID) (*entity.Organization, error)
	UpdateOrganization(ctx context.Context, org *entity.Organization) error
	DeleteOrganization(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) error
	RestoreOrganization(ctx context.Context, id uuid.UUID) error
	ListOrganizations(ctx context.Context, page, limit int) ([]*entity.Organization, int, error)
	SearchOrganizations(ctx context.Context, query string, limit int) ([]*entity.Organization, error)
	TransferOrganizationOwnership(ctx context.Context, orgID, newOwnerID uuid.UUID) error
	GetOrganizationsByOwner(ctx context.Context, ownerID uuid.UUID) ([]*entity.Organization, error)
	GetMemberOrganizations(ctx context.Context, userID uuid.UUID, page, limit int) ([]*entity.Organization, int, error)

	// persistence-agnostic methods for organization data access
	IsOwnedBy(ctx context.Context, orgID, userID uuid.UUID) (bool, error)
}
