package services

import (
	"agoravote-app-backend/src/database"
	"agoravote-app-backend/src/models"

	"github.com/google/uuid"
)

// UserService
// Business logic layer for user-related operations
// Frontend: Used by all authenticated components requiring user data
type UserService struct{}

// NewUserService
// Creates a new instance of UserService
// Frontend: Used internally during app initialization, no direct component usage
func NewUserService() *UserService {
	return &UserService{}
}

// AuthenticateUser
// Validates user credentials and generates authentication token
// Frontend: Called by LoginForm.vue component when clicking "Sign In" button
func (us *UserService) AuthenticateUser(user models.User) (string, error) {
	// Implement the logic to authenticate the user and return a token
	return "", nil
}

// FetchUser
// Retrieves user details from database
// Frontend: Called by UserProfile.vue when loading user profile page
func (us *UserService) FetchUser(userID uuid.UUID) (*models.User, error) {
	// Implement the logic to fetch the user by ID from the database
	return nil, nil
}

// GetUserByID
// Database query to retrieve user by UUID
// Frontend: Called by UserCard.vue when displaying user information in groups
func (us *UserService) GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail
// Looks up user by email address
// Frontend: Called by SignupForm.vue to check if email is already registered
func GetUserByEmail(email string, user *models.User) error {
	return database.DB.Where("email = ?", email).First(user).Error
}

// CreateUser
// Creates new user account in database
// Frontend: Called by SignupForm.vue when clicking "Create Account" button
func CreateUser(user *models.User) error {
	user.ID = uuid.New()
	return database.DB.Create(user).Error
}

// DeleteUserByID
// Removes user account from database
// Frontend: Called by AccountSettings.vue when clicking "Delete Account" button
func DeleteUserByID(userID uuid.UUID) error {
	if err := database.DB.Delete(&models.User{}, "id = ?", userID).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID retrieves a user by their ID from the database
func GetUserByID(userID uuid.UUID, user *models.User) error {
	if err := database.DB.Where("id = ?", userID).First(user).Error; err != nil {
		return err
	}
	return nil
}
