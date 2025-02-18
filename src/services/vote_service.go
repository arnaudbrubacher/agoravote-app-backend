package services

import (
	"agoravote-app-backend/src/database"
	"agoravote-app-backend/src/models"
)

type VoteService struct{}

func NewVoteService() VoteService {
	return VoteService{}
}

func (vs *VoteService) CreateVote(vote models.Vote) error {
	if err := database.DB.Create(&vote).Error; err != nil {
		return err
	}
	return nil
}

func (vs *VoteService) FetchVotes(groupID string) ([]models.Vote, error) {
	var votes []models.Vote
	if err := database.DB.Where("group_id = ?", groupID).Find(&votes).Error; err != nil {
		return nil, err
	}
	return votes, nil
}
