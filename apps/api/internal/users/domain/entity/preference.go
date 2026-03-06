// Package entity defines core domain entities for the users module.
//
// This file contains the Preference entity for user settings and preferences.

package entity

// Preference represents user-specific settings and preferences.
// Stores personalized configuration options.
//
// In production applications, Preference entities typically include:
// - UserID: reference to the user
// - Theme: light, dark, auto
// - EmailNotifications: enable/disable email alerts
// - PushNotifications: enable/disable push notifications
// - NotificationSettings: granular notification preferences
// - DefaultView: dashboard, kanban, list
// - ItemsPerPage: pagination preference
// - DateFormat: user's preferred date format
// - TimeFormat: 12h or 24h
//
// Example structure:
//   type Preference struct {
//       UserID               string            `json:"user_id"`
//       Theme                string            `json:"theme"`
//       EmailNotifications   bool              `json:"email_notifications"`
//       NotificationSettings map[string]bool   `json:"notification_settings"`
//       DefaultView          string            `json:"default_view"`
//       ItemsPerPage         int               `json:"items_per_page"`
//       DateFormat           string            `json:"date_format"`
//   }
//
// Example methods:
//   func (p *Preference) EnableNotification(notificationType string)
//   func (p *Preference) SetTheme(theme string) error
//   func (p *Preference) GetDefaults() *Preference
