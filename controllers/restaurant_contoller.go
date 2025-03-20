package controllers

import (
	"net/http"
	"strconv"

	"github.com/copausina/TheEats/db"
	"github.com/copausina/TheEats/handlers"
	"github.com/copausina/TheEats/models"
	"github.com/gin-gonic/gin"
)

// Read
// Get all restaurants
func GetRestaurants(c *gin.Context) {
	var restaurants []models.Restaurant
	db.GetDB().Find(&restaurants)
	c.JSON(http.StatusOK, restaurants)
}

// Get restaurant by ID
func GetRestaurantByID(c *gin.Context) {
	id := c.Param("id") // Get ID from URL

	var restaurant models.Restaurant
	if err := db.GetDB().First(&restaurant, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	c.JSON(http.StatusOK, restaurant)
}

// Create
// Add a new restaurant
func AddRestaurant(c *gin.Context) {
	// Parse form data
	name := c.PostForm("name")
	address := c.PostForm("address")
	cuisine := c.PostForm("cuisine")
	rating64, err := strconv.ParseFloat(c.PostForm("rating"), 32) // convert to float32
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rating. Please provide a number."})
	}
	rating := float32(rating64) // b/c ParseFloat returns float64 even w/ 32 parameter

	// Get image file
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"}) // TODO: make not required later
		return
	}
	// Open file to get multipart.File
	fileData, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image file"})
		return
	}
	defer fileData.Close()

	// Create temporary restaurant w/o image
	restaurant := models.Restaurant{Name: name, Address: address, Cuisine: cuisine, Rating: rating}
	result := db.GetDB().Create(&restaurant)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create restaurant"})
		return
	}

	// Upload the image and get the URL
	imageURL, err := handlers.UploadRestaurantImage(fileData, name) // image filename will be same as name of restaurant
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	// Update the restaurant with the image URL
	restaurant.ImageURL = imageURL
	db.GetDB().Save(&restaurant)

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"restaurant": restaurant})

	// var restaurant models.Restaurant
	// if err := c.ShouldBindJSON(&restaurant); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// db.GetDB().Create(&restaurant)
	// c.JSON(http.StatusOK, restaurant)
}

// Update
// Update a restaurant
func UpdateRestaurant(c *gin.Context) {
	var restaurant models.Restaurant
	id := c.Param("id")

	// Check if restaurant exists
	if err := db.GetDB().First(&restaurant, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	// Parse form data
	name := c.PostForm("name")
	address := c.PostForm("address")
	cuisine := c.PostForm("cuisine")
	if name != "" {
		restaurant.Name = name
	}
	if address != "" {
		restaurant.Address = address
	}
	if cuisine != "" {
		restaurant.Cuisine = cuisine
	}
	ratingStr := c.PostForm("rating")
	if ratingStr != "" {
		rating64, err := strconv.ParseFloat(c.PostForm("rating"), 32) // convert to float32
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rating. Please provide a number."})
		}
		restaurant.Rating = float32(rating64) // b/c ParseFloat returns float64 even w/ 32 parameter
	}

	// Get image file
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"}) // TODO: make not required later
		return
	}
	// Open file to get multipart.File
	fileData, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image file"})
		return
	}
	defer fileData.Close()

	// Before images were implemented
	// // Bind new data
	// if err := c.ShouldBindJSON(&restaurant); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// Save the updated restaurant
	db.GetDB().Save(&restaurant)
	c.JSON(http.StatusOK, restaurant)
}

// Delete
// Delete a restaurant
func DeleteRestaurant(c *gin.Context) {
	var restaurant models.Restaurant
	id := c.Param("id")

	// Check if restaurant exists
	if err := db.GetDB().First(&restaurant, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	// Delete image from supabase, then restaurant from Postgress
	imageName := []string{restaurant.Name} // for now only 1 image per restaurant
	imageDeleteMsg, err := handlers.DeleteRestaurantImage(imageName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	db.GetDB().Delete(&restaurant)

	// Respond with success message including both deletion messages
	c.JSON(http.StatusOK, gin.H{
		"message":       "Restaurant deleted successfully",
		"imageDeletion": imageDeleteMsg,
	})
}

func UploadRestaurantImageHandler(c *gin.Context) {
	// Get the file from the request
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}
	defer file.Close()

	// Upload image to Supabase Storage
	imageURL, err := handlers.UploadRestaurantImage(file, header.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	restaurantID := c.PostForm("restaurant_id")
	var restaurant models.Restaurant

	// Find restaurant by ID
	if err := db.GetDB().First(&restaurant, restaurantID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	// Update restaurant image URL
	restaurant.ImageURL = imageURL
	if err := db.GetDB().Save(&restaurant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update restaurant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "image_url": imageURL})
}
