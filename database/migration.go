package database

import (
	"fmt"
	"log"

	"github.com/mereb/v4/models"
)

func RunMigrations() {
	db := GetDB()

	err := db.AutoMigrate(&models.Person{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Database migrations applied successfully!")
}
