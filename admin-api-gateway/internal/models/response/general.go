package response

import "github.com/AbdulRahimOM/gov-services-app/shared/validation"

type SEE struct {
	Status       string `json:"status"`
	ResponseCode string `json:"error_code"`
	Error        string `json:"error"`
}
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

type SMsE struct {
	Status string   `json:"status"`
	Msg    string   `json:"message"`
	Error  []string `json:"error"`
}

type SM struct {
	Status string `json:"status"`
	Msg    string `json:"message"`
}

type SMT struct {
	Status string `json:"status"`
	Msg    string `json:"message"`
	Token  string `json:"token"`
}
