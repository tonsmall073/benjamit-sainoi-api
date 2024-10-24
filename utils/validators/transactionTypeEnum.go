package validators

import (
	"bjm/utils/enums"

	"github.com/go-playground/validator/v10"
)

func TransactionTypeEnum(validate *validator.Validate) {
	validate.RegisterValidation("transactionTypeEnum", func(fl validator.FieldLevel) bool {
		role := fl.Field().String()
		return role == string(enums.CREDIT) || role == string(enums.DEBIT)
	})
}
