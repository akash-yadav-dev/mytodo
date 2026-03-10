package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidOrganizationName = errors.New("organization name cannot be empty")
	ErrNotOrganizationOwner    = errors.New("user is not the owner of the organization")
)

type Organization struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Description string     `json:"description"`
	PlanID      string     `json:"plan_id"`
	OwnerID     string     `json:"owner_id"`
	IsActive    bool       `json:"is_active"`
	IsDeleted   bool       `json:"is_deleted"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func NewOrganization(ownerID, name, slug, description, planID string) (*Organization, error) {

	org := &Organization{
		ID:          uuid.New(),
		Name:        name,
		Slug:        slug,
		Description: description,
		OwnerID:     ownerID,
		PlanID:      planID,
		IsActive:    true,
		IsDeleted:   false,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	if err := org.validate(); err != nil {
		return nil, err
	}

	return org, nil
}

func (o *Organization) validate() error {
	if strings.TrimSpace(o.Name) == "" {
		return ErrInvalidOrganizationName
	}
	return nil
}

func (o *Organization) IsOwnedBy(userID string) bool {
	return o.OwnerID == userID
}

func (o *Organization) UpdateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return ErrInvalidOrganizationName
	}

	o.Name = name
	o.UpdatedAt = time.Now().UTC()
	return nil
}

func (o *Organization) Deactivate() {
	o.IsActive = false
	o.UpdatedAt = time.Now().UTC()
}

func (o *Organization) SoftDelete() {
	now := time.Now().UTC()
	o.IsDeleted = true
	o.DeletedAt = &now
	o.UpdatedAt = now
}
