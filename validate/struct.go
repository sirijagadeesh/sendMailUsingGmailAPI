package validate

import "github.com/go-playground/validator/v10"

// Struct validate given struct based translate error to english based on tagName.
// Incase of tagName empty, field name will be used.
func Struct(data any, tagName string) error {
	validator := validator.New()

	return validator.Struct(data)
}
