package validators

import (
	"bjm/db/benjamit/models"

	"github.com/go-playground/validator/v10"
)

func MessageTypeEnum(validate *validator.Validate) {
	validate.RegisterValidation("messageTypeEnum", func(fl validator.FieldLevel) bool {
		role := fl.Field().String()
		return role == string(models.TEXT) || role == string(models.EMOJI) || role == string(models.IMAGE)
	})
}
