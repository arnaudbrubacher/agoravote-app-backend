package models

import (
	"time"

	"github.com/google/uuid"
)

type GroupInvitation struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	GroupID   uuid.UUID `json:"group_id" gorm:"type:uuid;not null"`
	Email     string    `json:"email" gorm:"type:text;not null"`
	Token     string    `json:"token" gorm:"type:text;not null;unique"`
	ExpiresAt time.Time `json:"expires_at" gorm:"type:timestamp with time zone;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp with time zone;default:now()"`
	Used      bool      `json:"used" gorm:"type:boolean;default:false"`
}
