package queries

import (
	"errors"

	"github.com/google/uuid"
)

type ListOrganizationsQuery struct {
	Page  int
	Limit int
}

func (q ListOrganizationsQuery) Validate() error {
	if q.Page < 1 {
		return errors.New("page must be greater than 0")
	}
	if q.Limit < 1 || q.Limit > 100 {
		return errors.New("limit must be between 1 and 100")
	}
	return nil
}

type SearchOrganizationsQuery struct {
	Query string
	Limit int
}

func (q SearchOrganizationsQuery) Validate() error {
	if q.Query == "" {
		return errors.New("search query is required")
	}
	if q.Limit < 1 || q.Limit > 100 {
		return errors.New("limit must be between 1 and 100")
	}
	return nil
}

type GetMemberOrganizationsQuery struct {
	UserID uuid.UUID
	Page   int
	Limit  int
}

func (q GetMemberOrganizationsQuery) Validate() error {
	if q.UserID == uuid.Nil {
		return errors.New("user ID is required")
	}
	if q.Page < 1 {
		return errors.New("page must be greater than 0")
	}
	if q.Limit < 1 || q.Limit > 100 {
		return errors.New("limit must be between 1 and 100")
	}
	return nil
}
