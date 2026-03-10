package commands

import (
	"errors"
	"mytodo/apps/api/internal/organizations/domain/entity"

	"github.com/google/uuid"
)

type AddMemberCommand struct {
	OrgID     uuid.UUID
	UserID    uuid.UUID
	Role      entity.Role
	InvitedBy uuid.UUID
}

func (c AddMemberCommand) Validate() error {
	if c.OrgID == uuid.Nil {
		return errors.New("organization ID is required")
	}
	if c.UserID == uuid.Nil {
		return errors.New("user ID is required")
	}
	if c.InvitedBy == uuid.Nil {
		return errors.New("inviter ID is required")
	}
	if !entity.IsValidRole(c.Role) {
		return errors.New("invalid role")
	}
	return nil
}
