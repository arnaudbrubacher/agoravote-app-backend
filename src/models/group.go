package models

type Group struct {
    ID   string `gorm:"primary_key"`
    Name string
}