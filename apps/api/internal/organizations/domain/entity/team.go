package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidTeamName = errors.New("team name cannot be empty")
)

// Team represents a team within an organization
type Team struct {
	ID             uuid.UUID  `json:"id"`
	OrganizationID uuid.UUID  `json:"organization_id"`
	Name           string     `json:"name"`
	Slug           string     `json:"slug"`
	Description    string     `json:"description"`
	IsActive       bool       `json:"is_active"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
}

// NewTeam creates a new team
func NewTeam(orgID uuid.UUID, name, slug, description string) (*Team, error) {
	if strings.TrimSpace(name) == "" {
		return nil, ErrInvalidTeamName
	}

	team := &Team{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Name:           name,
		Slug:           slug,
		Description:    description,
		IsActive:       true,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}

	return team, nil
}

// UpdateName updates the team's name
func (t *Team) UpdateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return ErrInvalidTeamName
	}

	t.Name = name
	t.UpdatedAt = time.Now().UTC()
	return nil
}

// Deactivate deactivates the team
func (t *Team) Deactivate() {
	t.IsActive = false
	t.UpdatedAt = time.Now().UTC()
}

// Activate activates the team
func (t *Team) Activate() {
	t.IsActive = true
	t.UpdatedAt = time.Now().UTC()
}

// TeamMember represents a member of a team
type TeamMember struct {
	ID        uuid.UUID `json:"id"`
	TeamID    uuid.UUID `json:"team_id"`
	UserID    uuid.UUID `json:"user_id"`
	Role      Role      `json:"role"`
	JoinedAt  time.Time `json:"joined_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewTeamMember creates a new team member
func NewTeamMember(teamID, userID uuid.UUID, role Role) (*TeamMember, error) {
	if !IsValidRole(role) {
		return nil, ErrInvalidRole
	}

	member := &TeamMember{
		ID:        uuid.New(),
		TeamID:    teamID,
		UserID:    userID,
		Role:      role,
		JoinedAt:  time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	return member, nil
}
