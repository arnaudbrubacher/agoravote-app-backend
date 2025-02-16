package models

type User struct {
    ID       string `gorm:"primary_key"`
    Name     string
    Email    string `gorm:"unique"`
    Password string
    Role     string
}