package commands

import (
	"errors"

	"github.com/google/uuid"
)

type DeleteOrganizationCommand struct {
	OrgID     uuid.UUID
	DeletedBy uuid.UUID
}

func (c DeleteOrganizationCommand) Validate() error {
	if c.OrgID == uuid.Nil {
		return errors.New("organization ID is required")
	}
	if c.DeletedBy == uuid.Nil {
		return errors.New("deleter ID is required")
	}

	return nil
}

type TransferOwnershipCommand struct {
	OrgID          uuid.UUID
	NewOwnerID     uuid.UUID
	CurrentOwnerID uuid.UUID
}

func (c TransferOwnershipCommand) Validate() error {
	if c.OrgID == uuid.Nil {
		return errors.New("organization ID is required")
	}
	if c.NewOwnerID == uuid.Nil {
		return errors.New("new owner ID is required")
	}
	if c.CurrentOwnerID == uuid.Nil {
		return errors.New("current owner ID is required")
	}
	if c.NewOwnerID == c.CurrentOwnerID {
		return errors.New("new owner must be different from current owner")
	}
	return nil
}
