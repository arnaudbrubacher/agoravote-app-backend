package models

type Vote struct {
	ID      uint `gorm:"primary_key"`
	GroupID string
	Group   Group `gorm:"foreignKey:GroupID"`
	Value   string
}

type WriteInResponse struct {
	ID     string `gorm:"primary_key"`
	VoteID string
	Text   string
	Count  int
}
