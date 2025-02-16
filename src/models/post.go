package models

type Post struct {
	ID      uint `gorm:"primary_key"`
	GroupID string
	Group   Group `gorm:"foreignKey:GroupID"`
	Content string
}
