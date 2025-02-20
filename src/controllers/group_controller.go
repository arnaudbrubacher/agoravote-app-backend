package controllers

import (
	"agoravote-app-backend/src/database"
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
		GroupID:   group.ID.String(),
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

// GetGroups handles the request to get all groups
func GetGroups(c *gin.Context) {
	var groups []models.Group
	if err := database.DB.Find(&groups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch groups"})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// CreateGroup handles the request to create a new group
func CreateGroup(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := database.DB.Create(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create group"})
		return
	}

	c.JSON(http.StatusOK, group)
}

// GetUserGroups handles the request to get user groups
func GetUserGroups(c *gin.Context) {
	// Implement the logic to get user groups
	c.JSON(http.StatusOK, gin.H{"message": "GetUserGroups endpoint"})
}
