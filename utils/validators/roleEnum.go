package validators

import (
	"bjm/db/benjamit/models"

	"github.com/go-playground/validator/v10"
)

func RoleEnum(validate *validator.Validate) {
	validate.RegisterValidation("roleEnum", func(fl validator.FieldLevel) bool {
		role := fl.Field().String()
		return role == string(models.ADMIN) || role == string(models.USER)
	})
}
