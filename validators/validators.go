package validators

import (
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

func ValidateCoolTitle(fl validator.FieldLevel) bool {
	return strings.Contains(fl.Field().String(), "Cool")
}
