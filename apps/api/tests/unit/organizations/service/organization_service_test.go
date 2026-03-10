package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"mytodo/apps/api/internal/organizations/domain/entity"
	"mytodo/apps/api/internal/organizations/domain/service"

	"github.com/google/uuid"
)

type fakeOrgRepo struct {
	orgsByID    map[uuid.UUID]*entity.Organization
	orgsByOwner map[string][]*entity.Organization
	nextID      uuid.UUID
}

func newFakeOrgRepo() *fakeOrgRepo {
	return &fakeOrgRepo{
		orgsByID:    map[uuid.UUID]*entity.Organization{},
		orgsByOwner: map[string][]*entity.Organization{},
	}
}

func (r *fakeOrgRepo) CreateOrganization(ctx context.Context, org *entity.Organization) error {
	r.orgsByID[org.ID] = org
	r.orgsByOwner[org.OwnerID] = append(r.orgsByOwner[org.OwnerID], org)
	return nil
}

func (r *fakeOrgRepo) GetOrganizationByID(ctx context.Context, id uuid.UUID) (*entity.Organization, error) {
	org, ok := r.orgsByID[id]
	if !ok || org.IsDeleted {
		return nil, errors.New("organization not found")
	}
	return org, nil
}

func (r *fakeOrgRepo) UpdateOrganization(ctx context.Context, org *entity.Organization) error {
	if _, ok := r.orgsByID[org.ID]; !ok {
		return errors.New("organization not found")
	}
	r.orgsByID[org.ID] = org
	return nil
}

func (r *fakeOrgRepo) DeleteOrganization(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) error {
	org, ok := r.orgsByID[id]
	if !ok {
		return errors.New("organization not found")
	}
	now := time.Now().UTC()
	org.IsDeleted = true
	org.DeletedAt = &now
	org.DeletedBy = deletedBy.String()
	org.UpdatedAt = now
	return nil
}

func (r *fakeOrgRepo) RestoreOrganization(ctx context.Context, id uuid.UUID) error {
	org, ok := r.orgsByID[id]
	if !ok {
		return errors.New("organization not found")
	}
	org.IsDeleted = false
	org.DeletedAt = nil
	org.DeletedBy = ""
	org.UpdatedAt = time.Now().UTC()
	return nil
}

func (r *fakeOrgRepo) ListOrganizations(ctx context.Context, page, limit int) ([]*entity.Organization, int, error) {
	orgs := []*entity.Organization{}
	for _, org := range r.orgsByID {
		if !org.IsDeleted {
			orgs = append(orgs, org)
		}
	}
	return orgs, len(orgs), nil
}

func (r *fakeOrgRepo) SearchOrganizations(ctx context.Context, query string, limit int) ([]*entity.Organization, error) {
	orgs := []*entity.Organization{}
	for _, org := range r.orgsByID {
		if !org.IsDeleted {
			orgs = append(orgs, org)
		}
	}
	return orgs, nil
}

func (r *fakeOrgRepo) TransferOrganizationOwnership(ctx context.Context, orgID, newOwnerID uuid.UUID) error {
	org, ok := r.orgsByID[orgID]
	if !ok {
		return errors.New("organization not found")
	}
	org.OwnerID = newOwnerID.String()
	org.UpdatedAt = time.Now().UTC()
	return nil
}

func (r *fakeOrgRepo) GetOrganizationsByOwner(ctx context.Context, ownerID uuid.UUID) ([]*entity.Organization, error) {
	orgs := r.orgsByOwner[ownerID.String()]
	result := []*entity.Organization{}
	for _, org := range orgs {
		if !org.IsDeleted {
			result = append(result, org)
		}
	}
	return result, nil
}

func (r *fakeOrgRepo) GetMemberOrganizations(ctx context.Context, userID uuid.UUID, page, limit int) ([]*entity.Organization, int, error) {
	return []*entity.Organization{}, 0, nil
}

func (r *fakeOrgRepo) IsOwnedBy(ctx context.Context, orgID, userID uuid.UUID) (bool, error) {
	org, ok := r.orgsByID[orgID]
	if !ok {
		return false, errors.New("organization not found")
	}
	return org.OwnerID == userID.String(), nil
}

func TestOrganizationService_CreateOrganization(t *testing.T) {
	repo := newFakeOrgRepo()
	svc := service.NewOrganizationService(repo)

	ownerID := uuid.New()
	createdBy := ownerID.String()

	org, err := svc.CreateOrganization(context.Background(), ownerID, "Test Org", "", "Test Description", "free", createdBy)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if org.Name != "Test Org" {
		t.Errorf("expected name 'Test Org', got %s", org.Name)
	}

	if org.OwnerID != ownerID.String() {
		t.Errorf("expected ownerID to match")
	}

	if org.CreatedBy != createdBy {
		t.Errorf("expected createdBy to be %s, got %s", createdBy, org.CreatedBy)
	}

	if org.UpdatedBy != "" {
		t.Errorf("expected updatedBy to be empty on creation, got %s", org.UpdatedBy)
	}

	if org.DeletedBy != "" {
		t.Errorf("expected deletedBy to be empty on creation, got %s", org.DeletedBy)
	}

	if org.DeletedAt != nil {
		t.Errorf("expected deletedAt to be nil on creation")
	}

	if org.CreatedAt.IsZero() {
		t.Errorf("expected createdAt to be set")
	}

	if org.UpdatedAt.IsZero() {
		t.Errorf("expected updatedAt to be set")
	}

	if !org.IsActive {
		t.Errorf("expected organization to be active")
	}

	if org.IsDeleted {
		t.Errorf("expected organization to not be deleted")
	}
}

