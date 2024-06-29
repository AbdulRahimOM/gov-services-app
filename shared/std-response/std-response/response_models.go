package stdresponse

import "github.com/AbdulRahimOM/gov-services-app/shared/validation"

type SRE struct {
	Status string `json:"status"`
	// Msg string `json:"message"`
	ResponseCode string `json:"error_code"`
	Error        string `json:"error"`
}
type SMValidationErrors struct {
	Status string `json:"status"`
	// Msg string `json:"message"`
	ResponseCode string                               `json:"error_code"`
	Errors       []validation.ValidationErrorResponse `json:"error"`
}
