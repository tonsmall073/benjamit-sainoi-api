package validators

import (
	"bjm/utils/enums"

	"github.com/go-playground/validator/v10"
)

func SortEnum(validate *validator.Validate) {
	validate.RegisterValidation("sortEnum", func(fl validator.FieldLevel) bool {
		role := fl.Field().String()
		return role == string(enums.ASC) || role == string(enums.DESC)
	})
}
