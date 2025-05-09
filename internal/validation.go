package internal

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var phoneRegex = regexp.MustCompile(`^09[0-9]{9}$`) // Adjust to your needs

func PhoneValidation(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	if phone != "" {
		return phoneRegex.MatchString(phone)
	}
	return true
}
