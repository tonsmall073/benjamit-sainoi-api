package utils

import (
	con "bjm/db/benjamit"
	"bjm/db/benjamit/models"
	"bjm/utils/enums"
	"log"
	"os"
	"strconv"
	"time"
)

func deleteOldLogs(interfaceType enums.InterfaceTypeEnum) {
	db, dbErr := con.Connect()
	defer con.ConnectClose(db)
	if dbErr != nil {
		log.Printf("[ERROR] failed to connect to database: %v\n", dbErr)
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("[ERROR] failed to get underlying DB: %v\n", err)
		} else if pingErr := sqlDB.Ping(); pingErr != nil {
			log.Printf("[WARN] database connection lost: %v. Retrying...\n", pingErr)
		} else {
			getEnv := os.Getenv("LOG_CLEANING_DAY")
			day := 30
			if getEnv != "" {
				convInt, convIntErr := strconv.Atoi(getEnv)
				if convIntErr == nil {
					day = convInt
				}
			}
			thirtyDaysAgo := time.Now().AddDate(0, 0, -day)
			if err := db.Unscoped().Where("created_at < ? AND interface_type = ?", thirtyDaysAgo, interfaceType).Delete(&models.ApiTransactionLog{}).Error; err != nil {
				log.Printf("[ERROR] failed to delete old logs: %v\n", err)
			} else {
				log.Printf("[INFO] deleted logs older than %d days\n", day)
			}
		}
	}
}

func LogCleanupTask(interfaceType enums.InterfaceTypeEnum) {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		<-ticker.C
		deleteOldLogs(interfaceType)
	}
}
