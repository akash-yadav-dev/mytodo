// Package service contains domain services for the users module.
//
// This file implements profile management business logic.

package service

// ProfileService handles user profile operations.
//
// In production applications, profile services typically implement:
// - Complete profile information retrieval
// - Profile updates with validation
// - Profile completeness checking
// - Social link validation
// - Profile visibility settings
//
// Example interface:
//   type ProfileService interface {
//       GetProfile(userID string) (*Profile, error)
//       UpdateProfile(userID string, updates ProfileUpdates) (*Profile, error)
//       GetProfileCompleteness(userID string) (int, error)
//       UpdateSocialLinks(userID string, links map[string]string) error
//   }
//
// Example usage:
//   profile, err := profileService.UpdateProfile("user-123", ProfileUpdates{
//       FirstName: "John",
//       LastName:  "Doe",
//       Company:   "Acme Corp",
//   })
//   // Returns: &Profile{FirstName: "John", LastName: "Doe", ...}, nil
//
//   completeness, _ := profileService.GetProfileCompleteness("user-123")
//   // Returns: 75 (percentage of completed fields)
