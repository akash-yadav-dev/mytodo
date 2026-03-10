package commands

import (
	"errors"

	"github.com/google/uuid"
)

type RemoveMemberCommand struct {
	OrgID    uuid.UUID
	UserID   uuid.UUID
	RemovedBy uuid.UUID
}

func (c RemoveMemberCommand) Validate() error {
	if c.OrgID == uuid.Nil {
		return errors.New("organization ID is required")
	}
	if c.UserID == uuid.Nil {
		return errors.New("user ID is required")
	}
	if c.RemovedBy == uuid.Nil {
		return errors.New("remover ID is required")
	}
	return nil
}

type UpdateMemberRoleCommand struct {
	OrgID     uuid.UUID
	UserID    uuid.UUID
	NewRole   string
	UpdatedBy uuid.UUID
}

func (c UpdateMemberRoleCommand) Validate() error {
	if c.OrgID == uuid.Nil {
		return errors.New("organization ID is required")
	}
	if c.UserID == uuid.Nil {
		return errors.New("user ID is required")
	}
	if c.UpdatedBy == uuid.Nil {
		return errors.New("updater ID is required")
	}
	if c.NewRole == "" {
		return errors.New("new role is required")
	}
	return nil
}