func TestOrganizationService_CreateOrganization_WithoutCreatedBy(t *testing.T) {
	repo := newFakeOrgRepo()
	svc := service.NewOrganizationService(repo)

	ownerID := uuid.New()

	_, err := svc.CreateOrganization(context.Background(), ownerID, "Test Org", "", "Test Description", "free", "")
	if err == nil {
		t.Fatal("expected error when createdBy is empty")
	}

	if err.Error() != "created by cannot be empty" {
		t.Errorf("expected specific error message, got: %v", err)
	}
}

func TestOrganizationService_UpdateOrganization(t *testing.T) {
	repo := newFakeOrgRepo()
	svc := service.NewOrganizationService(repo)

	ownerID := uuid.New()
	createdBy := ownerID.String()

	// Create organization
	org, err := svc.CreateOrganization(context.Background(), ownerID, "Test Org", "", "Test Description", "free", createdBy)
	if err != nil {
		t.Fatalf("failed to create org: %v", err)
	}

	updaterID := uuid.New()
	updatedBy := updaterID.String()

	// Small delay to ensure different timestamps
	time.Sleep(10 * time.Millisecond)

	// Update organization
	updatedOrg, err := svc.UpdateOrganization(context.Background(), org.ID, "Updated Org", "", "Updated Description", "", updatedBy)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if updatedOrg.Name != "Updated Org" {
		t.Errorf("expected name 'Updated Org', got %s", updatedOrg.Name)
	}

	if updatedOrg.Description != "Updated Description" {
		t.Errorf("expected description to be updated")
	}

	if updatedOrg.UpdatedBy != updatedBy {
		t.Errorf("expected updatedBy to be %s, got %s", updatedBy, updatedOrg.UpdatedBy)
	}

	if updatedOrg.CreatedBy != createdBy {
		t.Errorf("expected createdBy to remain %s, got %s", createdBy, updatedOrg.CreatedBy)
	}

	if updatedOrg.UpdatedAt.Before(org.UpdatedAt) {
		t.Errorf("expected updatedAt to not be before original")
	}
}

func TestOrganizationService_UpdateOrganization_WithoutUpdatedBy(t *testing.T) {
	repo := newFakeOrgRepo()
	svc := service.NewOrganizationService(repo)

	ownerID := uuid.New()
	createdBy := ownerID.String()

	org, err := svc.CreateOrganization(context.Background(), ownerID, "Test Org", "", "Test Description", "free", createdBy)
	if err != nil {
		t.Fatalf("failed to create org: %v", err)
	}

	_, err = svc.UpdateOrganization(context.Background(), org.ID, "Updated Org", "", "", "", "")
	if err == nil {
		t.Fatal("expected error when updatedBy is empty")
	}

	if err.Error() != "updated by cannot be empty" {
		t.Errorf("expected specific error message, got: %v", err)
	}
}

