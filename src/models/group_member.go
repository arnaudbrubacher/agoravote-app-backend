package models

import "time"

type GroupMember struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	GroupID   string    `gorm:"type:uuid;not null"`
	UserID    string    `gorm:"type:uuid;not null"`
	CreatedAt time.Time `gorm:"not null"`
	Group     *Group    `gorm:"foreignKey:GroupID"`
}
