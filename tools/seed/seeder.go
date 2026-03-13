package main

import (
	"context"
	"database/sql"
	"log"
	"mytodo/tools/seed/factory"
	"time"
)

// runSeed executes all seed operations in dependency order.
// It is idempotent: running it multiple times will not create duplicates.
func runSeed(ctx context.Context, db *sql.DB, dataDir string) error {
	log.Println("=== Starting database seed ===")
	start := time.Now()

	// 1. Seed users (returns their UUIDs in insertion order).
	userIDs, err := factory.SeedUsers(ctx, db, dataDir)
	if err != nil {
		return err
	}
	if len(userIDs) == 0 {
		log.Println("No users seeded — skipping downstream seeds.")
		return nil
	}

	ownerID := userIDs[0]
	memberIDs := userIDs[1:]

	// 2. Seed organisation + projects.
	if err := factory.SeedOrganizationAndProjects(ctx, db, dataDir, ownerID, memberIDs); err != nil {
		return err
	}

	// 3. Seed issues against the first project.
	orgID, err := factory.FirstOrgID(ctx, db)
	if err != nil {
		return err
	}
	projectID, err := factory.FirstProjectID(ctx, db, orgID)
	if err != nil {
		return err
	}
	assigneeID := ownerID
	if len(memberIDs) > 0 {
		assigneeID = memberIDs[0]
	}
	if err := factory.SeedIssues(ctx, db, projectID, ownerID, assigneeID); err != nil {
		return err
	}

	log.Printf("=== Seed complete in %s ===", time.Since(start).Round(time.Millisecond))
	return nil
}
