package controllers

import (
	"agoravote-app-backend/src/database"
	"agoravote-app-backend/src/models"
	"agoravote-app-backend/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GroupController handles group-related requests
type GroupController struct {
	groupService *services.GroupService
}

// NewGroupController creates a new GroupController
func NewGroupController(groupService *services.GroupService) *GroupController {
	return &GroupController{groupService: groupService}
}

// CreateGroup handles the creation of a new group
func (gc *GroupController) CreateGroup(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from JWT token
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if err := gc.groupService.CreateGroup(&group, uuid.MustParse(userID.(string))); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, group)
}

// GetGroup
// Retrieves a single group by its ID with members
// Frontend: Called by GroupDetails.vue when loading /groups/:id page
func (gc *GroupController) GetGroup(c *gin.Context) {
	groupID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID format"})
		return
	}

	group, err := gc.groupService.GetGroupByID(groupID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}

	c.JSON(http.StatusOK, group)
}

// GetGroups
// Retrieves all groups accessible to the user
// Frontend: Called by GroupList.vue when loading /dashboard page
func (gc *GroupController) GetGroups(c *gin.Context) {
	groups, err := gc.groupService.FetchGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groups)
}

// GetUserGroups
// Retrieves all groups where the user is a member
// Frontend: Called by UserProfile.vue in "My Groups" section on /profile page
func (gc *GroupController) GetUserGroups(c *gin.Context) {
	userID := c.GetString("user_id") // Get the user ID from the context
	var groups []models.Group
	if err := services.GetUserGroups(userID, &groups); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groups)
}

// InviteToGroup invites a user to a group
func (gc *GroupController) InviteToGroup(c *gin.Context) {
	groupID := c.Param("id")
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invitation, err := gc.groupService.InviteToGroup(uuid.MustParse(groupID), req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO: Send email with invitation link
	c.JSON(http.StatusOK, invitation)
}

// AcceptInvitation accepts a group invitation
func (gc *GroupController) AcceptInvitation(c *gin.Context) {
	token := c.Param("token")
	userID := c.GetString("user_id")

	if err := gc.groupService.AcceptInvitation(token, uuid.MustParse(userID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully joined group"})
}

// [DEPRECATED] GetGroups
// Old implementation to be removed
// Frontend: No longer used, replaced by GroupController.GetGroups
func GetGroups(c *gin.Context) {
	var groups []models.Group
	if err := database.DB.Find(&groups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch groups"})
		return
	}

	c.JSON(http.StatusOK, groups)
}
