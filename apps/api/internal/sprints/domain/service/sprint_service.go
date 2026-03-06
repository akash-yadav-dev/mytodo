// Package service contains domain services for the sprints module.

package service

// SprintService handles sprint business logic.
//
// Example interface:
//   type SprintService interface {
//       CreateSprint(projectID, name string, startDate, endDate time.Time) (*Sprint, error)
//       StartSprint(sprintID string) error
//       CompleteSprint(sprintID string) error
//       GetActiveSprint(projectID string) (*Sprint, error)
//   }
//
// Example usage:
//   sprint, _ := svc.StartSprint("sprint-1")
//   // Changes sprint status to "active" and starts tracking
