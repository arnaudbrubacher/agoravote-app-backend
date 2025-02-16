package models

type Post struct {
    ID       string `gorm:"primary_key"`
    Title    string
    Content  string
    Author   string
    Date     string
    FileName *string
    FileUrl  *string
}