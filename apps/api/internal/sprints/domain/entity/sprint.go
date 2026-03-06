// Package entity defines core domain entities for the sprints module.
//
// Sprints represent time-boxed iterations in agile project management.

package entity

// Sprint represents a time-boxed iteration.
//
// In production, Sprint entities include:
// - ID, Name, ProjectID
// - StartDate, EndDate
// - Goals and objectives
// - Status (planning, active, completed)
// - Velocity and capacity
//
// Example structure:
//   type Sprint struct {
//       ID        string    `json:"id"`
//       Name      string    `json:"name"`
//       ProjectID string    `json:"project_id"`
//       StartDate time.Time `json:"start_date"`
//       EndDate   time.Time `json:"end_date"`
//       Status    string    `json:"status"`
//   }
//
// Example methods:
//   func (s *Sprint) Start() error
//   func (s *Sprint) Complete() error
//   func (s *Sprint) IsActive() bool
