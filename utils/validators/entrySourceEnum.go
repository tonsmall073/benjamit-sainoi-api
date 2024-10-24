package validators

import (
	"bjm/utils/enums"

	"github.com/go-playground/validator/v10"
)

func EntrySourceEnum(validate *validator.Validate) {
	validate.RegisterValidation("entrySourceEnum", func(fl validator.FieldLevel) bool {
		role := fl.Field().String()
		return role == string(enums.SYSTEM) || role == string(enums.MANUAL)
	})
}
