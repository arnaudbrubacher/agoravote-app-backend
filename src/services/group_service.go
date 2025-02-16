package services

import (
	"errors"
	"agoravote-app-backend/src/models"
	"agoravote-app-backend/src/database"
)

type GroupService struct{}

func NewGroupService() GroupService {
	return GroupService{}
}

func (gs *GroupService) CreateGroup(group models.Group) error {
	if err := database.DB.Create(&group).Error; err != nil {
		return err
	}
	return nil
}

func (gs *GroupService) GetGroupByID(groupID string) (models.Group, error) {
	var group models.Group
	if err := database.DB.First(&group, "id = ?", groupID).Error; err != nil {
		return group, errors.New("group not found")
	}
	return group, nil
}