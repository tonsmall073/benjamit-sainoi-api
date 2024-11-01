package middlewares

import (
	model "bjm/db/benjamit/models"
	"bjm/utils"
	"bjm/utils/enums"
	"encoding/json"
	"log"
	"regexp"
	"sync"

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
					Path:          path,
					Method:        method,
					ContentType:   contentType,
					StatusCode:    statusCode,
					ResponseBody:  responseBody,
					RequestBody:   requestBody,
					Headers:       string(headersJson),
					Origin:        origin,
					InterfaceType: enums.HTTP,
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
		go utils.LogCleanupTask(enums.HTTP)
	})

	return app.Group("", logMiddleware())
}
