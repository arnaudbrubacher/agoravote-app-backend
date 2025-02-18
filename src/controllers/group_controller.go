package controllers

import (
	"agoravote-app-backend/src/models"
	"agoravote-app-backend/src/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GroupController struct {
	GroupService *services.GroupService
}

func (gc *GroupController) CreateGroup(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group.CreatedAt = time.Now() // Set the CreatedAt field

	if err := gc.GroupService.CreateGroup(&group); err != nil {
		log.Println("Error creating group:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Add the user who created the group as a member
	userID := c.GetString("user_id") // Get the user ID from the context
	groupMember := models.GroupMember{
		GroupID:   group.ID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	if err := services.CreateGroupMember(&groupMember); err != nil {
		log.Println("Error creating group member:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Group created successfully:", group)
	c.JSON(http.StatusOK, group)
}

func (gc *GroupController) GetGroup(c *gin.Context) {
	id := c.Param("id")
	group, err := gc.GroupService.GetGroupByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}
	c.JSON(http.StatusOK, group)
}

func (gc *GroupController) GetGroups(c *gin.Context) {
	groups, err := gc.GroupService.FetchGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (gc *GroupController) GetUserGroups(c *gin.Context) {
	userID := c.GetString("user_id") // Get the user ID from the context
	var groups []models.Group
	if err := services.GetUserGroups(userID, &groups); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groups)
}

// GetUserGroups handles the request to get user groups
func GetUserGroups(c *gin.Context) {
	// Implement the logic to get user groups
	c.JSON(http.StatusOK, gin.H{"message": "GetUserGroups endpoint"})
}

// GetGroups handles the request to get all groups
func GetGroups(c *gin.Context) {
	// Implement your logic to get groups here
	c.JSON(http.StatusOK, gin.H{"message": "GetGroups endpoint"})
}

// CreateGroup handles the creation of a new group
func CreateGroup(c *gin.Context) {
	// Implement the logic to create a group
	c.JSON(http.StatusOK, gin.H{"message": "Group created successfully"})
}
