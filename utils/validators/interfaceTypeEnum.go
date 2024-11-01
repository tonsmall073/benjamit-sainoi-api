package validators

import (
	"bjm/utils/enums"

	"github.com/go-playground/validator/v10"
)

func InterfaceTypeEnum(validate *validator.Validate) {
	validate.RegisterValidation("interfaceTypeEnum", func(fl validator.FieldLevel) bool {
		role := fl.Field().String()
		return role == string(enums.HTTP) || role == string(enums.GRPC)
	})
}
