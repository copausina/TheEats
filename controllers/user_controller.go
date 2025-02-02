package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/copausina/TheEats/db"
	"github.com/copausina/TheEats/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Secret key for JWT
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Generate a JWT token
func generateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	return token.SignedString(jwtSecret)
}

// Register a new user
func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Ensure role is either "user" or "admin"
	if user.Role != "admin" {
		user.Role = "user"
	}

	// Save the user
	if err := db.GetDB().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login user
func Login(c *gin.Context) {
	var requestUser models.User
	var user models.User

	if err := c.ShouldBindJSON(&requestUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user by email
	if err := db.GetDB().Where("email = ?", requestUser.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare hashed passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestUser.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("auhthorization", token, 86400, "/", "", false, true) // 86400 = 24 hours

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// Logout and clear JWT cookie by setting setting its value to "" and expiration time to -1
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// Get all users
func GetUsers(context *gin.Context) {
	var users []models.User
	db.GetDB().Find(&users)
	context.JSON(http.StatusOK, users)
}
