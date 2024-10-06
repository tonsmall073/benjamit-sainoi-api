package db

import (
	"bjm/db/benjamit"
	"bjm/db/benjamit/seeds"
	"log"
	"reflect"

	"gorm.io/gorm"
)

func Seeder() {
	db, err := benjamit.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	data := []interface{}{
		seeds.Prefix(),
		seeds.User(),
		// seeds.ApiTransactionLog(),
	}
	insertData(db, data)
}

func insertData(db *gorm.DB, data []interface{}) {
	for _, item := range data {
		modelName := reflect.TypeOf(item).Elem().Name()
		if err := db.Create(item).Error; err != nil {
			log.Printf("[ERROR] insert data '%s' fail error : '%s'\n", modelName, err.Error())
		} else {
			log.Printf("insert data '%s' success.\n", modelName)
		}
	}
	log.Println("finished sprinkling the seeds.")
}
