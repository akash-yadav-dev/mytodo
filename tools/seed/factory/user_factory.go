package factory

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserSeed struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SeedUsers inserts seed users from data/users.json.
// Each user also gets a matching user_profile row.
// Conflicts on email are silently skipped (idempotent).
func SeedUsers(ctx context.Context, db *sql.DB, dataDir string) ([]uuid.UUID, error) {
	data, err := os.ReadFile(fmt.Sprintf("%s/users.json", dataDir))
	if err != nil {
		return nil, fmt.Errorf("read users.json: %w", err)
	}

	var seeds []UserSeed
	if err := json.Unmarshal(data, &seeds); err != nil {
		return nil, fmt.Errorf("parse users.json: %w", err)
	}

	var ids []uuid.UUID
	for _, s := range seeds {
		hash, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("hash password for %s: %w", s.Email, err)
		}

		id := uuid.New()
		now := time.Now().UTC()

		// Insert user; skip if email already exists.
		var userID uuid.UUID
		err = db.QueryRowContext(ctx, `
			INSERT INTO users (id, email, password_hash, name, is_active, created_at, updated_at)
			VALUES ($1, $2, $3, $4, true, $5, $5)
			ON CONFLICT (email) DO UPDATE SET email = EXCLUDED.email
			RETURNING id`,
			id, s.Email, string(hash), s.Name, now,
		).Scan(&userID)
		if err != nil {
			return nil, fmt.Errorf("insert user %s: %w", s.Email, err)
		}

		// Upsert profile.
		_, err = db.ExecContext(ctx, `
			INSERT INTO user_profiles (user_id, display_name, created_at, updated_at)
			VALUES ($1, $2, $3, $3)
			ON CONFLICT (user_id) DO NOTHING`,
			userID, s.Name, now,
		)
		if err != nil {
			return nil, fmt.Errorf("insert profile for %s: %w", s.Email, err)
		}

		ids = append(ids, userID)
		log.Printf("  seeded user: %s (%s)", s.Name, s.Email)
	}
	return ids, nil
}
