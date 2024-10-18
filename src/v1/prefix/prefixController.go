package prefix

import (
	db "bjm/db/benjamit"
	"bjm/src/v1/prefix/dto"
	"bjm/utils"
	"sync"

	"github.com/gofiber/fiber/v2"
)

// @Tags Prefix
// @Description Fetch all prefixes
// @Produce json
// @Success 200 {object} dto.GetAllPrefixResponseModel "ok"
// @Failure 400 {object} utils.ErrorResponseModel "invalid input"
// @Failure 500 {object} utils.ErrorResponseModel "internal server error"
// @Router /v1/prefix/getAll [get]
func getAllPrefix(c *fiber.Ctx) error {
	resultChan := make(chan *dto.GetAllPrefixResponseModel)
	errorChan := make(chan utils.ErrorResponseModel)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		context, contextErr := db.Connect()
		defer db.ConnectClose(context)
		if contextErr != nil {
			errorChan <- utils.ErrorResponseModel{
				MessageDesc: contextErr.Error(),
				StatusCode:  500,
			}
			return
		}

		resModel := &dto.GetAllPrefixResponseModel{}
		service := &PrefixService{context}
		serviceRes := service.GetAllPrefix(resModel)
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
