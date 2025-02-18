package services

import (
	"agoravote-app-backend/src/database"
	"agoravote-app-backend/src/models"
)

type PostService struct{}

func NewPostService() PostService {
	return PostService{}
}

func (ps *PostService) CreatePost(post models.Post) error {
	return database.DB.Create(&post).Error
}

func (ps *PostService) FetchPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := database.DB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
