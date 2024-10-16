package utils

import (
	"log"

	"gorm.io/gorm"
)

func DropAllTables(db *gorm.DB) error {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Fatal("[ERROR] failed to drop all tables:", err)
		return err
	}
	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			log.Fatal("[ERROR] failed to drop all tables:", err)
			return err
		}
	}
	if err := dropAllEnumTypes(db); err != nil {
		log.Fatal("[ERROR] failed to drop all enum types:", err)
		return err
	}
	log.Println("[INFO] drop all tables complete.")
	return nil
}

func dropAllEnumTypes(db *gorm.DB) error {
	var enumTypes []string

	// ค้นหา enum types ทั้งหมด
	if err := db.Raw(`
		SELECT typname 
		FROM pg_type 
		WHERE typtype = 'e';`).Scan(&enumTypes).Error; err != nil {
		return err
	}

	// ลบ enum types ทั้งหมด
	for _, enumType := range enumTypes {
		if err := db.Exec("DROP TYPE IF EXISTS " + enumType).Error; err != nil {
			return err
		}
	}

	return nil
}
