package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationErrorResponse struct {
	Error       bool        `json:"error"`
	FailedField string      `json:"field"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value"`
}

func ValidateRequest(req interface{}) []string {
	errResponse := []string{}
	errs := validate.Struct(req)

	if errs == nil {
		return nil
	}

	for _, err := range errs.(validator.ValidationErrors) {
		e := ValidationErrorResponse{
			// Error:       true,
			FailedField: err.Field(),
			Tag:         err.Tag(),
			Value:       err.Value(),
		}

		message := fmt.Sprintf("[%s]: '%v' | Needs to implement '%s'", e.FailedField, e.Value, e.Tag)

		errResponse = append(errResponse, message)
	}
	return errResponse
}

func ValidateRequestDetailed(req interface{}) []ValidationErrorResponse {

	errResponse := []ValidationErrorResponse{}
	errs := validate.Struct(req)

	if errs == nil {
		return nil
	}

	for _, err := range errs.(validator.ValidationErrors) {
		e := ValidationErrorResponse{
			// Error:       true,
			FailedField: err.Field(),
			Tag:         err.Tag(),
			Value:       err.Value(),
		}

		// message := fmt.Sprintf("[%s]: '%v' | Needs to implement '%s'", e.FailedField, e.Value, e.Tag)

		errResponse = append(errResponse, e)
	}
	return errResponse
}
