package db

import (
	"bjm/db/benjamit"
	"bjm/db/benjamit/models"
	"log"
)

func Migrate() {
	db, _ := benjamit.Connect()
	err := db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
