package routes

import (
	"github.com/copausina/TheEats/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	restaurantRoutes := router.Group("/api/restaurants")
	{
		restaurantRoutes.GET("/", controllers.GetRestaurants)
		restaurantRoutes.GET("/:id", controllers.GetRestaurantByID)
		restaurantRoutes.POST("/", controllers.AddRestaurant)
		restaurantRoutes.PUT("/:id", controllers.UpdateRestaurant)
		restaurantRoutes.DELETE("/:id", controllers.DeleteRestaurant)

	}
}
