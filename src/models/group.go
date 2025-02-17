package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Picture     string
	IsPrivate   bool          `gorm:"not null"`
	LastActive  string        `gorm:"not null"`
	Members     []GroupMember `gorm:"foreignKey:GroupID"`
}

func (group *Group) BeforeCreate(tx *gorm.DB) (err error) {
	group.ID = uuid.New().String()
	return
}
