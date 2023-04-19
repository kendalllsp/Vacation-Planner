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

// 1. Obtain private DB instance information.
// 2. Connect to the DB based on obtained information.
// 3. Check to see if there were any errors attempting to connect.
// 4. Use GORM to Auto-Migrate the tables to the given DB schemas within models.go
// 5. Print a string for successful DB connection.
// 6. Return the gorm.DB instance and potential error

func Connect() (*gorm.DB, error) {

	// 1.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT")

	// 2.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// 3.
	if err != nil {
		log.Fatalln(err)
	}

	// 4.
	db.AutoMigrate(&models.User{}, &models.SavedLocation{}, &models.SavedBusiness{})

	// 5.
	fmt.Println("successful connection")

	// 6.
	return db, err
}
