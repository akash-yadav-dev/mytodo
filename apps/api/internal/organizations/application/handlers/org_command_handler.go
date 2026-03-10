package handlers

import (
	"context"
	"fmt"
	"mytodo/apps/api/internal/organizations/application/commands"
	"mytodo/apps/api/internal/organizations/domain/service"
	"mytodo/apps/api/internal/organizations/interfaces/dto"
)

type OrganizationCommandHandler struct {
	organizationService service.OrganizationService
}

func NewOrganizationCommandHandler(organizationService service.OrganizationService) *OrganizationCommandHandler {
	return &OrganizationCommandHandler{
		organizationService: organizationService,
	}
}

func (h *OrganizationCommandHandler) HandleCreateOrganization(ctx context.Context, cmd commands.CreateOrganizationCommand) (*dto.OrganizationDTO, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	org, err := h.organizationService.CreateOrganization(ctx, cmd.OwnerID, cmd.Name, cmd.Slug, cmd.Description, cmd.PlanID, cmd.CreatedBy.String())
	if err != nil {
		return nil, fmt.Errorf("failed to create organization: %w", err)
	}

	return dto.ToOrganizationDTO(org), nil
}

func (h *OrganizationCommandHandler) HandleUpdateOrganization(ctx context.Context, cmd commands.UpdateOrganizationCommand) (*dto.OrganizationDTO, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	// Extract string values from pointers
	name := ""
	if cmd.Name != nil {
		name = *cmd.Name
	}

	slug := ""
	if cmd.Slug != nil {
		slug = *cmd.Slug
	}

	description := ""
	if cmd.Description != nil {
		description = *cmd.Description
	}

	planID := ""
	if cmd.PlanID != nil {
		planID = *cmd.PlanID
	}

	org, err := h.organizationService.UpdateOrganization(ctx, cmd.OrgID, name, slug, description, planID, cmd.UpdatedBy.String())
	if err != nil {
		return nil, fmt.Errorf("failed to update organization: %w", err)
	}

	return dto.ToOrganizationDTO(org), nil
}

func (h *OrganizationCommandHandler) HandleDeleteOrganization(ctx context.Context, cmd commands.DeleteOrganizationCommand) error {
	if err := cmd.Validate(); err != nil {
		return err
	}

	if err := h.organizationService.DeleteOrganization(ctx, cmd.OrgID, cmd.DeletedBy); err != nil {
		return fmt.Errorf("failed to delete organization: %w", err)
	}

	return nil
}

func (h *OrganizationCommandHandler) HandleTransferOwnership(ctx context.Context, cmd commands.TransferOwnershipCommand) error {
	if err := cmd.Validate(); err != nil {
		return err
	}

	// Verify current user is owner
	isOwner, err := h.organizationService.IsOwnedBy(ctx, cmd.OrgID, cmd.CurrentOwnerID)
	if err != nil {
		return fmt.Errorf("failed to verify ownership: %w", err)
	}

	if !isOwner {
		return fmt.Errorf("user is not the owner of this organization")
	}

	if err := h.organizationService.TransferOrganizationOwnership(ctx, cmd.OrgID, cmd.NewOwnerID); err != nil {
		return fmt.Errorf("failed to transfer ownership: %w", err)
	}

	return nil
}
