package routes

import (
	"github.com/copausina/TheEats/controllers"
	"github.com/copausina/TheEats/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth") //All
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
		authRoutes.POST("/logout", controllers.Logout)
		//authRoutes.Use(middlewares.AdminMiddleware()) // Require admin access
		authRoutes.GET("/", controllers.GetUsers) // Only admins
	}

	restaurantRoutes := router.Group("/api/restaurants")
	{
		restaurantRoutes.GET("/", controllers.GetRestaurants)       // Public
		restaurantRoutes.GET("/:id", controllers.GetRestaurantByID) // Public

		restaurantRoutes.Use(middlewares.UserMiddleware())    // Require login
		restaurantRoutes.POST("/", controllers.AddRestaurant) // Only logged-in users

		restaurantRoutes.Use(middlewares.AdminMiddleware())           // Require admin access
		restaurantRoutes.DELETE("/:id", controllers.DeleteRestaurant) // Only admins
	}
}
