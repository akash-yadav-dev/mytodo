// Package entity defines core domain entities for the users module.
//
// This file contains the Preference entity for user settings and preferences.

package entity

import (
	"time"

	"github.com/google/uuid"
)

// Preference represents user-specific settings and preferences.
type Preference struct {
	ID                     uuid.UUID `json:"id"`
	UserID                 uuid.UUID `json:"user_id"`
	EmailNotifications     bool      `json:"email_notifications"`
	PushNotifications      bool      `json:"push_notifications"`
	NewsletterSubscription bool      `json:"newsletter_subscription"`
	WeeklyDigest           bool      `json:"weekly_digest"`
	MentionsNotifications  bool      `json:"mentions_notifications"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

// NewUserPreferences creates default user preferences
func NewUserPreferences(userID uuid.UUID) *Preference {
	now := time.Now()
	return &Preference{
		ID:                     uuid.New(),
		UserID:                 userID,
		EmailNotifications:     true,
		PushNotifications:      true,
		NewsletterSubscription: false,
		WeeklyDigest:           true,
		MentionsNotifications:  true,
		CreatedAt:              now,
		UpdatedAt:              now,
	}
}

// UpdateNotificationSettings updates notification preferences
func (p *Preference) UpdateNotificationSettings(email, push, newsletter, digest, mentions bool) {
	p.EmailNotifications = email
	p.PushNotifications = push
	p.NewsletterSubscription = newsletter
	p.WeeklyDigest = digest
	p.MentionsNotifications = mentions
	p.UpdatedAt = time.Now()
}
