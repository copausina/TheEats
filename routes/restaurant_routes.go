package routes

import (
	"net/http"

	"github.com/copausina/TheEats/db"
	"github.com/copausina/TheEats/models"
	"github.com/gin-gonic/gin"
)

// Get all restaurants
func GetRestaurants(context *gin.Context) {
	var restaurants []models.Restaurant
	db.GetDB().Find(&restaurants)
	context.JSON(http.StatusOK, restaurants)
}

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

func SetupRoutes(router *gin.Engine) {
	restaurantRoutes := router.Group("/api/restaurants")
	{
		restaurantRoutes.GET("/", GetRestaurants)
		restaurantRoutes.POST("/", AddRestaurant)
	}
}
