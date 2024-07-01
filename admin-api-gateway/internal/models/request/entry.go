package request

type AdminLoginViaPassword struct {
	Username string `json:"username" binding:"required" validate:"required,min=3,max=50"`
	Password string `json:"password" binding:"required" validate:"required,min=6,max=50"`
}

/*
{
	"phone_number": "+919876543210",
	"password": "password"
}
*/
