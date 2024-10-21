package validators

import (
	"bjm/db/benjamit/models"

	"github.com/go-playground/validator/v10"
)

func TransactionTypeEnum(validate *validator.Validate) {
	validate.RegisterValidation("transactionTypeEnum", func(fl validator.FieldLevel) bool {
		role := fl.Field().String()
		return role == string(models.CREDIT) || role == string(models.DEBIT)
	})
}
