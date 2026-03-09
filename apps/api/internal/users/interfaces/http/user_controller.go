// Package http provides HTTP/REST API endpoints for users module.
//
// Controllers handle HTTP requests and delegate to application handlers.

package http

import (
	"mytodo/apps/api/internal/users/application/commands"
	"mytodo/apps/api/internal/users/application/handlers"
	"mytodo/apps/api/internal/users/application/queries"
	"mytodo/apps/api/internal/users/interfaces/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserController handles HTTP endpoints for user profile operations.
type UserController struct {
	queryHandler   *handlers.UserQueryHandler
	commandHandler *handlers.UserCommandHandler
}

// NewUserController creates a new UserController
func NewUserController(queryHandler *handlers.UserQueryHandler, commandHandler *handlers.UserCommandHandler) *UserController {
	return &UserController{
		queryHandler:   queryHandler,
		commandHandler: commandHandler,
	}
}

// GetCurrentUserProfile retrieves the authenticated user's profile
// GET /api/v1/users/me
func (c *UserController) GetCurrentUserProfile(ctx *gin.Context) {
	// Extract user ID from JWT token (set by auth middleware)
	userIDStr, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user ID",
		})
		return
	}

	query := queries.GetUserByIDQuery{
		UserID: userID.String(),
	}

	userProfile, err := c.queryHandler.HandleGetProfileByID(ctx.Request.Context(), query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "user profile not found",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userProfile,
	})
}

// GetUserProfileByID retrieves a user profile by ID
// GET /api/v1/users/:id
func (c *UserController) GetUserProfileByID(ctx *gin.Context) {
	userID := ctx.Param("id")

	query := queries.GetUserByIDQuery{
		UserID: userID,
	}

	userProfile, err := c.queryHandler.HandleGetProfileByID(ctx.Request.Context(), query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "user profile not found",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userProfile,
	})
}

// ListUserProfiles retrieves a paginated list of user profiles
// GET /api/v1/users
func (c *UserController) ListUserProfiles(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	query := queries.ListUsersQuery{
		Page:  page,
		Limit: pageSize,
	}

	users, err := c.queryHandler.HandleGetUserProfileList(ctx.Request.Context(), query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}

// SearchUserProfiles searches for user profiles
// GET /api/v1/users/search
func (c *UserController) SearchUserProfiles(ctx *gin.Context) {
	searchQuery := ctx.Query("q")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))

	query := queries.SearchUsersQuery{
		Query: searchQuery,
		Limit: limit,
	}

	users, err := c.queryHandler.HandleSearchUserProfiles(ctx.Request.Context(), query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}

// CreateUserProfile creates a new user profile (typically called after registration)
// POST /api/v1/users/profile
func (c *UserController) CreateUserProfile(ctx *gin.Context) {
	// Extract user ID from JWT token
	userIDStr, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	var req dto.CreateUserProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cmd := commands.CreateUserProfileCommand{
		AuthUserID:  userIDStr.(string),
		Username:    req.Username,
		DisplayName: req.DisplayName,
		AvatarURL:   req.AvatarURL,
	}

	userProfile, err := c.commandHandler.HandleCreateUserProfile(ctx.Request.Context(), cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    userProfile,
	})
}

// UpdateUserProfile updates user profile information
// PUT /api/v1/users/me
func (c *UserController) UpdateUserProfile(ctx *gin.Context) {
	// Extract user ID from JWT token
	userIDStr, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	var req dto.UpdateUserProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cmd := commands.UpdateUserProfileCommand{
		UserID:      userIDStr.(string),
		Username:    req.Username,
		DisplayName: req.DisplayName,
		Bio:         req.Bio,
		Location:    req.Location,
		Website:     req.Website,
		AvatarURL:   req.AvatarURL,
		Phone:       req.Phone,
		Timezone:    req.Timezone,
		Language:    req.Language,
		Theme:       req.Theme,
	}

	userProfile, err := c.commandHandler.HandleUpdateUserProfile(ctx.Request.Context(), cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userProfile,
	})
}

// DeleteUserProfile deletes the user's profile
// DELETE /api/v1/users/me
func (c *UserController) DeleteUserProfile(ctx *gin.Context) {
	// Extract user ID from JWT token
	userIDStr, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	cmd := commands.DeleteUserCommand{
		UserID: userIDStr.(string),
	}

	if err := c.commandHandler.HandleDeleteUserProfile(ctx.Request.Context(), cmd); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "user profile deleted successfully",
	})
}

// GetUserPreferences retrieves user preferences
// GET /api/v1/users/me/preferences
func (c *UserController) GetUserPreferences(ctx *gin.Context) {
	// Extract user ID from JWT token
	userIDStr, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user ID",
		})
		return
	}

	prefs, err := c.queryHandler.HandleGetUserPreferences(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    prefs,
	})
}

// UpdateUserPreferences updates user preferences
// PUT /api/v1/users/me/preferences
func (c *UserController) UpdateUserPreferences(ctx *gin.Context) {
	// Extract user ID from JWT token
	userIDStr, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	var req dto.UpdateUserPreferencesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cmd := commands.UpdateUserPreferencesCommand{
		UserID:                 userIDStr.(string),
		EmailNotifications:     req.EmailNotifications,
		PushNotifications:      req.PushNotifications,
		NewsletterSubscription: req.NewsletterSubscription,
		WeeklyDigest:           req.WeeklyDigest,
		MentionsNotifications:  req.MentionsNotifications,
	}

	prefs, err := c.commandHandler.HandleUpdateUserPreferences(ctx.Request.Context(), cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    prefs,
	})
}
