package routes

import (
	"gorm.io/gorm"
)

// Initalizing the Database Router struct so that the database connection can be passed as an object
type DBRouter struct {
	DB *gorm.DB
}

// Takes in database and creates object for it to be passed
func NewConnection(db *gorm.DB) DBRouter {
	return DBRouter{db}
}
