package services

import (
	"agoravote-app-backend/src/database"
	"agoravote-app-backend/src/models"
	"errors"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) AuthenticateUser(user models.User) (string, error) {
	var existingUser models.User
	if err := database.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&existingUser).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate a token (this is a placeholder, implement your own token generation logic)
	token := "generated-jwt-token"
	return token, nil
}

func (us *UserService) FetchUser(userId string) (models.User, error) {
	var user models.User
	if err := database.DB.First(&user, "id = ?", userId).Error; err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (us *UserService) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string, user *models.User) error {
	return database.DB.Where("email = ?", email).First(user).Error
}

func CreateUser(user *models.User) error {
	// Remove the role field from the user creation logic
	return database.DB.Create(user).Error
}
