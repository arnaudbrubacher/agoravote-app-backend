package controllers

import (
	"agoravote-app-backend/src/config"
	"agoravote-app-backend/src/models"
	"agoravote-app-backend/src/services"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Claims
// Data structure for JWT token payload containing user identification
type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}

// GenerateJWT
// Token generation utility that creates signed JWT tokens for users
func generateJWT(userID uuid.UUID) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JWTKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Login
// Authentication endpoint that validates credentials and returns JWT
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbUser models.User
	if err := services.GetUserByEmail(user.Email, &dbUser); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if dbUser.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	tokenString, err := generateJWT(dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "userId": dbUser.ID})
}

// Signup
// Registration endpoint that creates new users and returns JWT
func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := generateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "userId": user.ID})
}
