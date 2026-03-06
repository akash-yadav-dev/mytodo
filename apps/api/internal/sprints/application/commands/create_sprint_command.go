// Package commands implements Command pattern for sprints write operations.

package commands

// CreateSprintCommand creates a new sprint.
//
// Example:
//   cmd := CreateSprintCommand{
//       ProjectID: "proj-1",
//       Name:      "Sprint 1",
//       StartDate: time.Now(),
//       EndDate:   time.Now().AddDate(0, 0, 14),
//   }
