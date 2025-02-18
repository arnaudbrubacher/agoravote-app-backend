package controllers

import (
	"agoravote-app-backend/src/models"
	"agoravote-app-backend/src/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GroupController struct {
	GroupService services.GroupService
}

func (gc *GroupController) CreateGroup(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := gc.GroupService.CreateGroup(&group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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

func GetGroups(c *gin.Context) {
	var groups []models.Group
	if err := services.GetAllGroups(&groups); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groups)
}

func CreateGroup(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetString("user_id") // Get the user ID from the context

	if err := services.CreateGroup(&group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Add the user who created the group as a member
	groupMember := models.GroupMember{
		GroupID:   group.ID,
		UserID:    userID,
		CreatedAt: time.Now(), // Use time.Now() to get the current time as a time.Time value
	}
	if err := services.CreateGroupMember(&groupMember); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, group)
}

func GetUserGroups(c *gin.Context) {
	userID := c.GetString("user_id") // Get the user ID from the context
	var groups []models.Group
	if err := services.GetUserGroups(userID, &groups); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groups)
}
