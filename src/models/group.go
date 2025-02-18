package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID          uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Picture     string        `json:"picture"`
	IsPrivate   bool          `json:"is_private"`
	LastActive  string        `json:"last_active"`
	CreatedAt   time.Time     `gorm:"not null"`
	Members     []GroupMember `gorm:"foreignKey:GroupID"`
}

func (group *Group) BeforeCreate(tx *gorm.DB) (err error) {
	group.ID = uuid.New()
	return
}
