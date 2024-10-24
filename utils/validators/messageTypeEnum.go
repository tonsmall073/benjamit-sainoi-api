package validators

import (
	"bjm/utils/enums"

	"github.com/go-playground/validator/v10"
)

func MessageTypeEnum(validate *validator.Validate) {
	validate.RegisterValidation("messageTypeEnum", func(fl validator.FieldLevel) bool {
		role := fl.Field().String()
		return role == string(enums.TEXT) || role == string(enums.EMOJI) || role == string(enums.IMAGE)
	})
}
