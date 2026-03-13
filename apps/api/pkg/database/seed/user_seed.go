package seed

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type systemUser struct {
	name     string
	email    string
	password string
}

// systemUsers are the minimum required users for the system to operate.
// Passwords must be overridden via environment variables in production.
var systemUsers = []systemUser{
	{"System", "system@internal", "system-placeholder-not-for-login"},
}

func seedUsers(db *sql.DB) error {
	for _, u := range systemUsers {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		now := time.Now().UTC()
		_, err = db.Exec(`
			INSERT INTO users (id, email, password_hash, name, is_active, created_at, updated_at)
			VALUES ($1, $2, $3, $4, true, $5, $5)
			ON CONFLICT (email) DO NOTHING`,
			uuid.New(), u.email, string(hash), u.name, now,
		)
		if err != nil {
			return err
		}
		log.Printf("  seed: ensured user %s", u.email)
	}
	return nil
}
