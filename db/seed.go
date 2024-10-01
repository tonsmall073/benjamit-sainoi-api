package db

import (
	"bjm/db/benjamit"
	"bjm/db/benjamit/seeds"
	"log"
)

func Seed() {
	db, _ := benjamit.Connect()
	result := db.Create(seeds.Prefix())
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}
