// Package oauth provides OAuth 2.0 and social login integrations.
//
// OAuth infrastructure in Clean Architecture:
// - Integrates with third-party identity providers
// - Handles OAuth 2.0 authorization flow
// - Manages provider-specific configurations
// - Abstracts provider differences behind common interface
//
// In production applications, this folder typically contains:
// - OAuth provider client implementations (Google, GitHub, Facebook, etc.)
// - OAuth configuration (client IDs, secrets, redirect URIs)
// - Authorization code exchange logic
// - User profile fetching from providers
// - Token management for provider APIs
// - Provider-specific scopes and permissions
// - Account linking logic (connect OAuth to local user)
//
// Best practices:
// - Store credentials securely (environment variables, secrets manager)
// - Validate state parameter to prevent CSRF
// - Handle provider rate limits
// - Implement proper error handling for failed OAuth flows
// - Support multiple redirect URIs (dev, staging, prod)
// - Cache provider discovery documents

package oauth

// OAuthConfig manages configuration for OAuth providers.
// In production applications, OAuth configs typically include:
// - Provider-specific client IDs and secrets
// - Redirect URIs for each environment
// - OAuth scopes required
// - Provider endpoints (authorization, token, userinfo)
// - Timeout configurations
// - Provider-specific settings (PKCE, nonce)
