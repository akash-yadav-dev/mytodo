package factory

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ProjectSeed struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// SeedOrganizationAndProjects creates a default organization owned by ownerID,
// adds all userIDs as members, and seeds projects from data/projects.json.
func SeedOrganizationAndProjects(ctx context.Context, db *sql.DB, dataDir string, ownerID uuid.UUID, memberIDs []uuid.UUID) error {
	// --- Organization ---
	orgID := uuid.New()
	now := time.Now().UTC()
	orgSlug := "mytodo-default"

	err := db.QueryRowContext(ctx, `
		INSERT INTO organizations (id, name, slug, description, plan_id, owner_id, is_active, is_deleted, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, true, false, $6, $7, $7)
		ON CONFLICT (slug) DO UPDATE SET slug = EXCLUDED.slug
		RETURNING id`,
		orgID, "MyTodo Default Org", orgSlug,
		"Default organization created during seeding",
		"free", ownerID, now,
	).Scan(&orgID)
	if err != nil {
		return fmt.Errorf("insert organization: %w", err)
	}
	log.Printf("  seeded organization: MyTodo Default Org (id=%s)", orgID)

	// --- Org members ---
	for _, uid := range append([]uuid.UUID{ownerID}, memberIDs...) {
		role := "member"
		if uid == ownerID {
			role = "owner"
		}
		_, err := db.ExecContext(ctx, `
			INSERT INTO organization_members (id, organization_id, user_id, role, joined_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $5)
			ON CONFLICT (organization_id, user_id) DO NOTHING`,
			uuid.New(), orgID, uid, role, now,
		)
		if err != nil {
			return fmt.Errorf("insert member %s: %w", uid, err)
		}
	}
	log.Printf("  seeded %d organization members", len(memberIDs)+1)

	// --- Projects ---
	data, err := os.ReadFile(fmt.Sprintf("%s/projects.json", dataDir))
	if err != nil {
		return fmt.Errorf("read projects.json: %w", err)
	}

	var seeds []ProjectSeed
	if err := json.Unmarshal(data, &seeds); err != nil {
		return fmt.Errorf("parse projects.json: %w", err)
	}

	for _, s := range seeds {
		projectKey := strings.ToUpper(strings.TrimSpace(s.Key))
		_, err := db.ExecContext(ctx, `
			INSERT INTO projects (id, organization_id, name, key, description, status, owner_id, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $8)
			ON CONFLICT (key) DO NOTHING`,
			uuid.New(), orgID, s.Name, projectKey, s.Description, s.Status, ownerID, now,
		)
		if err != nil {
			return fmt.Errorf("insert project %s: %w", s.Key, err)
		}
		log.Printf("  seeded project: %s [%s]", s.Name, projectKey)
	}
	return nil
}
