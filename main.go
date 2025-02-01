package main

import (
	//"net/http"
	"fmt"
	"log"
	"os"

	"github.com/copausina/TheEats/db"
	"github.com/copausina/TheEats/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// load enviroment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"}) //only trust localhost for now

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db.InitDB(dsn)        // Connect database
	routes.SetupRoutes(r) // Setup API routes

	r.Run(":8080") // Start server
}
