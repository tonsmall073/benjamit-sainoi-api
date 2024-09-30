package db

import (
	"benjamit/db/benjamit"
	"benjamit/db/benjamit/models"
	"log"
)

func Migrate() {
	db, _ := benjamit.Connect()
	err := db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
