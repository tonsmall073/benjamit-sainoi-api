package middlewares

import (
	model "bjm/db/benjamit/models"
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	con "bjm/db/benjamit"

	"github.com/gofiber/fiber/v2"
)

var (
	cleanupOnce sync.Once
)

func isValidSSEPath(path string) bool {
	// ใช้ regex เพื่อให้ "events" เป็น path ที่ถูกต้อง
	re := regexp.MustCompile(`^(\/.*\/)?events(\/.*)?(\?.*)?$`)
	return re.MatchString(path)
}

func deleteOldLogs() {
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
			if err := db.Where("created_at < ?", thirtyDaysAgo).Delete(&model.ApiTransactionLog{}).Error; err != nil {
				log.Printf("[ERROR] failed to delete old logs: %v\n", err)
			} else {
				log.Printf("[INFO] deleted logs older than %d days\n", day)
			}
		}
	}
}

func logCleanupTask() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		<-ticker.C
		deleteOldLogs()
	}
}

func logMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		if isValidSSEPath(c.Path()) {
			log.Printf("[INFO] sse request for path: %s\n", c.Path())
		} else {
			rawHeaders := c.GetReqHeaders()
			headers := make(map[string]string)
			path := c.Path()
			method := c.Method()
			requestBody := string(c.Body())
			responseBody := string(c.Response().Body())
			statusCode := c.Response().StatusCode()
			contentType := c.Get("Content-Type")
			origin := c.Get("Origin")
			c.Set("Access-Control-Allow-Origin", origin)

			go func() {
				for key, values := range rawHeaders {
					if len(values) > 0 {
						headers[key] = values[0]
					}
				}

				headersJson, headersJsonErr := json.Marshal(headers)
				if headersJsonErr != nil {
					log.Printf("[ERROR] logging json marshal error: %v\n", headersJsonErr)
				}

				responseLog := model.ApiTransactionLog{
					Path:         path,
					Method:       method,
					ContentType:  contentType,
					StatusCode:   statusCode,
					ResponseBody: responseBody,
					RequestBody:  requestBody,
					Headers:      string(headersJson),
					Origin:       origin,
				}
				db, dbErr := con.Connect()
				defer con.ConnectClose(db)
				if dbErr != nil {
					log.Printf("[ERROR] failed to connect to database: %v\n", dbErr)
				} else {
					if err := db.Create(&responseLog).Error; err != nil {
						log.Printf("[ERROR] logging recording errors: %v\n", err)
					} else {
						log.Printf("[INFO] logging request for path: %s\n", c.Path())
					}
				}
			}()
		}
		return err
	}
}

func UseApiTransactionLog(app *fiber.App) fiber.Router {
	cleanupOnce.Do(func() {
		go logCleanupTask()
	})

	return app.Group("", logMiddleware())
}
