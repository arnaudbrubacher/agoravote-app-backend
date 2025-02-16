package models

type User struct {
    ID       string `gorm:"primary_key"`
    Email    string `gorm:"unique"`
    Password string
}