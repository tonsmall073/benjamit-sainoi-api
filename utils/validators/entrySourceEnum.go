package validators

import (
	"bjm/db/benjamit/models"

	"github.com/go-playground/validator/v10"
)

func EntrySourceEnum(validate *validator.Validate) {
	validate.RegisterValidation("entrySourceEnum", func(fl validator.FieldLevel) bool {
		role := fl.Field().String()
		return role == string(models.SYSTEM) || role == string(models.MANUAL)
	})
}
