package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"agoravote-app-backend/src/services"
)

type GroupController struct {
	service services.GroupService
}

func NewGroupController(service services.GroupService) *GroupController {
	return &GroupController{service: service}
}

func (gc *GroupController) CreateGroup(c *gin.Context) {
	var group services.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := gc.service.CreateGroup(&group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, group)
}

func (gc *GroupController) GetGroups(c *gin.Context) {
	groups, err := gc.service.FetchGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groups)
}