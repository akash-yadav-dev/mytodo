package dto

import (
	"mytodo/apps/api/internal/organizations/domain/entity"
	"time"
)

// MemberDTO represents an organization member response
type MemberDTO struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organization_id"`
	UserID         string `json:"user_id"`
	Role           string `json:"role"`
	JoinedAt       string `json:"joined_at"`
	InvitedBy      string `json:"invited_by,omitempty"`
}

// TeamDTO represents a team response
type TeamDTO struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organization_id"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	Description    string `json:"description,omitempty"`
	IsActive       bool   `json:"is_active"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// TeamMemberDTO represents a team member response
type TeamMemberDTO struct {
	ID       string `json:"id"`
	TeamID   string `json:"team_id"`
	UserID   string `json:"user_id"`
	Role     string `json:"role"`
	JoinedAt string `json:"joined_at"`
}

// Request DTOs

type AddMemberRequest struct {
	UserID string `json:"user_id" binding:"required,uuid"`
	Role   string `json:"role" binding:"required,oneof=owner admin member guest"`
}

type UpdateMemberRoleRequest struct {
	Role string `json:"role" binding:"required,oneof=owner admin member guest"`
}

type CreateTeamRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=120"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty" binding:"max=500"`
}

type UpdateTeamRequest struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=2,max=120"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=500"`
}

type AddTeamMemberRequest struct {
	UserID string `json:"user_id" binding:"required,uuid"`
	Role   string `json:"role" binding:"required,oneof=admin member"`
}

// Response DTOs

type MemberListResponse struct {
	Data  []*MemberDTO `json:"data"`
	Total int          `json:"total"`
	Page  int          `json:"page"`
	Limit int          `json:"limit"`
}

// Mapping functions

func ToMemberDTO(member *entity.OrganizationMember) *MemberDTO {
	if member == nil {
		return nil
	}

	dto := &MemberDTO{
		ID:             member.ID.String(),
		OrganizationID: member.OrganizationID.String(),
		UserID:         member.UserID.String(),
		Role:           string(member.Role),
		JoinedAt:       member.JoinedAt.Format(time.RFC3339),
	}

	if member.InvitedBy != nil {
		dto.InvitedBy = member.InvitedBy.String()
	}

	return dto
}

func ToMemberDTOList(members []*entity.OrganizationMember) []*MemberDTO {
	if members == nil {
		return nil
	}

	result := make([]*MemberDTO, 0, len(members))
	for _, member := range members {
		result = append(result, ToMemberDTO(member))
	}

	return result
}

func ToTeamDTO(team *entity.Team) *TeamDTO {
	if team == nil {
		return nil
	}

	return &TeamDTO{
		ID:             team.ID.String(),
		OrganizationID: team.OrganizationID.String(),
		Name:           team.Name,
		Slug:           team.Slug,
		Description:    team.Description,
		IsActive:       team.IsActive,
		CreatedAt:      team.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      team.UpdatedAt.Format(time.RFC3339),
	}
}

func ToTeamDTOList(teams []*entity.Team) []*TeamDTO {
	if teams == nil {
		return nil
	}

	result := make([]*TeamDTO, 0, len(teams))
	for _, team := range teams {
		result = append(result, ToTeamDTO(team))
	}

	return result
}

func ToTeamMemberDTO(member *entity.TeamMember) *TeamMemberDTO {
	if member == nil {
		return nil
	}

	return &TeamMemberDTO{
		ID:       member.ID.String(),
		TeamID:   member.TeamID.String(),
		UserID:   member.UserID.String(),
		Role:     string(member.Role),
		JoinedAt: member.JoinedAt.Format(time.RFC3339),
	}
}

func ToTeamMemberDTOList(members []*entity.TeamMember) []*TeamMemberDTO {
	if members == nil {
		return nil
	}

	result := make([]*TeamMemberDTO, 0, len(members))
	for _, member := range members {
		result = append(result, ToTeamMemberDTO(member))
	}

	return result
}
