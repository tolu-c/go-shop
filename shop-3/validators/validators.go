package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateCoolTitle(f validator.FieldLevel) bool {
	return strings.Contains(f.Field().String(), "Cool")
}
