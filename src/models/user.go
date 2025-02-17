package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Role     string
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	return
}
