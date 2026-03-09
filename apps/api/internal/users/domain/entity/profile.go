// Package entity defines core domain entities for the users module.
//
// This file contains the Profile entity for extended user information.

package entity

import "time"

// Profile represents extended user profile information.
// Separates core user data from optional profile details.
//
// In production applications, Profile entities typically include:
// - UserID: reference to the user
// - FirstName, LastName: personal information
// - Company: organization name
// - JobTitle: professional title
// - Location: city, country
// - Website: personal or company website
// - SocialLinks: GitHub, Twitter, LinkedIn URLs
// - Timezone: user's timezone for scheduling
// - Language: preferred language (i18n)
//
// Example structure:
type Profile struct {
	UserID      string            `json:"user_id"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	Company     string            `json:"company"`
	JobTitle    string            `json:"job_title"`
	Location    string            `json:"location"`
	Timezone    string            `json:"timezone"`
	SocialLinks map[string]string `json:"social_links"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

//
// Example methods:
func (p *Profile) GetFullName() string { return p.FirstName + " " + p.LastName }
func (p *Profile) UpdateContactInfo(company, jobTitle string) error {
	p.Company = company
	p.JobTitle = jobTitle
	p.UpdatedAt = time.Now()
	return nil
}
