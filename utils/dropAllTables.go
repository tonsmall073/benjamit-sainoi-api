package utils

import (
	"log"

	"gorm.io/gorm"
)

func DropAllTables(db *gorm.DB) error {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Fatal("failed to drop all tables:", err)
		return err
	}
	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			log.Fatal("failed to drop all tables:", err)
			return err
		}
	}
	log.Println("Drop all tables Complete.")
	return nil
}
