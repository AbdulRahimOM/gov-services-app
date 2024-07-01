package response

import (
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
)

type AdminLogin struct {
	Admin  *dto.LoggedInAdmin
	Token *string
}
type AdminSignUp struct {
	AdminId int32
	Token  *string
}

type NewAdminSetPw struct {
	Token *string
}

type AdminGetOTPForPwChange struct {
	Last4Digits string `json:"last4Digits"`
}

type AdminVerifyOTPForPwChange struct {
	TempToken string
}
