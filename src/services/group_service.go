package services

import (
	"agoravote-app-backend/src/database"
	"agoravote-app-backend/src/models"
)

type GroupService struct {
	// Add fields if necessary
}

func NewGroupService() *GroupService {
	return &GroupService{}
}

func (s *GroupService) CreateGroup(group *models.Group) error {
	return database.DB.Create(group).Error
}

func (s *GroupService) GetGroupByID(id string) (*models.Group, error) {
	var group models.Group
	if err := database.DB.Preload("Members").First(&group, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (s *GroupService) FetchGroups() ([]models.Group, error) {
	var groups []models.Group
	if err := database.DB.Preload("Members").Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func GetAllGroups(groups *[]models.Group) error {
	return database.DB.Preload("Members").Find(groups).Error
}

func CreateGroupMember(groupMember *models.GroupMember) error {
	return database.DB.Create(groupMember).Error
}

func GetUserGroups(userID string, groups *[]models.Group) error {
	return database.DB.Joins("JOIN group_members ON group_members.group_id = groups.id").
		Where("group_members.user_id = ?", userID).
		Find(groups).Error
}
