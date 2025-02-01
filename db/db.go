package db

import (
	"log"

	"github.com/copausina/TheEats/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB connection
var DB *gorm.DB

func InitDB(dsn string) {
	var errDB error
	DB, errDB = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatal("Failed to connect to database:\n", errDB)
	}

	DB.AutoMigrate(&models.Restaurant{}) // Auto-migrate tables
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
