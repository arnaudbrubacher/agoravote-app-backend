package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"agoravote-app-backend/src/services"
)

type VoteController struct {
	VoteService services.VoteService
}

func (vc *VoteController) CastVote(c *gin.Context) {
	var vote services.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := vc.VoteService.CastVote(vote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Vote cast successfully"})
}

func (vc *VoteController) GetVotes(c *gin.Context) {
	votes, err := vc.VoteService.FetchVotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, votes)
}