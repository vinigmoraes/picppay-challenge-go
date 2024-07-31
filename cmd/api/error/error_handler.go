package errorhandler

import "github.com/go-playground/validator/v10"

func HandleErrorMessage(field validator.FieldError) string {
	switch field.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return field.Error()
}
