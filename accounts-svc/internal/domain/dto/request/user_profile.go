package request

type UserUpdateProfile struct {
	UserId    int32
	FirstName string
	LastName  string
	Email     string
	Address   string
	Pincode   string
}

type UserUpdatePasswordUsingOldPw struct {
	UserId      int32
	OldPassword string
	NewPassword string
}
