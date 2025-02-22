package services

import (
	"agoravote-app-backend/src/database"
	"agoravote-app-backend/src/models"

	"github.com/google/uuid"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) AuthenticateUser(user models.User) (string, error) {
	// Implement the logic to authenticate the user and return a token
	return "", nil
}

func (us *UserService) FetchUser(userID uuid.UUID) (*models.User, error) {
	// Implement the logic to fetch the user by ID from the database
	// For example:
	// var user models.User
	// if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
	//     return nil, err
	// }
	// return &user, nil

	return nil, nil
}

func (us *UserService) GetUserByID(userID uuid.UUID) (*models.User, error) {
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
	user.ID = uuid.New()
	return database.DB.Create(user).Error
}

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
