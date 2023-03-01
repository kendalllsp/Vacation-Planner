package database

import (
	"fmt"
	"log"
	"vacation-planner/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Error return value to Connect function for possible errors being thrown while connecting
func Connect() (*gorm.DB, error) {

	//default code found in GormPG docs
	dsn := "host=24.199.69.57 user=dev password=K@Gf$+$n?EBV^K3% dbname=dev port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//test migrating User struct
	db.AutoMigrate(&models.User{}, &models.SavedLocation{})

	fmt.Println("successful connection")

	// Return potential error or nil as well as DB
	return db, err
}
