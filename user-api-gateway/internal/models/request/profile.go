package request

type UpdateProfile struct {
	FirstName string `json:"first_name" binding:"required" validate:"min=2,max=20,alpha"`
	LastName  string `json:"last_name" binding:"required" validate:"min=2,max=20,alpha"`
	Email     string `json:"email" binding:"required" validate:"email"`
	Address   string `json:"address" binding:"required" validate:"min=5,max=100"`
	Pincode   string   `json:"pincode" validate:"required,min=6,max=6,pincode"`
}
/*
{
	"first_name": "John",
	"last_name": "Doe",
	"email": "aaa@mymail.com",
	"address": "123, Main Street",
	"pincode": "123456"
}
*/

type VerifyOTPForPwChange struct {
	Otp string `json:"otp" binding:"required" `
}

type SettingNewPassword struct {
	NewPassword     string `json:"new_password" binding:"required" validate:"required,min=8,max=20"`
	ConfirmPassword string `json:"confirm_password" binding:"required" validate:"required,eqfield=NewPassword"`
}
/*
{
	"new_password": "password",
	"confirm_password": "password"
}
*/

type UpdatePasswordUsingOldPw struct {
	OldPassword string `json:"old_password" binding:"required" validate:"required,min=8,max=20"`
	NewPassword string `json:"new_password" binding:"required" validate:"required,min=8,max=20"`
	ConfirmPassword string `json:"confirm_password" binding:"required" validate:"required,eqfield=NewPassword"`
}
/*
{
	"old_password": "password",
	"new_password": "password",
	"confirm_password": "password"
}
*/