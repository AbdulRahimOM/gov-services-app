package request

type AdminLoginViaPassword struct {
	Username string `json:"username" binding:"required" validate:"min=3,max=50"`
	Password string `json:"password" binding:"required" validate:"min=6,max=50"`
}
