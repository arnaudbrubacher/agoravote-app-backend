package controllers

import (
	"agoravote-app-backend/src/models"
	"agoravote-app-backend/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserController
// Handles HTTP requests for user-related operations
// Frontend: Used across all authenticated pages for user management
type UserController struct {
	UserService *services.UserService
}

// NewUserController
// Creates a new user controller with injected user service
// Frontend: Not directly used - internal initialization
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

// CreateUser
// Handles HTTP request to create a new user account
// Frontend: Called from SignupForm.vue component's "Sign Up" button
func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUser
// Retrieves user details excluding sensitive information
// Frontend: Called from UserProfile.vue component when loading profile page
func (uc *UserController) GetUser(c *gin.Context) {
	userId := c.Param("id")
	user, err := uc.UserService.GetUserByID(uuid.MustParse(userId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// do not send the user password , store the hashed password not original ones
	user.Password = ""

	c.JSON(http.StatusOK, user)
}

// GetUserProfile
// Retrieves public user profile by ID
// Frontend: Called from UserCard.vue component when viewing other users' profiles
func (uc *UserController) GetUserProfile(c *gin.Context) {
	id := c.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := uc.UserService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	user.Password = ""

	c.JSON(http.StatusOK, user)
}

// DeleteUserAccount
// Removes user account after verifying ownership
// Frontend: Called from AccountSettings.vue component's "Delete Account" button
func (uc *UserController) DeleteUserAccount(c *gin.Context) {
	id := c.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Extract user ID from JWT token
	tokenUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Check if the user ID from the token matches the user ID in the request parameters
	if tokenUserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You can only delete your own account"})
		return
	}

	if err := services.DeleteUserByID(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User account deleted"})
}
