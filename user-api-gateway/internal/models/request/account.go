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

type UserUpdateProfile struct {
	FirstName string `json:"first_name" binding:"required" validate:"min=2,max=20,alpha"`
	LastName  string `json:"last_name" binding:"required" validate:"min=2,max=20,alpha"`
	Email     string `json:"email" binding:"required" validate:"email"`
	Address   string `json:"address" binding:"required" validate:"min=6,max=100"`
	Pincode   string `json:"pincode" validate:"min=6,max=6,pincode"`
}

type UserVerifyOTPForPwChange struct {
	Otp string `json:"otp" binding:"required" `
}

type SettingNewPassword struct {
	NewPassword     string `json:"new_password" binding:"required" validate:"min=5,max=50"`
	ConfirmPassword string `json:"confirm_password" binding:"required" validate:"eqfield=NewPassword"`
}

type UserUpdatePasswordUsingOldPw struct {
	OldPassword     string `json:"old_password" binding:"required" validate:"min=5,max=50"`
	NewPassword     string `json:"new_password" binding:"required" validate:"min=5,max=50,nefield=OldPassword"`
	ConfirmPassword string `json:"confirm_password" binding:"required" validate:"eqfield=NewPassword"`
}
