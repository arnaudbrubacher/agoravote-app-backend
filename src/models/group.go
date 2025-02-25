package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID          uuid.UUID     `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string        `json:"name" gorm:"type:text;not null"`
	Description string        `json:"description" gorm:"type:text"`
	Picture     string        `json:"picture" gorm:"type:text"`
	IsPrivate   bool          `json:"is_private" gorm:"type:boolean;not null"`
	LastActive  time.Time     `json:"last_active" gorm:"type:timestamp with time zone;not null"`
	CreatedAt   time.Time     `json:"created_at" gorm:"type:timestamp with time zone;default:now()"`
	UpdatedAt   time.Time     `json:"updated_at" gorm:"type:timestamp with time zone;default:now()"`
	Members     []GroupMember `json:"members" gorm:"foreignKey:GroupID"`
}

func (group *Group) BeforeCreate(tx *gorm.DB) (err error) {
	group.ID = uuid.New()
	return
}
