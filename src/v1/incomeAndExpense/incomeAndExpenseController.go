package incomeAndExpense

import (
	auth "bjm/auth/jwt"
	db "bjm/db/benjamit"
	"bjm/db/benjamit/models"
	"bjm/src/v1/incomeAndExpense/dto"
	"bjm/utils"

	"github.com/gofiber/fiber/v2"
)

// @Tags IncomeAndExpense
// @Description Create a income and expense
// @Description TransactionType="DEBIT" | "CREDIT"
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body dto.CreateListRequestModel true "create list request"
// @Success 201 {object} dto.CreateListResponseModel "created"
// @Failure 400 {object} utils.ErrorResponseModel "invalid input"
// @Failure 500 {object} utils.ErrorResponseModel "internal server error"
// @Router /v1/incomeAndExpense/user/create [post]
func createList(c *fiber.Ctx) error {
	reqModel := &dto.CreateListRequestModel{}
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

	getUuid := auth.DecodeToken(c)["uuid"].(string)
	resModel := &dto.CreateListResponseModel{}
	service := &IncomeAndExpenseService{context}

	serviceRes := service.CreateList(reqModel, resModel, models.MANUAL, getUuid)

	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
}
