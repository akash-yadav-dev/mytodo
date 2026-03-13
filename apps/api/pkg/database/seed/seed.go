// Package seed provides lightweight database seeding for the application server.
// It is distinct from the tools/seed CLI which does a full development data seed.
// This package is used by the server bootstrap to ensure critical baseline data exists.
package seed

import (
	"database/sql"
	"log"
)

// Run executes all app-level seeders in dependency order.
// All operations are idempotent and safe to run on every startup.
func Run(db *sql.DB) error {
	log.Println("Running database seeds...")

	if err := seedUsers(db); err != nil {
		return err
	}
	return nil
}
