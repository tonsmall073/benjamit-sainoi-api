package validators

import (
	"bjm/utils/enums"

	"github.com/go-playground/validator/v10"
)

func RoleEnum(validate *validator.Validate) {
	validate.RegisterValidation("roleEnum", func(fl validator.FieldLevel) bool {
		role := fl.Field().String()
		return role == string(enums.ADMIN) || role == string(enums.USER)
	})
}
