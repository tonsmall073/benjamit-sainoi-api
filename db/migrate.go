package db

import (
	"bjm/db/benjamit"
	"bjm/db/benjamit/models"
	"log"
	"os"

	"gorm.io/gorm"
)

func Migrate() {
	migrateBenjamitDatabase()
}

func migrateBenjamitDatabase() {
	db, _ := benjamit.Connect()
	if err := createUUIDExtension(db); err != nil {
		log.Printf("[ERROR] failed to create uuid-ossp extension: %v\n", err)
		return
	}
	if err := setTimeZone(db); err != nil {
		log.Printf("[ERROR] failed to set time zone: %v\n", err)
		return
	}

	err := db.AutoMigrate(
		&models.Prefix{},
		&models.User{},
		&models.ApiTransactionLog{},
		&models.Chat{},
		&models.Notification{},
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

func setTimeZone(db *gorm.DB) error {
	defaultStr := "Asia/Bangkok"

	if getTimeZone := os.Getenv("BENJAMIT_POSTGRESQL_TIME_ZONE"); getTimeZone != "" {
		defaultStr = getTimeZone
	}
	err := db.Exec("SET timezone = '" + defaultStr + "';").Error
	var timezone string
	db.Raw("SHOW TIMEZONE").Scan(&timezone)
	log.Println("Current Timezone:", timezone)
	return err
}
