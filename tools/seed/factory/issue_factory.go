package factory

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type issueSeed struct {
	title       string
	description string
	status      string
	priority    string
}

var defaultIssues = []issueSeed{
	{"Set up CI/CD pipeline", "Configure GitHub Actions for automated builds and deployments.", "open", "high"},
	{"Implement user authentication", "JWT-based login and registration endpoints.", "open", "high"},
	{"Design database schema", "ERD and initial SQL migrations for all core tables.", "open", "medium"},
	{"Build REST API endpoints", "CRUD operations for projects, issues and organisations.", "open", "medium"},
	{"Write unit tests", "Achieve ≥80 % code coverage across the domain layer.", "open", "medium"},
	{"Configure Docker Compose", "Local development environment with hot-reload.", "open", "low"},
}

// SeedIssues inserts a small set of default issues for projectID.
// reporterID and assigneeID are the seed users.
func SeedIssues(ctx context.Context, db *sql.DB, projectID, reporterID, assigneeID uuid.UUID) error {
	now := time.Now().UTC()
	for _, s := range defaultIssues {
		_, err := db.ExecContext(ctx, `
			INSERT INTO issues
				(id, project_id, title, description, status, priority,
				 assignee_id, reporter_id, created_at, updated_at)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$9)
			ON CONFLICT DO NOTHING`,
			uuid.New(), projectID, s.title, s.description,
			s.status, s.priority, assigneeID, reporterID, now,
		)
		if err != nil {
			return fmt.Errorf("insert issue '%s': %w", s.title, err)
		}
		log.Printf("  seeded issue: %s [%s/%s]", s.title, s.status, s.priority)
	}
	return nil
}

// FirstProjectID returns the UUID of the first project found for the org.
func FirstProjectID(ctx context.Context, db *sql.DB, orgID uuid.UUID) (uuid.UUID, error) {
	var id uuid.UUID
	err := db.QueryRowContext(ctx,
		`SELECT id FROM projects WHERE organization_id = $1 ORDER BY created_at LIMIT 1`,
		orgID,
	).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("first project: %w", err)
	}
	return id, nil
}

// FirstOrgID returns the first organization ID from the database.
func FirstOrgID(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	err := db.QueryRowContext(ctx,
		`SELECT id FROM organizations ORDER BY created_at LIMIT 1`,
	).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("first org: %w", err)
	}
	return id, nil
}
