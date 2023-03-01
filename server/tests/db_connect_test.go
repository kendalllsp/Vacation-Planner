package tests

import (
    "testing"
	"vacation-planner/database"
	"vacation-planner/models"
)

// Test function for DB connection
func TestDBConnection(t *testing.T) {
    db, err := database.Connect()
	if err != nil {
		t.Fatalf("Test connection failed, could not successfully connect to database.")
	} else {
		// Sample code for DB to be used.
		db.AutoMigrate(&models.User{})
	}
}
