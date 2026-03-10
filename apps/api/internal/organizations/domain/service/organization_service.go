package service

import (
	"context"
	"errors"
	"fmt"
	"mytodo/apps/api/internal/organizations/domain/entity"
	"mytodo/apps/api/internal/organizations/domain/repository"
	"strings"

	"github.com/google/uuid"
)

type OrganizationService interface {
	// Define methods for organization-related operations
	CreateOrganization(ctx context.Context, ownerID uuid.UUID, name, slug, description, planID string) (*entity.Organization, error)
	GetOrganizationByID(ctx context.Context, id uuid.UUID) (*entity.Organization, error)
	UpdateOrganization(ctx context.Context, id uuid.UUID, name, slug, description, planID string) (*entity.Organization, error)
	DeleteOrganization(ctx context.Context, id uuid.UUID) error
	RestoreOrganization(ctx context.Context, id uuid.UUID) error
	ListOrganizations(ctx context.Context, page, limit int) ([]*entity.Organization, int, error)
	SearchOrganizations(ctx context.Context, query string, limit int) ([]*entity.Organization, error)
	TransferOrganizationOwnership(ctx context.Context, orgID, newOwnerID uuid.UUID) error
	GetOrganizationsByOwner(ctx context.Context, ownerID uuid.UUID) ([]*entity.Organization, error)
	GetMemberOrganizations(ctx context.Context, userID uuid.UUID, page, limit int) ([]*entity.Organization, int, error)
	IsOwnedBy(ctx context.Context, orgID, userID uuid.UUID) (bool, error)
}

type OrganizationServiceImpl struct {
	organizationRepo repository.OrganizationRepository
}

func NewOrganizationService(organizationRepo repository.OrganizationRepository) OrganizationService {
	return &OrganizationServiceImpl{
		organizationRepo: organizationRepo,
	}
}

func (s *OrganizationServiceImpl) CreateOrganization(ctx context.Context, ownerID uuid.UUID, name, slug, description, planID string) (*entity.Organization, error) {
	// Validate owner ID
	if ownerID == uuid.Nil {
		return nil, errors.New("owner ID cannot be empty")
	}

	// Generate slug if not provided
	if slug == "" {
		slug = generateSlug(name)
	}

	// Set default plan if not provided
	if planID == "" {
		planID = "free" // Default plan
	}

	// Create organization entity
	org, err := entity.NewOrganization(ownerID.String(), name, slug, description, planID)
	if err != nil {
		return nil, fmt.Errorf("failed to create organization entity: %w", err)
	}

	// Save to repository
	if err := s.organizationRepo.CreateOrganization(ctx, org); err != nil {
		return nil, fmt.Errorf("failed to save organization: %w", err)
	}

	return org, nil
}

func (s *OrganizationServiceImpl) GetOrganizationByID(ctx context.Context, id uuid.UUID) (*entity.Organization, error) {
	if id == uuid.Nil {
		return nil, errors.New("organization ID cannot be empty")
	}

	return s.organizationRepo.GetOrganizationByID(ctx, id)
}

func (s *OrganizationServiceImpl) UpdateOrganization(ctx context.Context, id uuid.UUID, name, slug, description, planID string) (*entity.Organization, error) {
	if id == uuid.Nil {
		return nil, errors.New("organization ID cannot be empty")
	}

	// Get existing organization
	org, err := s.organizationRepo.GetOrganizationByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization: %w", err)
	}

	// Update fields if provided
	if name != "" {
		if err := org.UpdateName(name); err != nil {
			return nil, err
		}
	}

	if slug != "" {
		org.Slug = slug
	}

	if description != "" {
		org.Description = description
	}

	if planID != "" {
		org.PlanID = planID
	}

	// Save changes
	if err := s.organizationRepo.UpdateOrganization(ctx, org); err != nil {
		return nil, fmt.Errorf("failed to update organization: %w", err)
	}

	return org, nil
}

func (s *OrganizationServiceImpl) DeleteOrganization(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("organization ID cannot be empty")
	}

	return s.organizationRepo.DeleteOrganization(ctx, id)
}

func (s *OrganizationServiceImpl) RestoreOrganization(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("organization ID cannot be empty")
	}

	return s.organizationRepo.RestoreOrganization(ctx, id)
}

func (s *OrganizationServiceImpl) ListOrganizations(ctx context.Context, page, limit int) ([]*entity.Organization, int, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.organizationRepo.ListOrganizations(ctx, page, limit)
}

func (s *OrganizationServiceImpl) SearchOrganizations(ctx context.Context, query string, limit int) ([]*entity.Organization, error) {
	if query == "" {
		return []*entity.Organization{}, nil
	}

	if limit < 1 || limit > 100 {
		limit = 20
	}

	return s.organizationRepo.SearchOrganizations(ctx, query, limit)
}

func (s *OrganizationServiceImpl) TransferOrganizationOwnership(ctx context.Context, orgID, newOwnerID uuid.UUID) error {
	if orgID == uuid.Nil || newOwnerID == uuid.Nil {
		return errors.New("organization ID and new owner ID cannot be empty")
	}

	return s.organizationRepo.TransferOrganizationOwnership(ctx, orgID, newOwnerID)
}

func (s *OrganizationServiceImpl) GetOrganizationsByOwner(ctx context.Context, ownerID uuid.UUID) ([]*entity.Organization, error) {
	if ownerID == uuid.Nil {
		return nil, errors.New("owner ID cannot be empty")
	}

	return s.organizationRepo.GetOrganizationsByOwner(ctx, ownerID)
}

func (s *OrganizationServiceImpl) GetMemberOrganizations(ctx context.Context, userID uuid.UUID, page, limit int) ([]*entity.Organization, int, error) {
	if userID == uuid.Nil {
		return nil, 0, errors.New("user ID cannot be empty")
	}

	if page < 1 {
		page = 1
	}

	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.organizationRepo.GetMemberOrganizations(ctx, userID, page, limit)
}

func (s *OrganizationServiceImpl) IsOwnedBy(ctx context.Context, orgID, userID uuid.UUID) (bool, error) {
	if orgID == uuid.Nil || userID == uuid.Nil {
		return false, errors.New("organization ID and user ID cannot be empty")
	}

	return s.organizationRepo.IsOwnedBy(ctx, orgID, userID)
}

// Helper function to generate slug from name
func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			return r
		}
		return -1
	}, slug)
	return slug
}
