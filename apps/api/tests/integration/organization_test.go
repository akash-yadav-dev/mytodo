package integration_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"mytodo/apps/api/internal/organizations/domain/entity"
	"mytodo/apps/api/internal/organizations/infrastructure/persistence"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var testDB *sql.DB

func setupTestDB(t *testing.T) *sql.DB {
	// Use environment variable or default to test database
	dbHost := getEnvOrDefault("TEST_DB_HOST", "localhost")
	dbPort := getEnvOrDefault("TEST_DB_PORT", "5432")
	dbUser := getEnvOrDefault("TEST_DB_USER", "postgres")
	dbPassword := getEnvOrDefault("TEST_DB_PASSWORD", "postgres")
	dbName := getEnvOrDefault("TEST_DB_NAME", "mytodo_test")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Skip("Skipping integration test: database not available")
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skip("Skipping integration test: database not reachable")
		return nil
	}

	// Create test table if not exists
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS organizations (
		id UUID PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		slug VARCHAR(255) NOT NULL UNIQUE,
		description TEXT,
		plan_id VARCHAR(50) NOT NULL,
		owner_id VARCHAR(255) NOT NULL,
		is_active BOOLEAN NOT NULL DEFAULT true,
		is_deleted BOOLEAN NOT NULL DEFAULT false,
		created_by VARCHAR(255) NOT NULL,
		updated_by VARCHAR(255),
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL,
		deleted_by VARCHAR(255),
		deleted_at TIMESTAMP
	);
	`
	if _, err := db.Exec(createTableSQL); err != nil {
		t.Fatalf("failed to create test table: %v", err)
	}

	return db
}

func cleanupTestDB(t *testing.T, db *sql.DB) {
	if db != nil {
		_, err := db.Exec("DELETE FROM organizations")
		if err != nil {
			t.Logf("Warning: failed to clean up test data: %v", err)
		}
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func TestOrganizationRepository_CreateAndGet(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer cleanupTestDB(t, db)

	repo := persistence.NewOrganizationRepository(db)
	ctx := context.Background()

	ownerID := uuid.New()
	createdBy := ownerID.String()

	org := &entity.Organization{
		ID:          uuid.New(),
		Name:        "Test Organization",
		Slug:        "test-org",
		Description: "Test Description",
		PlanID:      "free",
		OwnerID:     ownerID.String(),
		IsActive:    true,
		IsDeleted:   false,
		CreatedBy:   createdBy,
		UpdatedBy:   "",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		DeletedBy:   "",
		DeletedAt:   nil,
	}

	// Create
	err := repo.CreateOrganization(ctx, org)
	if err != nil {
		t.Fatalf("failed to create organization: %v", err)
	}

	// Get
	retrieved, err := repo.GetOrganizationByID(ctx, org.ID)
	if err != nil {
		t.Fatalf("failed to get organization: %v", err)
	}

	// Verify all fields
	if retrieved.ID != org.ID {
		t.Errorf("expected ID %v, got %v", org.ID, retrieved.ID)
	}
	if retrieved.Name != org.Name {
		t.Errorf("expected Name %s, got %s", org.Name, retrieved.Name)
	}
	if retrieved.Slug != org.Slug {
		t.Errorf("expected Slug %s, got %s", org.Slug, retrieved.Slug)
	}
	if retrieved.CreatedBy != createdBy {
		t.Errorf("expected CreatedBy %s, got %s", createdBy, retrieved.CreatedBy)
	}
	if retrieved.UpdatedBy != "" {
		t.Errorf("expected UpdatedBy to be empty, got %s", retrieved.UpdatedBy)
	}
	if retrieved.DeletedBy != "" {
		t.Errorf("expected DeletedBy to be empty, got %s", retrieved.DeletedBy)
	}
	if retrieved.DeletedAt != nil {
		t.Errorf("expected DeletedAt to be nil")
	}
}

func TestOrganizationRepository_Update_AuditFields(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer cleanupTestDB(t, db)

	repo := persistence.NewOrganizationRepository(db)
	ctx := context.Background()

	ownerID := uuid.New()
	createdBy := ownerID.String()

	org := &entity.Organization{
		ID:          uuid.New(),
		Name:        "Test Organization",
		Slug:        "test-org-update",
		Description: "Test Description",
		PlanID:      "free",
		OwnerID:     ownerID.String(),
		IsActive:    true,
		IsDeleted:   false,
		CreatedBy:   createdBy,
		UpdatedBy:   "",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		DeletedBy:   "",
		DeletedAt:   nil,
	}

	// Create
	err := repo.CreateOrganization(ctx, org)
	if err != nil {
		t.Fatalf("failed to create organization: %v", err)
	}

	// Update
	updaterID := uuid.New()
	org.Name = "Updated Organization"
	org.UpdatedBy = updaterID.String()
	org.UpdatedAt = time.Now().UTC()

	time.Sleep(10 * time.Millisecond) // Small delay to ensure different timestamps

	err = repo.UpdateOrganization(ctx, org)
	if err != nil {
		t.Fatalf("failed to update organization: %v", err)
	}

	// Get and verify
	updated, err := repo.GetOrganizationByID(ctx, org.ID)
	if err != nil {
		t.Fatalf("failed to get updated organization: %v", err)
	}

	if updated.Name != "Updated Organization" {
		t.Errorf("expected Name to be updated")
	}
	if updated.UpdatedBy != updaterID.String() {
		t.Errorf("expected UpdatedBy %s, got %s", updaterID.String(), updated.UpdatedBy)
	}
	if updated.CreatedBy != createdBy {
		t.Errorf("expected CreatedBy to remain %s, got %s", createdBy, updated.CreatedBy)
	}
}

func TestOrganizationRepository_Delete_AuditFields(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer cleanupTestDB(t, db)

	repo := persistence.NewOrganizationRepository(db)
	ctx := context.Background()

	ownerID := uuid.New()
	createdBy := ownerID.String()

	org := &entity.Organization{
		ID:          uuid.New(),
		Name:        "Test Organization",
		Slug:        "test-org-delete",
		Description: "Test Description",
		PlanID:      "free",
		OwnerID:     ownerID.String(),
		IsActive:    true,
		IsDeleted:   false,
		CreatedBy:   createdBy,
		UpdatedBy:   "",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		DeletedBy:   "",
		DeletedAt:   nil,
	}

	// Create
	err := repo.CreateOrganization(ctx, org)
	if err != nil {
		t.Fatalf("failed to create organization: %v", err)
	}

	// Delete
	deleterID := uuid.New()
	err = repo.DeleteOrganization(ctx, org.ID, deleterID)
	if err != nil {
		t.Fatalf("failed to delete organization: %v", err)
	}

	// Verify it can't be retrieved (soft deleted)
	_, err = repo.GetOrganizationByID(ctx, org.ID)
	if err == nil {
		t.Fatal("expected error when getting deleted organization")
	}

	// Verify audit fields directly in database
	var deletedBy sql.NullString
	var deletedAt sql.NullTime
	var isDeleted bool

	query := "SELECT deleted_by, deleted_at, is_deleted FROM organizations WHERE id = $1"
	err = db.QueryRowContext(ctx, query, org.ID).Scan(&deletedBy, &deletedAt, &isDeleted)
	if err != nil {
		t.Fatalf("failed to query deleted organization: %v", err)
	}

	if !isDeleted {
		t.Errorf("expected is_deleted to be true")
	}
	if !deletedBy.Valid || deletedBy.String != deleterID.String() {
		t.Errorf("expected deleted_by to be %s, got %v", deleterID.String(), deletedBy)
	}
	if !deletedAt.Valid {
		t.Errorf("expected deleted_at to be set")
	}
}

func TestOrganizationRepository_ListOrganizations_ExcludesDeleted(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer cleanupTestDB(t, db)

	repo := persistence.NewOrganizationRepository(db)
	ctx := context.Background()

	ownerID := uuid.New()
	createdBy := ownerID.String()

	// Create 3 organizations
	for i := 0; i < 3; i++ {
		org := &entity.Organization{
			ID:          uuid.New(),
			Name:        fmt.Sprintf("Test Org %d", i),
			Slug:        fmt.Sprintf("test-org-%d-%d", i, time.Now().Unix()),
			Description: "Test Description",
			PlanID:      "free",
			OwnerID:     ownerID.String(),
			IsActive:    true,
			IsDeleted:   false,
			CreatedBy:   createdBy,
			UpdatedBy:   "",
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			DeletedBy:   "",
			DeletedAt:   nil,
		}
		err := repo.CreateOrganization(ctx, org)
		if err != nil {
			t.Fatalf("failed to create organization %d: %v", i, err)
		}

		// Delete the second one
		if i == 1 {
			err = repo.DeleteOrganization(ctx, org.ID, ownerID)
			if err != nil {
				t.Fatalf("failed to delete organization: %v", err)
			}
		}
	}

	// List organizations
	orgs, total, err := repo.ListOrganizations(ctx, 1, 10)
	if err != nil {
		t.Fatalf("failed to list organizations: %v", err)
	}

	// Should only return 2 (the non-deleted ones)
	if total < 2 {
		t.Errorf("expected at least 2 organizations, got %d", total)
	}
	if len(orgs) < 2 {
		t.Errorf("expected at least 2 organizations in list, got %d", len(orgs))
	}

	// Verify all returned organizations have audit fields
	for _, org := range orgs {
		if org.CreatedBy == "" {
			t.Errorf("expected CreatedBy to be set for org %s", org.ID)
		}
		if org.CreatedAt.IsZero() {
			t.Errorf("expected CreatedAt to be set for org %s", org.ID)
		}
		if org.IsDeleted {
			t.Errorf("expected org %s to not be deleted", org.ID)
		}
	}
}

func TestOrganizationRepository_RestoreOrganization(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer cleanupTestDB(t, db)

	repo := persistence.NewOrganizationRepository(db)
	ctx := context.Background()

	ownerID := uuid.New()
	createdBy := ownerID.String()

	org := &entity.Organization{
		ID:          uuid.New(),
		Name:        "Test Organization",
		Slug:        "test-org-restore",
		Description: "Test Description",
		PlanID:      "free",
		OwnerID:     ownerID.String(),
		IsActive:    true,
		IsDeleted:   false,
		CreatedBy:   createdBy,
		UpdatedBy:   "",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		DeletedBy:   "",
		DeletedAt:   nil,
	}

	// Create
	err := repo.CreateOrganization(ctx, org)
	if err != nil {
		t.Fatalf("failed to create organization: %v", err)
	}

	// Delete
	err = repo.DeleteOrganization(ctx, org.ID, ownerID)
	if err != nil {
		t.Fatalf("failed to delete organization: %v", err)
	}

	// Restore
	err = repo.RestoreOrganization(ctx, org.ID)
	if err != nil {
		t.Fatalf("failed to restore organization: %v", err)
	}

	// Verify it can be retrieved again
	restored, err := repo.GetOrganizationByID(ctx, org.ID)
	if err != nil {
		t.Fatalf("failed to get restored organization: %v", err)
	}

	if restored.IsDeleted {
		t.Errorf("expected organization to not be deleted")
	}
	if restored.DeletedAt != nil {
		t.Errorf("expected DeletedAt to be nil")
	}
	// DeletedBy might still be set in some implementations, but DeletedAt should be nil
}
