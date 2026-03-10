package handlers

import (
	"context"
	"fmt"
	"mytodo/apps/api/internal/organizations/application/commands"
	"mytodo/apps/api/internal/organizations/application/queries"
	"mytodo/apps/api/internal/organizations/domain/entity"
	"mytodo/apps/api/internal/organizations/domain/repository"
	"mytodo/apps/api/internal/organizations/interfaces/dto"
)

type MemberCommandHandler struct {
	membershipRepo repository.MembershipRepository
	orgRepo        repository.OrganizationRepository
}

func NewMemberCommandHandler(
	membershipRepo repository.MembershipRepository,
	orgRepo repository.OrganizationRepository,
) *MemberCommandHandler {
	return &MemberCommandHandler{
		membershipRepo: membershipRepo,
		orgRepo:        orgRepo,
	}
}

func (h *MemberCommandHandler) HandleAddMember(ctx context.Context, cmd commands.AddMemberCommand) (*dto.MemberDTO, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	// Check if user is already a member
	isMember, err := h.membershipRepo.IsMember(ctx, cmd.OrgID, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to check membership: %w", err)
	}
	if isMember {
		return nil, entity.ErrAlreadyMember
	}

	// Verify the inviter has permission (at least admin)
	inviterRole, err := h.membershipRepo.GetMemberRole(ctx, cmd.OrgID, cmd.InvitedBy)
	if err != nil {
		// Check if inviter is the owner
		isOwner, ownerErr := h.orgRepo.IsOwnedBy(ctx, cmd.OrgID, cmd.InvitedBy)
		if ownerErr != nil || !isOwner {
			return nil, fmt.Errorf("inviter must be an admin or owner")
		}
	} else {
		// Check if inviter has admin or higher permission
		if inviterRole != entity.RoleAdmin && inviterRole != entity.RoleOwner {
			return nil, fmt.Errorf("inviter must be an admin or owner")
		}
	}

	// Create new member
	member, err := entity.NewOrganizationMember(cmd.OrgID, cmd.UserID, cmd.Role, &cmd.InvitedBy)
	if err != nil {
		return nil, fmt.Errorf("failed to create member entity: %w", err)
	}

	// Save to repository
	if err := h.membershipRepo.AddMember(ctx, member); err != nil {
		return nil, fmt.Errorf("failed to add member: %w", err)
	}

	return dto.ToMemberDTO(member), nil
}

func (h *MemberCommandHandler) HandleRemoveMember(ctx context.Context, cmd commands.RemoveMemberCommand) error {
	if err := cmd.Validate(); err != nil {
		return err
	}

	// Verify the remover has permission
	removerRole, err := h.membershipRepo.GetMemberRole(ctx, cmd.OrgID, cmd.RemovedBy)
	if err != nil {
		// Check if remover is the owner
		isOwner, ownerErr := h.orgRepo.IsOwnedBy(ctx, cmd.OrgID, cmd.RemovedBy)
		if ownerErr != nil || !isOwner {
			return fmt.Errorf("remover must be an admin or owner")
		}
	} else {
		// Check if remover has admin or higher permission
		if removerRole != entity.RoleAdmin && removerRole != entity.RoleOwner {
			return fmt.Errorf("remover must be an admin or owner")
		}
	}

	// Cannot remove the owner
	isOwner, err := h.orgRepo.IsOwnedBy(ctx, cmd.OrgID, cmd.UserID)
	if err != nil {
		return fmt.Errorf("failed to check ownership: %w", err)
	}
	if isOwner {
		return fmt.Errorf("cannot remove the organization owner")
	}

	// Remove member
	if err := h.membershipRepo.RemoveMember(ctx, cmd.OrgID, cmd.UserID); err != nil {
		return fmt.Errorf("failed to remove member: %w", err)
	}

	return nil
}

func (h *MemberCommandHandler) HandleUpdateMemberRole(ctx context.Context, cmd commands.UpdateMemberRoleCommand) (*dto.MemberDTO, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	// Verify the updater has permission
	updaterRole, err := h.membershipRepo.GetMemberRole(ctx, cmd.OrgID, cmd.UpdatedBy)
	if err != nil {
		// Check if updater is the owner
		isOwner, ownerErr := h.orgRepo.IsOwnedBy(ctx, cmd.OrgID, cmd.UpdatedBy)
		if ownerErr != nil || !isOwner {
			return nil, fmt.Errorf("updater must be an admin or owner")
		}
	} else {
		// Check if updater has admin or higher permission
		if updaterRole != entity.RoleAdmin && updaterRole != entity.RoleOwner {
			return nil, fmt.Errorf("updater must be an admin or owner")
		}
	}

	// Cannot change the owner's role
	isOwner, err := h.orgRepo.IsOwnedBy(ctx, cmd.OrgID, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to check ownership: %w", err)
	}
	if isOwner {
		return nil, fmt.Errorf("cannot change the organization owner's role")
	}

	// Update role
	newRole := entity.Role(cmd.NewRole)
	if !entity.IsValidRole(newRole) {
		return nil, entity.ErrInvalidRole
	}

	if err := h.membershipRepo.UpdateMemberRole(ctx, cmd.OrgID, cmd.UserID, newRole); err != nil {
		return nil, fmt.Errorf("failed to update member role: %w", err)
	}

	// Get updated member
	member, err := h.membershipRepo.GetMember(ctx, cmd.OrgID, cmd.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get updated member: %w", err)
	}

	return dto.ToMemberDTO(member), nil
}

func (h *MemberCommandHandler) HandleListMembers(ctx context.Context, query queries.ListMembersQuery) (*dto.MemberListResponse, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}

	members, total, err := h.membershipRepo.ListMembers(ctx, query.OrgID, query.Page, query.Limit)
	if err != nil {
		return nil, fmt.Errorf("failed to list members: %w", err)
	}

	return &dto.MemberListResponse{
		Data:  dto.ToMemberDTOList(members),
		Total: total,
		Page:  query.Page,
		Limit: query.Limit,
	}, nil
}
