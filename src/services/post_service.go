package services

import (
    "agoravote-app-backend/src/models"
    "agoravote-app-backend/src/database"
)

type PostService struct{}

func NewPostService() PostService {
    return PostService{}
}

func (ps *PostService) CreatePost(post models.Post) error {
    if err := database.DB.Create(&post).Error; err != nil {
        return err
    }
    return nil
}

func (ps *PostService) FetchPosts() ([]models.Post, error) {
    var posts []models.Post
    if err := database.DB.Find(&posts).Error; err != nil {
        return nil, err
    }
    return posts, nil
}