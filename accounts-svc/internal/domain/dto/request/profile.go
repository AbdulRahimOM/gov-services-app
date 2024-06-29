package request

type UpdateProfile struct {
	UserId   int32
	FirstName string
	LastName  string
	Email     string
	Address   string
	Pincode   string
}

type UpdatePasswordUsingOldPw struct {
	UserId      int32
	OldPassword string
	NewPassword string
}