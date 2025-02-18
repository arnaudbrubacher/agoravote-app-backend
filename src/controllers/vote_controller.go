package controllers

import (
	"agoravote-app-backend/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VoteController struct {
	VoteService services.VoteService
}

func NewVoteController(service services.VoteService) *VoteController {
	return &VoteController{VoteService: service}
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

// CreateVote handles the creation of a new vote
func CreateVote(c *gin.Context) {
	// Implement the logic to create a vote
	c.JSON(http.StatusOK, gin.H{"message": "Vote created successfully"})
}
