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

type TeamRepositoryImpl struct {
	db *sql.DB
}

// NewTeamRepository creates a new instance of TeamRepositoryImpl
func NewTeamRepository(db *sql.DB) repository.TeamRepository {
	return &TeamRepositoryImpl{db: db}
}

func (r *TeamRepositoryImpl) CreateTeam(ctx context.Context, team *entity.Team) error {
	query := `
		INSERT INTO teams (id, organization_id, name, slug, description, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.ExecContext(ctx, query,
		team.ID, team.OrganizationID, team.Name, team.Slug, team.Description,
		team.IsActive, team.CreatedAt, team.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create team: %w", err)
	}

	return nil
}

func (r *TeamRepositoryImpl) GetTeamByID(ctx context.Context, id uuid.UUID) (*entity.Team, error) {
	query := `
		SELECT id, organization_id, name, slug, description, is_active, created_at, updated_at, deleted_at
		FROM teams
		WHERE id = $1 AND deleted_at IS NULL
	`

	team := &entity.Team{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&team.ID, &team.OrganizationID, &team.Name, &team.Slug, &team.Description,
		&team.IsActive, &team.CreatedAt, &team.UpdatedAt, &team.DeletedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("team not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get team: %w", err)
	}

	return team, nil
}

func (r *TeamRepositoryImpl) UpdateTeam(ctx context.Context, team *entity.Team) error {
	query := `
		UPDATE teams
		SET name = $2, slug = $3, description = $4, is_active = $5, updated_at = $6
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query,
		team.ID, team.Name, team.Slug, team.Description, team.IsActive, team.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to update team: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("team not found or deleted")
	}

	return nil
}

func (r *TeamRepositoryImpl) DeleteTeam(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE teams
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete team: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("team not found or already deleted")
	}

	return nil
}

func (r *TeamRepositoryImpl) ListTeamsByOrganization(ctx context.Context, orgID uuid.UUID) ([]*entity.Team, error) {
	query := `
		SELECT id, organization_id, name, slug, description, is_active, created_at, updated_at, deleted_at
		FROM teams
		WHERE organization_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, orgID)
	if err != nil {
		return nil, fmt.Errorf("failed to list teams: %w", err)
	}
	defer rows.Close()

	teams := make([]*entity.Team, 0)
	for rows.Next() {
		team := &entity.Team{}
		err := rows.Scan(
			&team.ID, &team.OrganizationID, &team.Name, &team.Slug, &team.Description,
			&team.IsActive, &team.CreatedAt, &team.UpdatedAt, &team.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan team: %w", err)
		}
		teams = append(teams, team)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating teams: %w", err)
	}

	return teams, nil
}

func (r *TeamRepositoryImpl) SearchTeams(ctx context.Context, orgID uuid.UUID, query string) ([]*entity.Team, error) {
	searchQuery := `
		SELECT id, organization_id, name, slug, description, is_active, created_at, updated_at, deleted_at
		FROM teams
		WHERE organization_id = $1 AND deleted_at IS NULL
		  AND (name ILIKE $2 OR description ILIKE $2 OR slug ILIKE $2)
		ORDER BY name
	`

	searchPattern := "%" + query + "%"
	rows, err := r.db.QueryContext(ctx, searchQuery, orgID, searchPattern)
	if err != nil {
		return nil, fmt.Errorf("failed to search teams: %w", err)
	}
	defer rows.Close()

	teams := make([]*entity.Team, 0)
	for rows.Next() {
		team := &entity.Team{}
		err := rows.Scan(
			&team.ID, &team.OrganizationID, &team.Name, &team.Slug, &team.Description,
			&team.IsActive, &team.CreatedAt, &team.UpdatedAt, &team.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan team: %w", err)
		}
		teams = append(teams, team)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating teams: %w", err)
	}

	return teams, nil
}

func (r *TeamRepositoryImpl) AddTeamMember(ctx context.Context, member *entity.TeamMember) error {
	query := `
		INSERT INTO team_members (id, team_id, user_id, role, joined_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		member.ID, member.TeamID, member.UserID, member.Role,
		member.JoinedAt, member.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to add team member: %w", err)
	}

	return nil
}

func (r *TeamRepositoryImpl) RemoveTeamMember(ctx context.Context, teamID, userID uuid.UUID) error {
	query := `DELETE FROM team_members WHERE team_id = $1 AND user_id = $2`

	result, err := r.db.ExecContext(ctx, query, teamID, userID)
	if err != nil {
		return fmt.Errorf("failed to remove team member: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("team member not found")
	}

	return nil
}

func (r *TeamRepositoryImpl) ListTeamMembers(ctx context.Context, teamID uuid.UUID) ([]*entity.TeamMember, error) {
	query := `
		SELECT id, team_id, user_id, role, joined_at, updated_at
		FROM team_members
		WHERE team_id = $1
		ORDER BY joined_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to list team members: %w", err)
	}
	defer rows.Close()

	members := make([]*entity.TeamMember, 0)
	for rows.Next() {
		member := &entity.TeamMember{}
		err := rows.Scan(
			&member.ID, &member.TeamID, &member.UserID, &member.Role,
			&member.JoinedAt, &member.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan team member: %w", err)
		}
		members = append(members, member)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating team members: %w", err)
	}

	return members, nil
}

func (r *TeamRepositoryImpl) GetTeamsByUser(ctx context.Context, userID uuid.UUID) ([]*entity.Team, error) {
	query := `
		SELECT t.id, t.organization_id, t.name, t.slug, t.description, t.is_active, t.created_at, t.updated_at, t.deleted_at
		FROM teams t
		INNER JOIN team_members tm ON t.id = tm.team_id
		WHERE tm.user_id = $1 AND t.deleted_at IS NULL
		ORDER BY t.created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get teams by user: %w", err)
	}
	defer rows.Close()

	teams := make([]*entity.Team, 0)
	for rows.Next() {
		team := &entity.Team{}
		err := rows.Scan(
			&team.ID, &team.OrganizationID, &team.Name, &team.Slug, &team.Description,
			&team.IsActive, &team.CreatedAt, &team.UpdatedAt, &team.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan team: %w", err)
		}
		teams = append(teams, team)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating teams: %w", err)
	}

	return teams, nil
}
