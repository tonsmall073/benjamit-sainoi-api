package db

import (
	"bjm/db/benjamit"
	"bjm/db/benjamit/models"
	"log"

	"gorm.io/gorm"
)

func Migrate() {
	db, _ := benjamit.Connect()
	if err := createUUIDExtension(db); err != nil {
		log.Printf("[ERROR] failed to create uuid-ossp extension: %v\n", err)
		return
	}

	err := db.AutoMigrate(
		&models.Prefix{},
		&models.User{},
		&models.Log{},
	)

	if err != nil {
		log.Printf("[ERROR] failed to migrate database: %v\n", err)
		return
	}

	log.Println("The migration is complete.")
}

func createUUIDExtension(db *gorm.DB) error {
	return db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
}
