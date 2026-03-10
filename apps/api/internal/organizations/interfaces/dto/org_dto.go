package dto

import (
	"strings"
	"time"

	"mytodo/apps/api/internal/organizations/domain/entity"
)

type OrganizationDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
	PlanID      string `json:"plan_id"`
	OwnerID     string `json:"owner_id"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

/*
Request DTOs
*/

type CreateOrganizationRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=120"`
	Slug        string `json:"slug,omitempty" binding:"omitempty,min=2,max=120"`
	Description string `json:"description,omitempty" binding:"max=500"`
	PlanID      string `json:"plan_id,omitempty"`
}

type UpdateOrganizationRequest struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=2,max=120"`
	Slug        *string `json:"slug,omitempty" binding:"omitempty,min=2,max=120"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=500"`
	PlanID      *string `json:"plan_id,omitempty"`
}

type TransferOwnershipRequest struct {
	NewOwnerID string `json:"new_owner_id" binding:"required,uuid"`
}

/*
Response DTOs
*/

type OrganizationListResponse struct {
	Data  []*OrganizationDTO `json:"data"`
	Total int                `json:"total"`
	Page  int                `json:"page"`
	Limit int                `json:"limit"`
}

/*
Mapping functions
*/

func ToOrganizationDTO(org *entity.Organization) *OrganizationDTO {
	if org == nil {
		return nil
	}

	return &OrganizationDTO{
		ID:          org.ID.String(),
		Name:        org.Name,
		Slug:        org.Slug,
		Description: org.Description,
		PlanID:      org.PlanID,
		OwnerID:     org.OwnerID,
		IsActive:    org.IsActive,
		CreatedAt:   org.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   org.UpdatedAt.Format(time.RFC3339),
	}
}

func ToOrganizationDTOList(orgs []*entity.Organization) []*OrganizationDTO {
	if orgs == nil {
		return nil
	}

	result := make([]*OrganizationDTO, 0, len(orgs))

	for _, org := range orgs {
		result = append(result, ToOrganizationDTO(org))
	}

	return result
}

/*
Optional helper for sanitizing request data
*/

func (r *CreateOrganizationRequest) Normalize() {
	r.Name = strings.TrimSpace(r.Name)
	r.Slug = strings.TrimSpace(r.Slug)
	r.Description = strings.TrimSpace(r.Description)
	r.PlanID = strings.TrimSpace(r.PlanID)
}
