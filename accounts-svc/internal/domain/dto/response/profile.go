package response

type UserGetOTPForPwChange struct {
	Last4Digits string `json:"last4Digits"`
}

type UserVerifyOTPForPwChange struct {
	TempToken string
}
