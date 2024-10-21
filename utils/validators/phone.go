package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func Phone(validate *validator.Validate) {
	validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		phoneRegex := regexp.MustCompile(`^\+?[0-9]{10,15}$`) // ตัวอย่าง regex สำหรับหมายเลขโทรศัพท์
		return phoneRegex.MatchString(fl.Field().String())
	})
}
