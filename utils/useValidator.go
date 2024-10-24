package utils

import (
	"bjm/utils/validators"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func UseValidator() {
	Validate = validator.New()

	/*----------------RegisterValidation----------------*/
	validators.Phone(Validate)
	validators.RoleEnum(Validate)
	validators.EntrySourceEnum(Validate)
	validators.MessageTypeEnum(Validate)
	validators.TransactionTypeEnum(Validate)
	validators.SortEnum(Validate)
}
