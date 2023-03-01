package models

import (
	"gorm.io/gorm"
)

type Business struct {
    Name        string  `json: "name"`
    Price       string  `json: "price"`
    Rating      float64 `json: "rating"`
    Location    string  `json: "location"`
    Type        string  `json: "type"`
}

// Removed gorm.Model for Destination because it won't be saved, only the name of the destination
// Removed ImageLink because it's not being used yet
type Destination struct {
    Location        [3]string       `json: "location"`
    Restaurants     [10]Business    `json: "restaurants"`
    Entertainment   [10]Business    `json: "entertainment"`
    Shopping        [10]Business    `json: "shopping"`
}

type SavedLocation struct {
    Email       string  `json: "email"`
    Location    string  `json: "location"`
}

// introducing user type for database
type User struct {
    gorm.Model
    Email           string      `json: "email"          gorm: "primaryKey"`
    Password        string      `json: "password"`
}