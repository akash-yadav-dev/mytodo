// Package persistence provides concrete repository implementations for users.
//
// This implements UserRepository interface using PostgreSQL/database.

package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mytodo/apps/api/internal/users/domain/entity"
	"mytodo/apps/api/internal/users/domain/repository"

	"github.com/google/uuid"
)

// UserRepositoryImpl implements UserRepository using PostgreSQL
type UserRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepositoryImpl
func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

// CreateProfile creates a new user profile
func (r *UserRepositoryImpl) CreateProfile(ctx context.Context, user *entity.User) error {
	query := `
		INSERT INTO user_profiles (
			id, user_id, username, display_name, avatar_url, bio, 
			location, website, phone, timezone, language, theme,
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`
	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.UserID, user.Username, user.DisplayName, user.AvatarURL,
		user.Bio, user.Location, user.Website, user.Phone, user.Timezone,
		user.Language, user.Theme, user.CreatedAt, user.UpdatedAt,
	)
	return err
}

// FindProfileByID retrieves a user profile by profile ID
func (r *UserRepositoryImpl) FindProfileByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	query := `
		SELECT id, user_id, username, display_name, avatar_url, bio,
		       location, website, phone, timezone, language, theme,
		       created_at, updated_at
		FROM user_profiles
		WHERE id = $1
	`
	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.UserID, &user.Username, &user.DisplayName, &user.AvatarURL,
		&user.Bio, &user.Location, &user.Website, &user.Phone, &user.Timezone,
		&user.Language, &user.Theme, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("user profile not found")
	}
	return user, err
}

// FindProfileByUserID retrieves a user profile by auth user ID
func (r *UserRepositoryImpl) FindProfileByUserID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	query := `
		SELECT id, user_id, username, display_name, avatar_url, bio,
		       location, website, phone, timezone, language, theme,
		       created_at, updated_at
		FROM user_profiles
		WHERE user_id = $1
	`
	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&user.ID, &user.UserID, &user.Username, &user.DisplayName, &user.AvatarURL,
		&user.Bio, &user.Location, &user.Website, &user.Phone, &user.Timezone,
		&user.Language, &user.Theme, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("user profile not found")
	}
	return user, err
}

// FindProfileByUsername retrieves a user profile by username
func (r *UserRepositoryImpl) FindProfileByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := `
		SELECT id, user_id, username, display_name, avatar_url, bio,
		       location, website, phone, timezone, language, theme,
		       created_at, updated_at
		FROM user_profiles
		WHERE username = $1
	`
	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID, &user.UserID, &user.Username, &user.DisplayName, &user.AvatarURL,
		&user.Bio, &user.Location, &user.Website, &user.Phone, &user.Timezone,
		&user.Language, &user.Theme, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("user profile not found")
	}
	return user, err
}

// UpdateProfile updates a user profile
func (r *UserRepositoryImpl) UpdateProfile(ctx context.Context, user *entity.User) error {
	query := `
		UPDATE user_profiles
		SET username = $1, display_name = $2, avatar_url = $3, bio = $4,
		    location = $5, website = $6, phone = $7, timezone = $8,
		    language = $9, theme = $10, updated_at = $11
		WHERE user_id = $12
	`
	result, err := r.db.ExecContext(ctx, query,
		user.Username, user.DisplayName, user.AvatarURL, user.Bio,
		user.Location, user.Website, user.Phone, user.Timezone,
		user.Language, user.Theme, user.UpdatedAt, user.UserID,
	)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user profile not found")
	}
	return nil
}

