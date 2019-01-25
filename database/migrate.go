package database

import (
	"examplego/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Note{})
	db.AutoMigrate(&models.URL{})
	// set up foreign keys
	db.Model(&models.Note{})
	db.Model(&models.URL{})
	fmt.Println("Auto Migration has been processed")
}
