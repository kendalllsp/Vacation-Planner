package routes

import (
	"gorm.io/gorm"
)

// 1. Initliaze DB Router struct that can be passed around as object.
// 2. Create function that takes in gorm DB instance and creates instance of new DB object so gorm.DB can be passed

// 1.
type DBRouter struct {
	DB *gorm.DB
}

// 2.
func NewConnection(db *gorm.DB) DBRouter {
	return DBRouter{db}
}
