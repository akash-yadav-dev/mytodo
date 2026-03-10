package handlers

import (
	"context"
	"fmt"
	"mytodo/apps/api/internal/organizations/application/queries"
	"mytodo/apps/api/internal/organizations/domain/service"
	"mytodo/apps/api/internal/organizations/interfaces/dto"
)

type OrganizationQueryHandler struct {
	organizationService service.OrganizationService
}

func NewOrganizationQueryHandler(organizationService service.OrganizationService) *OrganizationQueryHandler {
	return &OrganizationQueryHandler{
		organizationService: organizationService,
	}
}

func (h *OrganizationQueryHandler) HandleGetOrganization(ctx context.Context, query queries.GetOrganizationQuery) (*dto.OrganizationDTO, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}
	
	org, err := h.organizationService.GetOrganizationByID(ctx, query.OrgID)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization: %w", err)
	}
	
	return dto.ToOrganizationDTO(org), nil
}

func (h *OrganizationQueryHandler) HandleListOrganizations(ctx context.Context, query queries.ListOrganizationsQuery) (*dto.OrganizationListResponse, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}
	
	orgs, total, err := h.organizationService.ListOrganizations(ctx, query.Page, query.Limit)
	if err != nil {
		return nil, fmt.Errorf("failed to list organizations: %w", err)
	}
	
	return &dto.OrganizationListResponse{
		Data:  dto.ToOrganizationDTOList(orgs),
		Total: total,
		Page:  query.Page,
		Limit: query.Limit,
	}, nil
}

func (h *OrganizationQueryHandler) HandleSearchOrganizations(ctx context.Context, query queries.SearchOrganizationsQuery) ([]*dto.OrganizationDTO, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}
	
	orgs, err := h.organizationService.SearchOrganizations(ctx, query.Query, query.Limit)
	if err != nil {
		return nil, fmt.Errorf("failed to search organizations: %w", err)
	}
	
	return dto.ToOrganizationDTOList(orgs), nil
}

func (h *OrganizationQueryHandler) HandleGetOrganizationsByOwner(ctx context.Context, query queries.GetOrganizationsByOwnerQuery) ([]*dto.OrganizationDTO, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}
	
	orgs, err := h.organizationService.GetOrganizationsByOwner(ctx, query.OwnerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get organizations by owner: %w", err)
	}
	
	return dto.ToOrganizationDTOList(orgs), nil
}

func (h *OrganizationQueryHandler) HandleGetMemberOrganizations(ctx context.Context, query queries.GetMemberOrganizationsQuery) (*dto.OrganizationListResponse, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}
	
	orgs, total, err := h.organizationService.GetMemberOrganizations(ctx, query.UserID, query.Page, query.Limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get member organizations: %w", err)
	}
	
	return &dto.OrganizationListResponse{
		Data:  dto.ToOrganizationDTOList(orgs),
		Total: total,
		Page:  query.Page,
		Limit: query.Limit,
	}, nil
}
