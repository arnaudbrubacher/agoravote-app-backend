package models

type GroupMember struct {
    ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    GroupID   string `gorm:"type:uuid;not null"`
    UserID    string `gorm:"type:uuid;not null"`
    CreatedAt string `gorm:"not null"`
}