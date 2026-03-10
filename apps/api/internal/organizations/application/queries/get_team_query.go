package queries

import (
	"errors"

	"github.com/google/uuid"
)

type GetTeamQuery struct {
	TeamID uuid.UUID
}

func (q GetTeamQuery) Validate() error {
	if q.TeamID == uuid.Nil {
		return errors.New("team ID is required")
	}
	return nil
}

type ListTeamsQuery struct {
	OrgID uuid.UUID
}

func (q ListTeamsQuery) Validate() error {
	if q.OrgID == uuid.Nil {
		return errors.New("organization ID is required")
	}
	return nil
}

type ListMembersQuery struct {
	OrgID uuid.UUID
	Page  int
	Limit int
}

func (q ListMembersQuery) Validate() error {
	if q.OrgID == uuid.Nil {
		return errors.New("organization ID is required")
	}
	if q.Page < 1 {
		return errors.New("page must be greater than 0")
	}
	if q.Limit < 1 || q.Limit > 100 {
		return errors.New("limit must be between 1 and 100")
	}
	return nil
}
