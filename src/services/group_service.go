package services

import (
	"agoravote-app-backend/src/database"
	"agoravote-app-backend/src/models"
)

// GroupService
// Business logic layer for group-related database operations
type GroupService struct {
	// Add fields if necessary
}

// NewGroupService
// Creates a new instance of GroupService
func NewGroupService() *GroupService {
	return &GroupService{}
}

// CreateGroup
// Database operation to create new groups with validation
func (s *GroupService) CreateGroup(group *models.Group) error {
	return database.DB.Create(group).Error
}

// GetGroupByID
// Database query to fetch single group with members
func (s *GroupService) GetGroupByID(id string) (*models.Group, error) {
	var group models.Group
	if err := database.DB.Preload("Members").First(&group, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

// FetchGroups
// Database query to retrieve all accessible groups
func (s *GroupService) FetchGroups() ([]models.Group, error) {
	var groups []models.Group
	if err := database.DB.Preload("Members").Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

// CreateGroupMember
// Adds a new member to a group
func CreateGroupMember(groupMember *models.GroupMember) error {
	return database.DB.Create(groupMember).Error
}

// GetUserGroups
// Retrieves groups where user is a member
func GetUserGroups(userID string, groups *[]models.Group) error {
	return database.DB.Joins("JOIN group_members ON group_members.group_id = groups.id").
		Where("group_members.user_id = ?", userID).
		Find(groups).Error
}
