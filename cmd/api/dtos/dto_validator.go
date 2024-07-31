package dtos

import (
	"github.com/go-playground/validator/v10"
	errorhandler "picpay-challenge-go/cmd/api/error"
)

func Validate(dto interface{}) []errorhandler.ErrorResponse {
	var validate *validator.Validate
	validate = validator.New(validator.WithRequiredStructEnabled())
	var validationErrors []errorhandler.ErrorResponse

	errors := validate.Struct(dto)

	if errors != nil {
		for _, err := range errors.(validator.ValidationErrors) {
			var response errorhandler.ErrorResponse

			validationErrors = append(validationErrors, handleError(response, err))
		}
		return validationErrors
	}
	return validationErrors
}

func handleError(response errorhandler.ErrorResponse, err validator.FieldError) errorhandler.ErrorResponse {
	switch err.Tag() {
	case "required":
		response.Message = "Field " + err.Field() + " is required"
		response.Error = true
		response.FieldName = err.Field()
		response.Value = err.Value()
	case "email":
		response.Message = "Field " + err.Field() + " is invalid email"
		response.Error = true
		response.FieldName = err.Field()
		response.Value = err.Value()
	}
	return response
}
