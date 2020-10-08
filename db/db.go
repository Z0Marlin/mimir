package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init initiliaes the database
func Init() {
	connectDatabase()
}

// GetDatabase returns a pointer to the databse handle
func GetDatabase() *gorm.DB {
	return db
}

// D is a shothand for GetDatabase
func D() *gorm.DB {
	return db
}

func connectDatabase() {
	databse, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db!")
	}

	db = databse
}
