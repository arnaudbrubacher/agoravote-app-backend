package services

import (
	"errors"
	"sync"

	"agoravote-app-backend/src/models"
)

type PostService struct {
	mu    sync.Mutex
	posts []models.Post
}

func NewPostService() *PostService {
	return &PostService{
		posts: []models.Post{},
	}
}

func (s *PostService) CreatePost(post models.Post) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if post.Title == "" || post.Content == "" {
		return errors.New("post title and content cannot be empty")
	}

	s.posts = append(s.posts, post)
	return nil
}

func (s *PostService) FetchPosts() []models.Post {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.posts
}