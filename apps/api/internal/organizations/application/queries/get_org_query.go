package queries

import (
	"errors"

	"github.com/google/uuid"
)

type GetOrganizationQuery struct {
	OrgID uuid.UUID
}

func (q GetOrganizationQuery) Validate() error {
	if q.OrgID == uuid.Nil {
		return errors.New("organization ID is required")
	}
	return nil
}

type GetOrganizationsByOwnerQuery struct {
	OwnerID uuid.UUID
}

func (q GetOrganizationsByOwnerQuery) Validate() error {
	if q.OwnerID == uuid.Nil {
		return errors.New("owner ID is required")
	}
	return nil
}
