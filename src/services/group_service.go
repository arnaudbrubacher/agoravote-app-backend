package services

import (
	"agoravote-app-backend/src/database"
	"agoravote-app-backend/src/models"
)

func GetAllGroups(groups *[]models.Group) error {
	return database.DB.Preload("Members").Find(groups).Error
}

func CreateGroup(group *models.Group) error {
	return database.DB.Create(group).Error
}

func CreateGroupMember(groupMember *models.GroupMember) error {
	return database.DB.Create(groupMember).Error
}

func GetUserGroups(userID string, groups *[]models.Group) error {
	return database.DB.Joins("JOIN group_members ON group_members.group_id = groups.id").
		Where("group_members.user_id = ?", userID).
		Find(groups).Error
}
