package request

type UserLoginGetOTP struct {
	PhoneNumber string `json:"phone_number" binding:"required" validate:"e164,len=13"`
}

type UserLoginVerifyOTP struct {
	PhoneNumber string `json:"phone_number" binding:"required" validate:"e164,len=13"`
	OTP         string `json:"otp" binding:"required"`
}

type GetOTPForSignup struct {
	PhoneNumber string `json:"phone_number" binding:"required" validate:"e164,len=13"`
}

type UserSignpViaOTP struct {
	PhoneNumber string `json:"phone_number" binding:"required" validate:"e164,len=13"`
	OTP         string `json:"otp" binding:"required"`
}

type UserLoginViaPassword struct {
	PhoneNumber string `json:"phone_number" binding:"required" validate:"e164,len=13"`
	Password    string `json:"password" binding:"required" validate:"min=5,max=50"`
}
