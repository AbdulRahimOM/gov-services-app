package response

import (
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
)

type UserLogin struct {
	User  *dto.LoggedInUser
	Token *string
}
type UserSignUp struct {
	UserId int32
	Token  *string
}

type NewUserSetPw struct {
	Token *string
}

type UserGetOTPForPwChange struct {
	Last4Digits string `json:"last4Digits"`
}

type UserVerifyOTPForPwChange struct {
	TempToken string
}
