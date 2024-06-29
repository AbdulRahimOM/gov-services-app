package request

type UserLoginGetOTP struct {
	PhoneNumber string `json:"phone_number" binding:"required" validate:"required,e164,len=13"`
}

type UserLoginVerifyOTP struct {
	PhoneNumber string `json:"phone_number"`
	OTP         string `json:"otp"`
}

type GetOTPForSignup struct {
	PhoneNumber string `json:"phone_number"`
}

type UserSignpViaOTP struct {
	PhoneNumber string `json:"phone_number"`
	OTP         string `json:"otp"`
}

type UserLoginViaPassword struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

/*
{
	"phone_number": "+919876543210",
	"password": "password"
}
*/
