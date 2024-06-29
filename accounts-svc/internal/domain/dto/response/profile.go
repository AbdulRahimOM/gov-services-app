package response

type GetOTPForPwChange struct {
	Last4Digits string `json:"last4Digits"`
}

type VerifyOTPForPwChange struct {
	TempToken string
}
