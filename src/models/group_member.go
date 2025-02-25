package models

import (
	"time"

	"github.com/google/uuid"
)

type GroupMember struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	GroupID   uuid.UUID `json:"group_id" gorm:"type:uuid;not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	IsAdmin   bool      `json:"is_admin" gorm:"type:boolean;default:false;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp with time zone;not null"`
	Group     *Group    `json:"group,omitempty" gorm:"foreignKey:GroupID"`
}
