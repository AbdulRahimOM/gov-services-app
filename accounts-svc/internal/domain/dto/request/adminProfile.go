package request

type AdminUpdateProfile struct {
	AdminId     int32
	FirstName   string
	LastName    string
	Email       string
	Address     string
	Pincode     string
	PhoneNumber string
}

type AdminUpdatePasswordUsingOldPw struct {
	AdminId     int32
	OldPassword string
	NewPassword string
}
