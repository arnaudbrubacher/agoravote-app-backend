package models

type Group struct {
	ID          string `gorm:"primary_key"`
	Name        string
	Description string
	Picture     *string
	IsPrivate   bool
	Members     []User `gorm:"many2many:group_members;"`
	Votes       []Vote `gorm:"foreignKey:GroupID"`
	Posts       []Post `gorm:"foreignKey:GroupID"`
	LastActive  string
}
