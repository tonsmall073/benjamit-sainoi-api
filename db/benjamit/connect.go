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
	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " port=" + post + " sslmode=" + sslmode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}

	// log.Println("Connected to the PostgreSQL database successfully!")
	return db, nil
}
