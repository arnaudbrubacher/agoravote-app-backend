package controllers

import (
	"agoravote-app-backend/src/models"
	"agoravote-app-backend/src/services"
	"net/http"

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
