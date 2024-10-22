package incomeAndExpense

import (
	auth "bjm/auth/jwt"
	db "bjm/db/benjamit"
	"bjm/db/benjamit/models"
	"bjm/src/v1/incomeAndExpense/dto"
	"bjm/utils"
	"sync"

	"github.com/gofiber/fiber/v2"
)

// @Tags IncomeAndExpense
// @Description Create a income and expense
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body dto.CreateListRequestModel true "create list request"
// @Success 201 {object} dto.CreateListResponseModel "created"
// @Failure 400 {object} utils.ErrorResponseModel "invalid input"
// @Failure 500 {object} utils.ErrorResponseModel "internal server error"
// @Router /v1/incomeAndExpense/user/create [post]
func createList(c *fiber.Ctx) error {
	resultChan := make(chan *dto.CreateListResponseModel)
	errorChan := make(chan utils.ErrorResponseModel)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		reqModel := &dto.CreateListRequestModel{}
		if err := c.BodyParser(reqModel); err != nil {
			errorChan <- utils.ErrorResponseModel{MessageDesc: err.Error(), StatusCode: 400}
			return
		}
		if err := utils.Validate.Struct(reqModel); err != nil {
			errorChan <- utils.ErrorResponseModel{MessageDesc: err.Error(), StatusCode: 400}
			return
		}
		context, contextErr := db.Connect()
		defer db.ConnectClose(context)
		if contextErr != nil {
			errorChan <- utils.ErrorResponseModel{MessageDesc: contextErr.Error(), StatusCode: 500}
			return
		}

		getUuid := auth.DecodeToken(c)["uuid"].(string)
		resModel := &dto.CreateListResponseModel{}
		service := &IncomeAndExpenseService{context}
		serviceRes := service.CreateList(reqModel, resModel, models.MANUAL, getUuid)
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
