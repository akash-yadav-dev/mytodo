package entity

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidOrganizationName = errors.New("organization name cannot be empty")
	ErrInvalidOrganizationSlug = errors.New("organization slug cannot be empty")
	ErrNotOrganizationOwner    = errors.New("user is not the owner of the organization")
)

type Organization struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	PlanID      string    `json:"plan_id"`
	OwnerID     string    `json:"owner_id"`
	IsActive    bool      `json:"is_active"`
	IsDeleted   bool      `json:"is_deleted"`

	// Audit fields
	CreatedBy string     `json:"created_by"`
	UpdatedBy string     `json:"updated_by,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedBy string     `json:"deleted_by,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func NewOrganization(ownerID, name, description, planID, createdBy string) (*Organization, error) {
	// Slug generation belongs in service layer (uniqueness check needed)
	slug := generateSlug(name)

	org := &Organization{
		ID:          uuid.New(),
		Name:        strings.TrimSpace(name),
		Slug:        slug,
		Description: description,
		OwnerID:     ownerID,
		PlanID:      planID,
		IsActive:    true,
		IsDeleted:   false,
		CreatedBy:   createdBy,
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
	if strings.TrimSpace(o.Slug) == "" {
		return ErrInvalidOrganizationSlug
	}
	if o.OwnerID == "" {
		return errors.New("owner ID is required")
	}
	if o.PlanID == "" {
		return errors.New("plan ID is required")
	}
	if o.CreatedBy == "" {
		return errors.New("created by is required")
	}
	return nil
}

func (o *Organization) IsOwnedBy(userID string) bool {
	return o.OwnerID == userID
}

func (o *Organization) CanBeModifiedBy(userID string) bool {
	return o.IsOwnedBy(userID) && o.IsActive && !o.IsDeleted
}

// Business methods - return errors instead of panicking
func (o *Organization) UpdateName(name, updatedBy string) error {
	if strings.TrimSpace(name) == "" {
		return ErrInvalidOrganizationName
	}

	o.Name = strings.TrimSpace(name)
	o.UpdatedBy = updatedBy
	o.UpdatedAt = time.Now().UTC()
	return nil
}

func (o *Organization) UpdateDescription(description, updatedBy string) {
	o.Description = description
	o.UpdatedBy = updatedBy
	o.UpdatedAt = time.Now().UTC()
}

func (o *Organization) UpdatePlanID(planID, updatedBy string) error {
	if planID == "" {
		return errors.New("plan ID cannot be empty")
	}
	o.PlanID = planID
	o.UpdatedBy = updatedBy
	o.UpdatedAt = time.Now().UTC()
	return nil
}

func (o *Organization) UpdateSlug(slug, updatedBy string) error {
	if strings.TrimSpace(slug) == "" {
		return ErrInvalidOrganizationSlug
	}
	o.Slug = strings.TrimSpace(slug)
	o.UpdatedBy = updatedBy
	o.UpdatedAt = time.Now().UTC()
	return nil
}

func (o *Organization) Deactivate(updatedBy string) error {
	if o.IsDeleted {
		return errors.New("cannot deactivate deleted organization")
	}
	o.IsActive = false
	o.UpdatedBy = updatedBy
	o.UpdatedAt = time.Now().UTC()
	return nil
}

func (o *Organization) SoftDelete(deletedBy string) error {
	if o.IsDeleted {
		return errors.New("organization already deleted")
	}

	now := time.Now().UTC()
	o.IsDeleted = true
	o.DeletedAt = &now
	o.DeletedBy = deletedBy
	return nil
}

// Utility function (package-private)
func generateSlug(name string) string {
	slug := strings.ToLower(strings.TrimSpace(name))
	slug = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	if len(slug) > 50 {
		slug = slug[:50]
	}
	return slug
}