func TestOrganizationService_DeleteOrganization(t *testing.T) {
	repo := newFakeOrgRepo()
	svc := service.NewOrganizationService(repo)

	ownerID := uuid.New()
	createdBy := ownerID.String()

	// Create organization
	org, err := svc.CreateOrganization(context.Background(), ownerID, "Test Org", "", "Test Description", "free", createdBy)
	if err != nil {
		t.Fatalf("failed to create org: %v", err)
	}

	deleterID := uuid.New()

	// Delete organization
	err = svc.DeleteOrganization(context.Background(), org.ID, deleterID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Try to get the deleted organization
	_, err = svc.GetOrganizationByID(context.Background(), org.ID)
	if err == nil {
		t.Fatal("expected error when getting deleted organization")
	}

	// Verify it's marked as deleted in the repo
	orgFromRepo := repo.orgsByID[org.ID]
	if !orgFromRepo.IsDeleted {
		t.Errorf("expected organization to be marked as deleted")
	}

	if orgFromRepo.DeletedBy != deleterID.String() {
		t.Errorf("expected deletedBy to be %s, got %s", deleterID.String(), orgFromRepo.DeletedBy)
	}

	if orgFromRepo.DeletedAt == nil {
		t.Errorf("expected deletedAt to be set")
	}
}

func TestOrganizationService_DeleteOrganization_WithoutDeletedBy(t *testing.T) {
	repo := newFakeOrgRepo()
	svc := service.NewOrganizationService(repo)

	ownerID := uuid.New()
	createdBy := ownerID.String()

	org, err := svc.CreateOrganization(context.Background(), ownerID, "Test Org", "", "Test Description", "free", createdBy)
	if err != nil {
		t.Fatalf("failed to create org: %v", err)
	}

	err = svc.DeleteOrganization(context.Background(), org.ID, uuid.Nil)
	if err == nil {
		t.Fatal("expected error when deletedBy is empty")
	}

	if err.Error() != "deleted by cannot be empty" {
		t.Errorf("expected specific error message, got: %v", err)
	}
}

func TestOrganizationService_GetOrganizationByID(t *testing.T) {
	repo := newFakeOrgRepo()
	svc := service.NewOrganizationService(repo)

	ownerID := uuid.New()
	createdBy := ownerID.String()

	org, err := svc.CreateOrganization(context.Background(), ownerID, "Test Org", "", "Test Description", "free", createdBy)
	if err != nil {
		t.Fatalf("failed to create org: %v", err)
	}

	retrievedOrg, err := svc.GetOrganizationByID(context.Background(), org.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if retrievedOrg.ID != org.ID {
		t.Errorf("expected IDs to match")
	}

	if retrievedOrg.Name != org.Name {
		t.Errorf("expected names to match")
	}

	if retrievedOrg.CreatedBy != org.CreatedBy {
		t.Errorf("expected createdBy to match")
	}
}

func TestOrganizationService_GetOrganizationsByOwner(t *testing.T) {
	repo := newFakeOrgRepo()
	svc := service.NewOrganizationService(repo)

	ownerID := uuid.New()
	createdBy := ownerID.String()

	// Create multiple organizations
	_, err := svc.CreateOrganization(context.Background(), ownerID, "Org 1", "", "Description 1", "free", createdBy)
	if err != nil {
		t.Fatalf("failed to create org 1: %v", err)
	}

	_, err = svc.CreateOrganization(context.Background(), ownerID, "Org 2", "", "Description 2", "premium", createdBy)
	if err != nil {
		t.Fatalf("failed to create org 2: %v", err)
	}

	orgs, err := svc.GetOrganizationsByOwner(context.Background(), ownerID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(orgs) != 2 {
		t.Errorf("expected 2 organizations, got %d", len(orgs))
	}
}

func TestOrganizationService_RestoreOrganization(t *testing.T) {
	repo := newFakeOrgRepo()
	svc := service.NewOrganizationService(repo)

	ownerID := uuid.New()
	createdBy := ownerID.String()

	org, err := svc.CreateOrganization(context.Background(), ownerID, "Test Org", "", "Test Description", "free", createdBy)
	if err != nil {
		t.Fatalf("failed to create org: %v", err)
	}

	// Delete the organization
	err = svc.DeleteOrganization(context.Background(), org.ID, ownerID)
	if err != nil {
		t.Fatalf("failed to delete org: %v", err)
	}

	// Restore the organization
	err = svc.RestoreOrganization(context.Background(), org.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify it's restored
	restoredOrg, err := svc.GetOrganizationByID(context.Background(), org.ID)
	if err != nil {
		t.Fatalf("expected no error when getting restored organization, got %v", err)
	}

	if restoredOrg.IsDeleted {
		t.Errorf("expected organization to not be deleted")
	}

	if restoredOrg.DeletedAt != nil {
		t.Errorf("expected deletedAt to be nil")
	}
}

func TestOrganizationService_AuditFieldsImmutability(t *testing.T) {
	repo := newFakeOrgRepo()
	svc := service.NewOrganizationService(repo)

	ownerID := uuid.New()
	createdBy := ownerID.String()

	org, err := svc.CreateOrganization(context.Background(), ownerID, "Test Org", "", "Test Description", "free", createdBy)
	if err != nil {
		t.Fatalf("failed to create org: %v", err)
	}

	originalCreatedAt := org.CreatedAt
	originalCreatedBy := org.CreatedBy
	originalUpdatedAt := org.UpdatedAt

	// Update the organization
	updaterID := uuid.New()
	time.Sleep(10 * time.Millisecond) // Small delay to ensure different timestamps

	updatedOrg, err := svc.UpdateOrganization(context.Background(), org.ID, "Updated Org", "", "", "", updaterID.String())
	if err != nil {
		t.Fatalf("failed to update org: %v", err)
	}

	// Verify CreatedAt and CreatedBy are immutable
	if updatedOrg.CreatedAt != originalCreatedAt {
		t.Errorf("expected CreatedAt to remain unchanged")
	}

	if updatedOrg.CreatedBy != originalCreatedBy {
		t.Errorf("expected CreatedBy to remain unchanged")
	}

	// Verify UpdatedAt changed
	if !updatedOrg.UpdatedAt.After(originalUpdatedAt) {
		t.Errorf("expected UpdatedAt to change")
	}

	// Verify UpdatedBy was set
	if updatedOrg.UpdatedBy != updaterID.String() {
		t.Errorf("expected UpdatedBy to be set to updater ID")
	}
}
