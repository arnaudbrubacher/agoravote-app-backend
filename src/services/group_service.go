package services

import (
	"agoravote-app-backend/src/database"
	"agoravote-app-backend/src/models"
)

type GroupService struct{}

func (s GroupService) GetGroupByID(id string) (*models.Group, error) {
	var group models.Group
	if err := database.DB.First(&group, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (s GroupService) FetchGroups() ([]models.Group, error) {
	var groups []models.Group
	if err := database.DB.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}
