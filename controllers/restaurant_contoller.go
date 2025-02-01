package controllers

import (
	"net/http"

	"github.com/copausina/TheEats/db"
	"github.com/copausina/TheEats/models"
	"github.com/gin-gonic/gin"
)

// Read
// Get all restaurants
func GetRestaurants(context *gin.Context) {
	var restaurants []models.Restaurant
	db.GetDB().Find(&restaurants)
	context.JSON(http.StatusOK, restaurants)
}

// Get restaurant by ID
func GetRestaurantByID(context *gin.Context) {
	id := context.Param("id") // Get ID from URL

	var restaurant models.Restaurant
	if err := db.GetDB().First(&restaurant, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	context.JSON(http.StatusOK, restaurant)
}

// Create
// Add a new restaurant
func AddRestaurant(context *gin.Context) {
	var restaurant models.Restaurant
	if err := context.ShouldBindJSON(&restaurant); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.GetDB().Create(&restaurant)
	context.JSON(http.StatusOK, restaurant)
}

// Update
// Update a restaurant
func UpdateRestaurant(context *gin.Context) {
	var restaurant models.Restaurant
	id := context.Param("id")

	// Check if restaurant exists
	if err := db.GetDB().First(&restaurant, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	// Bind new data
	if err := context.ShouldBindJSON(&restaurant); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated restaurant
	db.GetDB().Save(&restaurant)
	context.JSON(http.StatusOK, restaurant)
}

// Delete
// Delete a restaurant
func DeleteRestaurant(context *gin.Context) {
	var restaurant models.Restaurant
	id := context.Param("id")

	// Check if restaurant exists
	if err := db.GetDB().First(&restaurant, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	// Delete the restaurant
	db.GetDB().Delete(&restaurant)
	context.JSON(http.StatusOK, gin.H{"message": "Restaurant deleted successfully"})
}
