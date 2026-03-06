// Package oauth provides OAuth 2.0 and social login integrations.
//
// This file implements GitHub OAuth integration.

package oauth

// GitHubOAuth implements OAuth 2.0 flow for GitHub authentication.
// In production applications, GitHub OAuth implementations typically:
// - Use GitHub's OAuth App or GitHub App authentication
// - Request scopes: user:email, read:user
// - Exchange code for access token
// - Fetch user profile from GitHub API
// - Handle primary/verified email selection
// - Support GitHub Enterprise Server
// - Respect API rate limits
// - Handle organization membership checks if needed
