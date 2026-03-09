// Package http provides HTTP/REST API endpoints for the auth module.
//
// Interfaces/HTTP layer in Clean Architecture:
// - Handles HTTP requests and responses
// - Implements REST API endpoints
// - Manages request/response serialization (JSON, XML)
// - Applies HTTP-specific concerns (routing, middleware)
//
// In production applications, this folder typically contains:
// - HTTP controllers/handlers for each resource
// - Route definitions and registration
// - HTTP middleware (auth, logging, CORS, rate limiting)
// - Request validation and binding
// - Response formatting and error handling
// - API versioning logic
// - OpenAPI/Swagger documentation
//
// Best practices:
// - Keep controllers thin - delegate to application handlers
// - Use DTOs for request/response (don't expose domain entities)
// - Handle HTTP status codes appropriately
// - Implement proper error responses (RFC 7807)
// - Apply security headers
// - Implement request ID tracking
// - Use content negotiation

package http

import (
	"mytodo/apps/api/internal/auth/domain/service"
	"mytodo/apps/api/internal/auth/interfaces/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AuthController handles HTTP endpoints for authentication operations.
type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Register handles user registration
// POST /api/v1/auth/register
func (h *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.RegisterUser(c.Request.Context(), req.Email, req.Password, req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Automatically login after registration
	_, tokens, err := h.authService.AuthenticateUser(
		c.Request.Context(),
		req.Email,
		req.Password,
		c.Request.UserAgent(),
		c.ClientIP(),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "registration succeeded but login failed"})
		return
	}

	response := dto.AuthResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    tokens.ExpiresIn,
		User:         dto.ToUserDTO(user),
	}

	c.JSON(http.StatusCreated, response)
}

// Login handles user authentication
// POST /api/v1/auth/login
func (h *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, tokens, err := h.authService.AuthenticateUser(
		c.Request.Context(),
		req.Email,
		req.Password,
		c.Request.UserAgent(),
		c.ClientIP(),
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	response := dto.AuthResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    tokens.ExpiresIn,
		User:         dto.ToUserDTO(user),
	}

	c.JSON(http.StatusOK, response)
}

// RefreshToken handles token refresh
// POST /api/v1/auth/refresh
func (h *AuthController) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.authService.RefreshAccessToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	response := dto.AuthResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    tokens.ExpiresIn,
	}

	c.JSON(http.StatusOK, response)
}

// Logout handles user logout
// POST /api/v1/auth/logout
func (h *AuthController) Logout(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.authService.Logout(c.Request.Context(), req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{Message: "logged out successfully"})
}

// Me returns the current authenticated user
// GET /api/v1/auth/me
func (h *AuthController) Me(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userUUID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.authService.GetUserByID(c.Request.Context(), userUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserDTO(user))
}
