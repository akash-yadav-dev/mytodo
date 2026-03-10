package http

import (
	"mytodo/apps/api/internal/organizations/application/commands"
	"mytodo/apps/api/internal/organizations/application/handlers"
	"mytodo/apps/api/internal/organizations/application/queries"
	"mytodo/apps/api/internal/organizations/domain/entity"
	"mytodo/apps/api/internal/organizations/interfaces/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MemberController handles HTTP endpoints for organization member operations.
// This is included in the team_controller.go but can be separated if needed
type MemberController struct {
	queryHandler   *handlers.OrganizationQueryHandler
	commandHandler *handlers.OrganizationCommandHandler
	memberHandler  *handlers.MemberCommandHandler
}

// NewMemberController creates a new MemberController
func NewMemberController(
	queryHandler *handlers.OrganizationQueryHandler,
	commandHandler *handlers.OrganizationCommandHandler,
	memberHandler *handlers.MemberCommandHandler,
) *MemberController {
	return &MemberController{
		queryHandler:   queryHandler,
		commandHandler: commandHandler,
		memberHandler:  memberHandler,
	}
}

// AddMember adds a member to an organization
// POST /api/v1/organizations/:id/members
func (c *MemberController) AddMember(ctx *gin.Context) {
	// Extract user ID from JWT token
	userIDStr, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	inviterID, err := uuid.Parse(userIDStr.(string))
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

	var req dto.AddMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUserID, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user ID",
		})
		return
	}

	cmd := commands.AddMemberCommand{
		OrgID:     orgID,
		UserID:    newUserID,
		Role:      entity.Role(req.Role),
		InvitedBy: inviterID,
	}

	member, err := c.memberHandler.HandleAddMember(ctx.Request.Context(), cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    member,
	})
}

// RemoveMember removes a member from an organization
// DELETE /api/v1/organizations/:id/members/:userId
func (c *MemberController) RemoveMember(ctx *gin.Context) {
	// Extract user ID from JWT token
	userIDStr, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	removerID, err := uuid.Parse(userIDStr.(string))
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

	memberUserIDStr := ctx.Param("userId")
	memberUserID, err := uuid.Parse(memberUserIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid member user ID",
		})
		return
	}

	cmd := commands.RemoveMemberCommand{
		OrgID:     orgID,
		UserID:    memberUserID,
		RemovedBy: removerID,
	}

	err = c.memberHandler.HandleRemoveMember(ctx.Request.Context(), cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "member removed successfully",
	})
}

// UpdateMemberRole updates a member's role in an organization
// PATCH /api/v1/organizations/:id/members/:userId/role
func (c *MemberController) UpdateMemberRole(ctx *gin.Context) {
	// Extract user ID from JWT token
	userIDStr, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	updaterID, err := uuid.Parse(userIDStr.(string))
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

	memberUserIDStr := ctx.Param("userId")
	memberUserID, err := uuid.Parse(memberUserIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid member user ID",
		})
		return
	}

	var req dto.UpdateMemberRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cmd := commands.UpdateMemberRoleCommand{
		OrgID:     orgID,
		UserID:    memberUserID,
		NewRole:   req.Role,
		UpdatedBy: updaterID,
	}

	member, err := c.memberHandler.HandleUpdateMemberRole(ctx.Request.Context(), cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    member,
	})
}

// ListMembers lists all members of an organization
// GET /api/v1/organizations/:id/members
func (c *MemberController) ListMembers(ctx *gin.Context) {
	orgIDStr := ctx.Param("id")
	orgID, err := uuid.Parse(orgIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid organization ID",
		})
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))

	query := queries.ListMembersQuery{
		OrgID: orgID,
		Page:  page,
		Limit: limit,
	}

	result, err := c.memberHandler.HandleListMembers(ctx.Request.Context(), query)
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
