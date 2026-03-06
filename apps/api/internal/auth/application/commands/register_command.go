// Package commands implements Command pattern for write operations.
//
// This file defines the user registration command.

package commands

// RegisterCommand represents a new user registration request.
// In production applications, registration commands typically include:
// - Email address (unique identifier)
// - Password (will be hashed before storage)
// - Username or display name
// - Optional: first name and last name
// - Optional: phone number
// - Optional: profile picture URL
// - Optional: invitation code or referral token
// - Optional: terms of service acceptance timestamp
// - Optional: marketing consent flags
// - Device and location info for fraud detection
