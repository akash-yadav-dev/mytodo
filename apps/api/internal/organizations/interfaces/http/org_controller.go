package http

import (
	"mytodo/apps/api/internal/organizations/application/commands"
	"mytodo/apps/api/internal/organizations/application/handlers"
	"mytodo/apps/api/internal/organizations/application/queries"
	"mytodo/apps/api/internal/organizations/interfaces/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// OrganizationController handles HTTP endpoints for organization operations.
type OrganizationController struct {
	queryHandler   *handlers.OrganizationQueryHandler
	commandHandler *handlers.OrganizationCommandHandler
}

// NewOrganizationController creates a new OrganizationController
func NewOrganizationController(
	queryHandler *handlers.OrganizationQueryHandler,
	commandHandler *handlers.OrganizationCommandHandler,
) *OrganizationController {
	return &OrganizationController{
		queryHandler:   queryHandler,
		commandHandler: commandHandler,
	}
}

// CreateOrganization creates a new organization
// POST /api/v1/organizations
func (c *OrganizationController) CreateOrganization(ctx *gin.Context) {
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

	var req dto.CreateOrganizationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.Normalize()

	cmd := commands.CreateOrganizationCommand{
		OwnerID:     userID,
		CreatedBy:   userID,
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		PlanID:      req.PlanID,
	}

	organization, err := c.commandHandler.HandleCreateOrganization(ctx.Request.Context(), cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    organization,
	})
}

// GetOrganization retrieves an organization by ID
// GET /api/v1/organizations/:id
func (c *OrganizationController) GetOrganization(ctx *gin.Context) {
	orgIDStr := ctx.Param("id")
	orgID, err := uuid.Parse(orgIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid organization ID",
		})
		return
	}

	query := queries.GetOrganizationQuery{
		OrgID: orgID,
	}

	organization, err := c.queryHandler.HandleGetOrganization(ctx.Request.Context(), query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "organization not found",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    organization,
	})
}

// ListOrganizations retrieves a paginated list of organizations
// GET /api/v1/organizations
func (c *OrganizationController) ListOrganizations(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	query := queries.ListOrganizationsQuery{
		Page:  page,
		Limit: limit,
	}

	result, err := c.queryHandler.HandleListOrganizations(ctx.Request.Context(), query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result.Data,
		"total":   result.Total,
		"page":    result.Page,
		"limit":   result.Limit,
	})
}

// SearchOrganizations searches for organizations
// GET /api/v1/organizations/search
func (c *OrganizationController) SearchOrganizations(ctx *gin.Context) {
	searchQuery := ctx.Query("q")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))

	if searchQuery == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "search query is required",
		})
		return
	}

	query := queries.SearchOrganizationsQuery{
		Query: searchQuery,
		Limit: limit,
	}

	organizations, err := c.queryHandler.HandleSearchOrganizations(ctx.Request.Context(), query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    organizations,
	})
}

// UpdateOrganization updates an organization
// PUT /api/v1/organizations/:id
func (c *OrganizationController) UpdateOrganization(ctx *gin.Context) {
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

	orgIDStr := ctx.Param("id")
	orgID, err := uuid.Parse(orgIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid organization ID",
		})
		return
	}

	var req dto.UpdateOrganizationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: Add authorization check - only owner or admin can update

	cmd := commands.UpdateOrganizationCommand{
		OrgID:       orgID,
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		PlanID:      req.PlanID,
		UpdatedBy:   userID,
	}

	organization, err := c.commandHandler.HandleUpdateOrganization(ctx.Request.Context(), cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    organization,
	})
}

// DeleteOrganization deletes an organization (soft delete)
// DELETE /api/v1/organizations/:id
func (c *OrganizationController) DeleteOrganization(ctx *gin.Context) {
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

	orgIDStr := ctx.Param("id")
	orgID, err := uuid.Parse(orgIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid organization ID",
		})
		return
	}

	// TODO: Add authorization check - only owner can delete

	cmd := commands.DeleteOrganizationCommand{
		OrgID: orgID,
		DeletedBy: userID,
	}

	err = c.commandHandler.HandleDeleteOrganization(ctx.Request.Context(), cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "organization deleted successfully",
	})
}

// GetMyOrganizations retrieves organizations owned by the current user
// GET /api/v1/organizations/me/owned
func (c *OrganizationController) GetMyOrganizations(ctx *gin.Context) {
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

	query := queries.GetOrganizationsByOwnerQuery{
		OwnerID: userID,
	}

	organizations, err := c.queryHandler.HandleGetOrganizationsByOwner(ctx.Request.Context(), query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    organizations,
	})
}

// GetMemberOrganizations retrieves organizations where the current user is a member
// GET /api/v1/organizations/me/member
func (c *OrganizationController) GetMemberOrganizations(ctx *gin.Context) {
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

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	query := queries.GetMemberOrganizationsQuery{
		UserID: userID,
		Page:   page,
		Limit:  limit,
	}

	result, err := c.queryHandler.HandleGetMemberOrganizations(ctx.Request.Context(), query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result.Data,
		"total":   result.Total,
		"page":    result.Page,
		"limit":   result.Limit,
	})
}

// TransferOwnership transfers organization ownership to another user
// POST /api/v1/organizations/:id/transfer
func (c *OrganizationController) TransferOwnership(ctx *gin.Context) {
	// Extract user ID from JWT token
	userIDStr, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	currentOwnerID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user ID",
		})
		return
	}

	orgIDStr := ctx.Param("id")
	orgID, err := uuid.Parse(orgIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid organization ID",
		})
		return
	}

	var req dto.TransferOwnershipRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newOwnerID, err := uuid.Parse(req.NewOwnerID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid new owner ID",
		})
		return
	}

	cmd := commands.TransferOwnershipCommand{
		OrgID:          orgID,
		NewOwnerID:     newOwnerID,
		CurrentOwnerID: currentOwnerID,
	}

	err = c.commandHandler.HandleTransferOwnership(ctx.Request.Context(), cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ownership transferred successfully",
	})
}
