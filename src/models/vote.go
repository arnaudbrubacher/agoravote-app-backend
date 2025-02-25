package models

import (
	"time"

	"github.com/google/uuid"
)

type Vote struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupID   uuid.UUID `json:"group_id" gorm:"type:uuid;not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Value     string    `json:"value" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp with time zone;not null"`
	Group     *Group    `json:"group,omitempty" gorm:"foreignKey:GroupID"`
	User      *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
