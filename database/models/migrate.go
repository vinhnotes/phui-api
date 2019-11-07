package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Team{})

	db.Model(&Team{}).AddForeignKey("owner_id", "users(id)", "CASCADE", "CASCADE")
	fmt.Println("Auto Migration has beed processed")
}
