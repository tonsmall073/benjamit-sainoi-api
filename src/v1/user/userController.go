package user

import (
	auth "bjm/auth/jwt"
	db "bjm/db/benjamit"
	"bjm/src/v1/user/dto"
	"bjm/utils"
	"sync"

	"github.com/gofiber/fiber/v2"
)

// @Tags User
// @Description Create a user by filling in user information.
// @Accept json
// @Produce json
// @Param input body dto.CreateUserRequestModel true "create user request"
// @Success 201 {object} dto.CreateUserResponseModel "created"
// @Failure 400 {object} utils.ErrorResponseModel "invalid input"
// @Failure 500 {object} utils.ErrorResponseModel "internal server error"
// @Router /v1/user/create [post]
func createUser(c *fiber.Ctx) error {
	resultChan := make(chan *dto.CreateUserResponseModel)
	errorChan := make(chan utils.ErrorResponseModel)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		reqModel := &dto.CreateUserRequestModel{}
		err := c.BodyParser(reqModel)
		if err != nil {
			errorChan <- utils.ErrorResponseModel{MessageDesc: err.Error(), StatusCode: 400}
			return
		}
		context, contextErr := db.Connect()
		defer db.ConnectClose(context)
		if contextErr != nil {
			errorChan <- utils.ErrorResponseModel{MessageDesc: contextErr.Error(), StatusCode: 500}
			return
		}

		resModel := &dto.CreateUserResponseModel{}
		service := &UserService{context}
		serviceRes := service.CreateUser(reqModel, resModel)
		resultChan <- serviceRes
	}()

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case err := <-errorChan:
		return utils.FiberResponseErrorJson(c, err.MessageDesc, err.StatusCode)
	case result := <-resultChan:
		return utils.FiberResponseJson(c, result, result.StatusCode)
	}
}

// @Tags User
// @Description Login user with username and password
// @Accept json
// @Produce json
// @Param input body dto.LoginRequestModel true "login request"
// @Success 200 {object} dto.LoginResponseModel "ok"
// @Failure 400 {object} utils.ErrorResponseModel "invalid input"
// @Failure 500 {object} utils.ErrorResponseModel "internal server error"
// @Router /v1/user/login [post]
func login(c *fiber.Ctx) error {
	resultChan := make(chan *dto.LoginResponseModel)
	errorChan := make(chan utils.ErrorResponseModel)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		reqModel := &dto.LoginRequestModel{}
		err := c.BodyParser(reqModel)
		if err != nil {
			errorChan <- utils.ErrorResponseModel{MessageDesc: err.Error(), StatusCode: 400}
			return
		}
		context, contextErr := db.Connect()
		defer db.ConnectClose(context)
		if contextErr != nil {
			errorChan <- utils.ErrorResponseModel{MessageDesc: contextErr.Error(), StatusCode: 500}
			return
		}

		resModel := &dto.LoginResponseModel{}
		service := &UserService{context}
		serviceRes := service.Login(reqModel, resModel)
		resultChan <- serviceRes
	}()

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case err := <-errorChan:
		return utils.FiberResponseErrorJson(c, err.MessageDesc, err.StatusCode)
	case result := <-resultChan:
		return utils.FiberResponseJson(c, result, result.StatusCode)
	}
}

// @Tags User
// @Description Fetch profile information
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.GetProfileResponseModel "ok"
// @Failure 400 {object} utils.ErrorResponseModel "invalid input"
// @Failure 401 {object} utils.ErrorResponseModel "unauthorized"
// @Failure 500 {object} utils.ErrorResponseModel "internal server error"
// @Router /v1/user/profile [get]
func getProfile(c *fiber.Ctx) error {
	resultChan := make(chan *dto.GetProfileResponseModel)
	errorChan := make(chan utils.ErrorResponseModel)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		context, contextErr := db.Connect()
		defer db.ConnectClose(context)
		if contextErr != nil {
			errorChan <- utils.ErrorResponseModel{MessageDesc: contextErr.Error(), StatusCode: 500}
			return
		}

		data := auth.DecodeToken(c)

		resModel := &dto.GetProfileResponseModel{}
		service := &UserService{context}
		serviceRes := service.GetProfile(data["uuid"].(string), resModel)
		resultChan <- serviceRes
	}()

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case err := <-errorChan:
		return utils.FiberResponseErrorJson(c, err.MessageDesc, err.StatusCode)
	case result := <-resultChan:
		return utils.FiberResponseJson(c, result, result.StatusCode)
	}
}

func updateUser(c *fiber.Ctx) error {
	data := auth.DecodeToken(c)
	return c.JSON(fiber.Map{
		"uuid":        data["uuid"],
		"messageDesc": "User updated",
	})
}

func deleteUser(c *fiber.Ctx) error {
	data := auth.DecodeToken(c)
	return c.JSON(fiber.Map{
		"uuid":        data["uuid"],
		"messageDesc": "User updated",
	})
}
