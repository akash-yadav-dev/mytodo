package repository

import (
	"context"
	"mytodo/apps/api/internal/organizations/domain/entity"

	"github.com/google/uuid"
)

type MembershipRepository interface {
	// Member management
	AddMember(ctx context.Context, member *entity.OrganizationMember) error
	RemoveMember(ctx context.Context, orgID, userID uuid.UUID) error
	UpdateMemberRole(ctx context.Context, orgID, userID uuid.UUID, role entity.Role) error
	GetMember(ctx context.Context, orgID, userID uuid.UUID) (*entity.OrganizationMember, error)
	ListMembers(ctx context.Context, orgID uuid.UUID, page, limit int) ([]*entity.OrganizationMember, int, error)
	IsMember(ctx context.Context, orgID, userID uuid.UUID) (bool, error)
	GetMemberRole(ctx context.Context, orgID, userID uuid.UUID) (entity.Role, error)
}
