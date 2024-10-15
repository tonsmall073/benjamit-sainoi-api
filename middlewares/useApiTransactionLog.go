package middlewares

import (
	model "bjm/db/benjamit/models"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	con "bjm/db/benjamit"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	cleanupOnce sync.Once
)

func deleteOldLogs(db *gorm.DB) {
	getEnv := os.Getenv("LOG_CLEANING_DAY")
	day := 30
	if getEnv != "" {
		convInt, convIntErr := strconv.Atoi(getEnv)
		if convIntErr == nil {
			day = convInt
		}
	}
	thirtyDaysAgo := time.Now().AddDate(0, 0, -day)
	if err := db.Where("created_at < ?", thirtyDaysAgo).Delete(&model.ApiTransactionLog{}).Error; err != nil {
		log.Printf("[ERROR] failed to delete old logs: %v\n", err)
	} else {
		log.Printf("[INFO] Deleted logs older than %d days\n", day)
	}
}

func logCleanupTask(db *gorm.DB) {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		<-ticker.C
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("[ERROR] failed to get underlying DB: %v\n", err)
			continue
		}

		if pingErr := sqlDB.Ping(); pingErr != nil {
			log.Printf("[WARN] Database connection lost: %v. Retrying...\n", pingErr)
			continue
		}
		deleteOldLogs(db)
	}
}

func logMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		var requestBody string
		bodyBytes := c.Body()
		requestBody = string(bodyBytes)
		rawHeaders := c.GetReqHeaders()
		headers := make(map[string]string)
		for key, values := range rawHeaders {
			if len(values) > 0 {
				headers[key] = values[0]
			}
		}
		headersJson, headersJsonErr := json.Marshal(headers)
		if headersJsonErr != nil {
			log.Printf("[ERROR] failed to marshal headers: %v\n", headersJsonErr)
		}
		origin := c.Get("Origin")
		c.Set("Access-Control-Allow-Origin", origin)

		if strings.Contains(c.Path(), "/events") || strings.Contains(c.Path(), "events/") {
			log.Printf("[INFO] SSE request for path: %s\n", c.Path())
		} else {
			responseLog := model.ApiTransactionLog{
				Path:         c.Path(),
				Method:       c.Method(),
				ContentType:  c.Get("Content-Type"),
				StatusCode:   c.Response().StatusCode(),
				ResponseBody: string(c.Response().Body()),
				RequestBody:  requestBody,
				Headers:      string(headersJson),
				Origin:       origin,
			}
			if err := db.Create(&responseLog).Error; err != nil {
				log.Printf("[ERROR] failed to log error: %v\n", err)
			} else {
				log.Printf("[INFO] Logging request for path: %s\n", c.Path())
			}
		}

		return err
	}
}

func UseApiTransactionLog(app *fiber.App) fiber.Router {
	db, dbErr := con.Connect()
	if dbErr != nil {
		log.Printf("[ERROR] failed to connect to database: %v\n", dbErr)
		return app.Group("")
	}

	cleanupOnce.Do(func() {
		go logCleanupTask(db)
	})

	return app.Group("", logMiddleware(db))
}
