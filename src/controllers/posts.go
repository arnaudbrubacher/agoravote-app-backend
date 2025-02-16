package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"agoravote-app-backend/src/models"
	"agoravote-app-backend/src/services"
)

type PostController struct {
	PostService services.PostService
}

func NewPostController(service services.PostService) *PostController {
	return &PostController{PostService: service}
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.PostService.CreatePost(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (pc *PostController) FetchPosts(c *gin.Context) {
	posts, err := pc.PostService.FetchPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}