package services

import (
	"agoravote-app-backend/src/database"
	"agoravote-app-backend/src/models"
)

// GroupService
// Business logic layer for group-related database operations
// Frontend: Used by GroupList.vue and GroupDetails.vue components
type GroupService struct {
	// Add fields if necessary
}

// NewGroupService
// Creates a new instance of GroupService
// Frontend: Used internally during app initialization, no direct component usage
func NewGroupService() *GroupService {
	return &GroupService{}
}

// CreateGroup
// Database operation to create new groups with validation
// Frontend: Called by NewGroupDialog.vue when clicking "Create Group" button
func (s *GroupService) CreateGroup(group *models.Group) error {
	return database.DB.Create(group).Error
}

// GetGroupByID
// Database query to fetch single group with members
// Frontend: Called by GroupDetails.vue when loading group page
func (s *GroupService) GetGroupByID(id string) (*models.Group, error) {
	var group models.Group
	if err := database.DB.Preload("Members").First(&group, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

// FetchGroups
// Database query to retrieve all accessible groups
// Frontend: Called by GroupList.vue when loading dashboard page
func (s *GroupService) FetchGroups() ([]models.Group, error) {
	var groups []models.Group
	if err := database.DB.Preload("Members").Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

// CreateGroupMember
// Adds a new member to a group
// Frontend: Called by InviteDialog.vue when clicking "Send Invitation" button
func CreateGroupMember(groupMember *models.GroupMember) error {
	return database.DB.Create(groupMember).Error
}

// GetUserGroups
// Retrieves groups where user is a member
// Frontend: Called by UserProfile.vue in "My Groups" section
func GetUserGroups(userID string, groups *[]models.Group) error {
	return database.DB.Joins("JOIN group_members ON group_members.group_id = groups.id").
		Where("group_members.user_id = ?", userID).
		Find(groups).Error
}
