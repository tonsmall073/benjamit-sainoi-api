package benjamit

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	host := os.Getenv("BENJAMIT_CONNECT_POSTGRESQL_HOST")
	user := os.Getenv("BENJAMIT_CONNECT_POSTGRESQL_USER")
	pass := os.Getenv("BENJAMIT_CONNECT_POSTGRESQL_PASS")
	dbname := os.Getenv("BENJAMIT_CONNECT_POSTGRESQL_DBNAME")
	post := os.Getenv("BENJAMIT_CONNECT_POSTGRESQL_POST")
	sslmode := os.Getenv("BENJAMIT_CONNECT_POSTGRESQL_SSLMODE")
	timeZone := os.Getenv("BENJAMIT_POSTGRESQL_TIME_ZONE")
	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " port=" + post + " sslmode=" + sslmode + " TimeZone=" + timeZone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("[ERROR] failed to connect to database: %v\n", err)
		return nil, err
	}

	// log.Println("Connected to the PostgreSQL database successfully!")
	return db, nil
}

func ConnectClose(db *gorm.DB) {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("[ERROR] failed to get underlying DB: %v\n", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("[ERROR] failed to close database connection: %v\n", err)
			return
		}
	} else {
		log.Printf("[ERROR] failed to close database connection: %s\n", "there is no connection to the database")
		return
	}
}
