package models

type Post struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	GroupID   string `gorm:"type:uuid;not null"`
	Content   string `gorm:"not null"`
	CreatedAt string `gorm:"not null"`
}
