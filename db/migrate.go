package db

import (
	"bjm/db/benjamit"
	"bjm/db/benjamit/models"
	"fmt"
	"log"
	"os"
	"strings"

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

	errs := createEnums(db)
	for err := range errs {
		log.Printf("[ERROR] failed to create enum type : %v\n", err)
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

	log.Println("[INFO] the migration is complete.")
}

func createEnums(db *gorm.DB) []error {
	var errors []error

	enumDefinitions := map[string][]string{
		"role_enum": {
			string(models.USER),
			string(models.ADMIN),
		},
		"message_type_enum": {
			string(models.TEXT),
			string(models.IMAGE),
			string(models.EMOJI),
		},
	}

	for name, values := range enumDefinitions {
		joinedValues := strings.Join(values, "', '")
		query := fmt.Sprintf("CREATE TYPE %s AS ENUM ('%s');", name, joinedValues)
		if err := db.Exec(query).Error; err != nil {
			errors = append(errors, err)
		}
	}

	return errors
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
