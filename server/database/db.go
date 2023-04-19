package database

import (
	"fmt"
	"log"
	"os"
	"vacation-planner/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 1. Declare specific

// Error return value to Connect function for possible errors being thrown while connecting
func Connect() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}

	//default code found in GormPG docs
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// Auto Migrating User and SavedLocation structs
	db.AutoMigrate(&models.User{}, &models.SavedLocation{}, &models.SavedBusiness{})

	// Visuals for Successful Connection
	fmt.Println("successful connection")

	// Return potential error or nil as well as DB
	return db, err
}
