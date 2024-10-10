package user

import (
	"bjm/src/v2/user/dto"
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
// @Router /v2/user/create [post]
func createUser(c *fiber.Ctx) error {
	reqModel := &dto.CreateUserRequestModel{}

	err := c.BodyParser(reqModel)
	if err != nil {
		return utils.FiberResponseErrorJson(c, err.Error(), 400)
	}

	resModel := &dto.CreateUserResponseModel{
		Data: &dto.CreateUserDataListResponseModel{
			Nickname:  reqModel.Nickname,
			Firstname: reqModel.Firstname,
			Lastname:  reqModel.Lastname,
		},
		MessageDesc: utils.HttpStatusCodes[201],
		StatusCode:  201,
	}
	return utils.FiberResponseJson(c, resModel, resModel.StatusCode)
}
