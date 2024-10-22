package jwt

import (
	"bjm/utils"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
)

func UseUserGuard(c *fiber.Ctx) error {
	errorChan := make(chan utils.ErrorResponseModel)
	resultChan := make(chan error)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			errorChan <- utils.ErrorResponseModel{
				MessageDesc: utils.HttpStatusCodes[401],
				StatusCode:  401,
			}
			return
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			errorChan <- utils.ErrorResponseModel{
				MessageDesc: utils.HttpStatusCodes[401],
				StatusCode:  401,
			}
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		res, err := VerifyToken(token)
		if err != nil {
			errorChan <- utils.ErrorResponseModel{
				MessageDesc: utils.HttpStatusCodes[401],
				StatusCode:  401,
			}
			return
		}
		if res["role"].(string) != "USER" && res["role"].(string) != "ADMIN" {
			errorChan <- utils.ErrorResponseModel{
				MessageDesc: utils.HttpStatusCodes[403],
				StatusCode:  403,
			}
			return
		}
		resultChan <- c.Next()
	}()

	go func() {
		wg.Wait()
		close(errorChan)
		close(resultChan)
	}()

	select {
	case err := <-errorChan:
		return utils.FiberResponseErrorJson(c, err.MessageDesc, err.StatusCode)
	case result := <-resultChan:
		return result
	}
}

func UseAdminGuard(c *fiber.Ctx) error {
	errorChan := make(chan utils.ErrorResponseModel)
	resultChan := make(chan error)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			errorChan <- utils.ErrorResponseModel{
				MessageDesc: utils.HttpStatusCodes[401],
				StatusCode:  401,
			}
			return
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			errorChan <- utils.ErrorResponseModel{
				MessageDesc: utils.HttpStatusCodes[401],
				StatusCode:  401,
			}
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		res, err := VerifyToken(token)
		if err != nil {
			errorChan <- utils.ErrorResponseModel{
				MessageDesc: utils.HttpStatusCodes[401],
				StatusCode:  401,
			}
			return
		}
		if res["role"].(string) != "ADMIN" {
			errorChan <- utils.ErrorResponseModel{
				MessageDesc: utils.HttpStatusCodes[403],
				StatusCode:  403,
			}
			return
		}
		resultChan <- c.Next()
	}()

	go func() {
		wg.Wait()
		close(errorChan)
		close(resultChan)
	}()

	select {
	case err := <-errorChan:
		return utils.FiberResponseErrorJson(c, err.MessageDesc, err.StatusCode)
	case result := <-resultChan:
		return result
	}
}