// DeleteProfile deletes a user profile
func (r *UserRepositoryImpl) DeleteProfile(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM user_profiles WHERE user_id = $1`
	result, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user profile not found")
	}
	return nil
}

// ListProfiles retrieves a paginated list of user profiles
func (r *UserRepositoryImpl) ListProfiles(ctx context.Context, page, limit int) ([]*entity.User, int, error) {
	// Calculate offset
	offset := (page - 1) * limit

	// Get total count
	var total int
	countQuery := `SELECT COUNT(*) FROM user_profiles`
	if err := r.db.QueryRowContext(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	// Get paginated results
	query := `
		SELECT id, user_id, username, display_name, avatar_url, bio,
		       location, website, phone, timezone, language, theme,
		       created_at, updated_at
		FROM user_profiles
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		user := &entity.User{}
		err := rows.Scan(
			&user.ID, &user.UserID, &user.Username, &user.DisplayName, &user.AvatarURL,
			&user.Bio, &user.Location, &user.Website, &user.Phone, &user.Timezone,
			&user.Language, &user.Theme, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	return users, total, rows.Err()
}

// ExistsByUsername checks if a username already exists
func (r *UserRepositoryImpl) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM user_profiles WHERE username = $1)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, username).Scan(&exists)
	return exists, err
}

// SearchProfiles searches user profiles by display name or username
func (r *UserRepositoryImpl) SearchProfiles(ctx context.Context, query string, limit int) ([]*entity.User, error) {
	searchQuery := `
		SELECT id, user_id, username, display_name, avatar_url, bio,
		       location, website, phone, timezone, language, theme,
		       created_at, updated_at
		FROM user_profiles
		WHERE display_name ILIKE $1 OR username ILIKE $1
		LIMIT $2
	`
	searchPattern := fmt.Sprintf("%%%s%%", query)
	rows, err := r.db.QueryContext(ctx, searchQuery, searchPattern, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		user := &entity.User{}
		err := rows.Scan(
			&user.ID, &user.UserID, &user.Username, &user.DisplayName, &user.AvatarURL,
			&user.Bio, &user.Location, &user.Website, &user.Phone, &user.Timezone,
			&user.Language, &user.Theme, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

// CreatePreferences creates user preferences
func (r *UserRepositoryImpl) CreatePreferences(ctx context.Context, pref *entity.Preference) error {
	query := `
		INSERT INTO user_preferences (
			id, user_id, email_notifications, push_notifications,
			newsletter_subscription, weekly_digest, mentions_notifications,
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.ExecContext(ctx, query,
		pref.ID, pref.UserID, pref.EmailNotifications, pref.PushNotifications,
		pref.NewsletterSubscription, pref.WeeklyDigest, pref.MentionsNotifications,
		pref.CreatedAt, pref.UpdatedAt,
	)
	return err
}

// FindPreferencesByUserID retrieves user preferences by user ID
func (r *UserRepositoryImpl) FindPreferencesByUserID(ctx context.Context, userID uuid.UUID) (*entity.Preference, error) {
	query := `
		SELECT id, user_id, email_notifications, push_notifications,
		       newsletter_subscription, weekly_digest, mentions_notifications,
		       created_at, updated_at
		FROM user_preferences
		WHERE user_id = $1
	`
	pref := &entity.Preference{}
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&pref.ID, &pref.UserID, &pref.EmailNotifications, &pref.PushNotifications,
		&pref.NewsletterSubscription, &pref.WeeklyDigest, &pref.MentionsNotifications,
		&pref.CreatedAt, &pref.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("user preferences not found")
	}
	return pref, err
}

// UpdatePreferences updates user preferences
func (r *UserRepositoryImpl) UpdatePreferences(ctx context.Context, pref *entity.Preference) error {
	query := `
		UPDATE user_preferences
		SET email_notifications = $1, push_notifications = $2,
		    newsletter_subscription = $3, weekly_digest = $4,
		    mentions_notifications = $5, updated_at = $6
		WHERE user_id = $7
	`
	result, err := r.db.ExecContext(ctx, query,
		pref.EmailNotifications, pref.PushNotifications,
		pref.NewsletterSubscription, pref.WeeklyDigest,
		pref.MentionsNotifications, pref.UpdatedAt, pref.UserID,
	)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user preferences not found")
	}
	return nil
}
