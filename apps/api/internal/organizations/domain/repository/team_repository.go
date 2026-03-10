package repository

import (
	"context"
	"mytodo/apps/api/internal/organizations/domain/entity"

	"github.com/google/uuid"
)

type TeamRepository interface {
	// Team CRUD operations
	CreateTeam(ctx context.Context, team *entity.Team) error
	GetTeamByID(ctx context.Context, id uuid.UUID) (*entity.Team, error)
	UpdateTeam(ctx context.Context, team *entity.Team) error
	DeleteTeam(ctx context.Context, id uuid.UUID) error
	ListTeamsByOrganization(ctx context.Context, orgID uuid.UUID) ([]*entity.Team, error)
	SearchTeams(ctx context.Context, orgID uuid.UUID, query string) ([]*entity.Team, error)

	// Team member operations
	AddTeamMember(ctx context.Context, member *entity.TeamMember) error
	RemoveTeamMember(ctx context.Context, teamID, userID uuid.UUID) error
	ListTeamMembers(ctx context.Context, teamID uuid.UUID) ([]*entity.TeamMember, error)
	GetTeamsByUser(ctx context.Context, userID uuid.UUID) ([]*entity.Team, error)
}
