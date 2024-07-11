package response

import "github.com/AbdulRahimOM/gov-services-app/internal/validation"

type SRE struct {
	Status string `json:"status"`
	ResponseCode string `json:"error_code"`
	Error        string `json:"error"`
}

type SMValidationErrors struct {
	Status string `json:"status"`
	ResponseCode string                               `json:"error_code"`
	Errors       []validation.ValidationErrorResponse `json:"error"`
}

type SM struct {
	Status string `json:"status"`
	Msg    string `json:"message"`
}

