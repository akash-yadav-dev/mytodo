package commands

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type UpdateOrganizationCommand struct {
	OrgID       uuid.UUID
	Name        *string
	Slug        *string
	Description *string
	PlanID      *string
	UpdatedBy   uuid.UUID
}

func (c UpdateOrganizationCommand) Validate() error {
	if c.OrgID == uuid.Nil {
		return errors.New("organization ID is required")
	}

	if c.Name != nil {
		trimmedName := strings.TrimSpace(*c.Name)
		if len(trimmedName) < 2 || len(trimmedName) > 120 {
			return errors.New("name must be between 2 and 120 characters")
		}
	}

	if c.Description != nil && len(*c.Description) > 500 {
		return errors.New("description must not exceed 500 characters")
	}

	if c.UpdatedBy == uuid.Nil {
		return errors.New("updater ID is required")
	}

	return nil
}
