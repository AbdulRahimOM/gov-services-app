package request

type AdminSearchCriteria struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	DeptID      int32
	RankID      int32
	Designation string
}

type AdminAddSubAdmin struct {
	AdminID int32
	NewSubAdmin NewSubAdmin
}

type NewSubAdmin struct {
	FirstName   string `json:"first_name" binding:"required" validate:"gte=2,lte=20"`
	LastName    string `json:"last_name" binding:"required" validate:"gte=2,lte=20"`
	Email       string `json:"email" binding:"required" validate:"email"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,e164,len=13"`
	DeptID      int32  `json:"dept_id" binding:"required" validate:"min=1,number"`
	Designation string `json:"designation" binding:"required" validate:"gte=2,lte=50"`
	RankID      int32  `json:"rank_id" binding:"required" validate:"min=1,number"`
	OfficeID    int32  `json:"office_id" binding:"required" validate:"min=1,number"`
}