package incomeAndExpense

import (
	auth "bjm/auth/jwt"
	db "bjm/db/benjamit"
	"bjm/src/http/v1/incomeAndExpense/dto"
	"bjm/utils"
	"bjm/utils/enums"

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
// @Failure 401 {object} utils.ErrorResponseModel "unauthorized"
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

	serviceRes := service.CreateList(reqModel, resModel, enums.MANUAL, getUuid)

	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
}

// @Tags IncomeAndExpense
// @Description Get all income and expense
// @Description unlimited: take=-1 or skip=-1
// @Description sort: "ASC" | "DESC"
// @Description Didn't search in between date is not request body StartDate and EndDate
// @Description Didn't sort in column is not request body sort and sortColumn
// @Description sortColumn: "income_and_expenses.amount" | "product_sellings.sell_price" | "product_sellings.cost_price" | "unit_types.name" | "income_and_expenses.quantity" | "income_and_expenses.description" | "products.name"
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body dto.GetAllListRequestModel true "get all list request"
// @Success 200 {object} dto.GetAllListResponseModel "ok"
// @Failure 400 {object} utils.ErrorResponseModel "invalid input"
// @Failure 401 {object} utils.ErrorResponseModel "unauthorized"
// @Failure 500 {object} utils.ErrorResponseModel "internal server error"
// @Router /v1/incomeAndExpense/user/getAll [post]
func getAllList(c *fiber.Ctx) error {
	reqModel := &dto.GetAllListRequestModel{}
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

	resModel := &dto.GetAllListResponseModel{}
	service := &IncomeAndExpenseService{context}

	serviceRes := service.GetAllList(reqModel, resModel)

	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
}
