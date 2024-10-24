package user

import (
	auth "bjm/auth/jwt"
	db "bjm/db/benjamit"
	"bjm/src/v1/user/dto"
	"bjm/utils"

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
	reqModel := &dto.CreateUserRequestModel{}

	if err := c.BodyParser(reqModel); err != nil {
		return utils.FiberResponseErrorJson(c, err.Error(), 400)
	}

	if err := utils.Validate.Struct(reqModel); err != nil {
		return utils.FiberResponseErrorJson(c, err.Error(), 400)
	}

	context, contextErr := db.Connect()
	defer db.ConnectClose(context)
	if contextErr != nil {
		return utils.FiberResponseErrorJson(c, contextErr.Error(), 500)
	}

	resModel := &dto.CreateUserResponseModel{}
	service := &UserService{context}

	serviceRes := service.CreateUser(reqModel, resModel)

	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
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
	reqModel := &dto.LoginRequestModel{}

	if err := c.BodyParser(reqModel); err != nil {
		return utils.FiberResponseErrorJson(c, err.Error(), 400)
	}

	if err := utils.Validate.Struct(reqModel); err != nil {
		return utils.FiberResponseErrorJson(c, err.Error(), 400)
	}

	context, contextErr := db.Connect()
	defer db.ConnectClose(context)
	if contextErr != nil {
		return utils.FiberResponseErrorJson(c, contextErr.Error(), 500)
	}

	resModel := &dto.LoginResponseModel{}
	service := &UserService{context}

	serviceRes := service.Login(reqModel, resModel)

	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
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
	context, contextErr := db.Connect()
	defer db.ConnectClose(context)
	if contextErr != nil {
		return utils.FiberResponseErrorJson(c, contextErr.Error(), 500)
	}

	data := auth.DecodeToken(c)

	resModel := &dto.GetProfileResponseModel{}
	service := &UserService{context}

	serviceRes := service.GetProfile(data["uuid"].(string), resModel)

	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
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
