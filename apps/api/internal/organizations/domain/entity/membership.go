package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidRole   = errors.New("invalid role")
	ErrAlreadyMember = errors.New("user is already a member")
)

type Role string

const (
	RoleOwner  Role = "owner"
	RoleAdmin  Role = "admin"
	RoleMember Role = "member"
	RoleGuest  Role = "guest"
)

// OrganizationMember represents a user's membership in an organization
type OrganizationMember struct {
	ID             uuid.UUID  `json:"id"`
	OrganizationID uuid.UUID  `json:"organization_id"`
	UserID         uuid.UUID  `json:"user_id"`
	Role           Role       `json:"role"`
	JoinedAt       time.Time  `json:"joined_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	InvitedBy      *uuid.UUID `json:"invited_by,omitempty"`
}

// NewOrganizationMember creates a new organization member
func NewOrganizationMember(orgID, userID uuid.UUID, role Role, invitedBy *uuid.UUID) (*OrganizationMember, error) {
	if !IsValidRole(role) {
		return nil, ErrInvalidRole
	}

	member := &OrganizationMember{
		ID:             uuid.New(),
		OrganizationID: orgID,
		UserID:         userID,
		Role:           role,
		JoinedAt:       time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		InvitedBy:      invitedBy,
	}

	return member, nil
}

// IsValidRole checks if the role is valid
func IsValidRole(role Role) bool {
	switch role {
	case RoleOwner, RoleAdmin, RoleMember, RoleGuest:
		return true
	default:
		return false
	}
}

// UpdateRole updates the member's role
func (m *OrganizationMember) UpdateRole(newRole Role) error {
	if !IsValidRole(newRole) {
		return ErrInvalidRole
	}

	m.Role = newRole
	m.UpdatedAt = time.Now().UTC()
	return nil
}

// HasPermission checks if the member has a specific permission level
func (m *OrganizationMember) HasPermission(requiredRole Role) bool {
	roleHierarchy := map[Role]int{
		RoleGuest:  1,
		RoleMember: 2,
		RoleAdmin:  3,
		RoleOwner:  4,
	}

	memberLevel := roleHierarchy[m.Role]
	requiredLevel := roleHierarchy[requiredRole]

	return memberLevel >= requiredLevel
}
