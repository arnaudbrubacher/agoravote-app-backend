package controllers

import (
	"agoravote-app-backend/src/models"
	"agoravote-app-backend/src/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type VoteController struct {
	VoteService services.VoteService
}

func NewVoteController(service services.VoteService) *VoteController {
	return &VoteController{VoteService: service}
}

func (vc *VoteController) CreateVote(c *gin.Context) {
	var vote models.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vote.CreatedAt = time.Now()

	if err := vc.VoteService.CreateVote(vote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, vote)
}

func (vc *VoteController) FetchVotes(c *gin.Context) {
	groupID := c.Param("group_id")
	votes, err := vc.VoteService.FetchVotes(groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, votes)
}
