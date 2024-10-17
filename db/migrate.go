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
	for _, err := range errs {
		log.Printf("[ERROR] failed to create enum type : %v\n", err)
		return
	}

	err := db.AutoMigrate(
		&models.Prefix{},
		&models.User{},
		&models.ApiTransactionLog{},
		&models.Chat{},
		&models.Notification{},
		&models.Product{},
		&models.UnitType{},
		&models.ProductSelling{},
		&models.IncomeAndExpense{},
	)

	if err != nil {
		log.Printf("[ERROR] failed to migrate database: %v\n", err)
		return
	}

	log.Println("[INFO] the migration is complete.")
}

func createEnums(db *gorm.DB) []error {
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
		"transaction_type_enum": {
			string(models.DEBIT),
			string(models.CREDIT),
		},
		"entry_source_enum": {
			string(models.MANUAL),
			string(models.SYSTEM),
		},
	}

	return createEnumCondition(db, enumDefinitions)
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
	log.Println("[INFO] current timezone:", timezone)
	return err
}

func createEnumCondition(db *gorm.DB, enumDefinitions map[string][]string) []error {
	var errors []error
	for name, values := range enumDefinitions {
		var exists bool
		if err := db.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = ?)", name).Scan(&exists).Error; err != nil {
			errors = append(errors, err)
		}

		if exists {
			for _, value := range values {
				var eumExists bool
				query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM pg_enum WHERE enumlabel = '%s' AND enumtypid = (SELECT oid FROM pg_type WHERE typname = '%s'));", value, name)
				if err := db.Raw(query).Scan(&eumExists).Error; err != nil {
					errors = append(errors, err)
					eumExists = false
				}
				if !eumExists {
					if err := db.Exec(fmt.Sprintf("ALTER TYPE %s ADD VALUE '%s';", name, value)).Error; err != nil {
						errors = append(errors, err)
					}
				}
			}
		} else {
			joinedValues := strings.Join(values, "', '")
			query := fmt.Sprintf("CREATE TYPE %s AS ENUM ('%s');", name, joinedValues)
			if err := db.Exec(query).Error; err != nil {
				errors = append(errors, err)
			}
		}
	}
	return errors
}
