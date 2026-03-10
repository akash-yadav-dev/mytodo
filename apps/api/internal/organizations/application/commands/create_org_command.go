package commands

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type CreateOrganizationCommand struct {
	OwnerID     uuid.UUID
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
	PlanID      string `json:"plan_id,omitempty"`
}

func (c CreateOrganizationCommand) Validate() error {
	if c.OwnerID == uuid.Nil {
		return errors.New("owner ID is required")
	}

	if strings.TrimSpace(c.Name) == "" {
		return errors.New("organization name is required")
	}

	if len(c.Name) < 2 || len(c.Name) > 120 {
		return errors.New("organization name must be between 2 and 120 characters")
	}

	return nil
}
