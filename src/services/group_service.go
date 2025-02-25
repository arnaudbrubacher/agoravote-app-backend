package services

import (
	"agoravote-app-backend/src/database"
	"agoravote-app-backend/src/models"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/google/uuid"
)

// GroupService
// Business logic layer for group-related database operations
// Frontend: Used by GroupList.vue and GroupDetails.vue components
type GroupService struct{}

// NewGroupService
// Creates a new instance of GroupService
// Frontend: Used internally during app initialization, no direct component usage
func NewGroupService() *GroupService {
	return &GroupService{}
}

// CreateGroup creates a new group and adds the creator as admin
func (s *GroupService) CreateGroup(group *models.Group, creatorID uuid.UUID) error {
	// Start a transaction
	tx := database.DB.Begin()

	// Set creation time
	now := time.Now()
	group.CreatedAt = now
	group.UpdatedAt = now
	group.LastActive = now

	// Create the group
	if err := tx.Create(group).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Create group member entry for creator as admin
	member := &models.GroupMember{
		GroupID:   group.ID,
		UserID:    creatorID,
		IsAdmin:   true,
		CreatedAt: now,
	}

	if err := tx.Create(member).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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

func (s *GroupService) InviteToGroup(groupID uuid.UUID, email string) (*models.GroupInvitation, error) {
	// Generate invitation token
	b := make([]byte, 32)
	rand.Read(b)
	token := base64.URLEncoding.EncodeToString(b)

	invitation := &models.GroupInvitation{
		GroupID:   groupID,
		Email:     email,
		Token:     token,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		CreatedAt: time.Now(),
	}

	if err := database.DB.Create(invitation).Error; err != nil {
		return nil, err
	}

	return invitation, nil
}

func (s *GroupService) AcceptInvitation(token string, userID uuid.UUID) error {
	tx := database.DB.Begin()

	var invitation models.GroupInvitation
	if err := tx.Where("token = ? AND used = false AND expires_at > ?", token, time.Now()).First(&invitation).Error; err != nil {
		tx.Rollback()
		return err
	}

	member := &models.GroupMember{
		GroupID:   invitation.GroupID,
		UserID:    userID,
		IsAdmin:   false,
		CreatedAt: time.Now(),
	}

	if err := tx.Create(member).Error; err != nil {
		tx.Rollback()
		return err
	}

	invitation.Used = true
	if err := tx.Save(&invitation).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
