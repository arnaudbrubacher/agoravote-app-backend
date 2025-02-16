package models

type Vote struct {
    ID       string   `gorm:"primary_key"`
    Question string
    Choices  []string `gorm:"type:text[]"`
}