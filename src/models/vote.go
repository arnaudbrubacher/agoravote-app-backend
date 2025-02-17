package models

type Vote struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	GroupID   string `gorm:"type:uuid;not null"`
	UserID    string `gorm:"not null"`
	CreatedAt string `gorm:"not null"`
}
