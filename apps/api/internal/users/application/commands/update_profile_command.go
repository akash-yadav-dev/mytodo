// Package commands implements Command pattern for write operations.

package commands

// UpdateProfileCommand represents a request to update user profile.
//
// Example structure:
//   type UpdateProfileCommand struct {
//       UserID    string            `json:"user_id" validate:"required"`
//       FirstName string            `json:"first_name,omitempty"`
//       LastName  string            `json:"last_name,omitempty"`
//       Company   string            `json:"company,omitempty"`
//       JobTitle  string            `json:"job_title,omitempty"`
//       Location  string            `json:"location,omitempty"`
//       Timezone  string            `json:"timezone,omitempty"`
//       Website   string            `json:"website,omitempty" validate:"omitempty,url"`
//       SocialLinks map[string]string `json:"social_links,omitempty"`
//   }
//
// Example usage:
//   cmd := UpdateProfileCommand{
//       UserID:    "user-123",
//       FirstName: "John",
//       LastName:  "Doe",
//       Company:   "Acme Corp",
//       JobTitle:  "Senior Developer",
//   }
//   result, err := handler.Handle(cmd)
//   // Returns: &Profile{FirstName: "John", LastName: "Doe",...}, nil
