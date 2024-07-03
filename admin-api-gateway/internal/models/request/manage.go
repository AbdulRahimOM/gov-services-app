package request

type AdminAddSubAdmin struct {
	FirstName   string `json:"first_name" binding:"required" validate:"gte=2,lte=20"`
	LastName    string `json:"last_name" binding:"required" validate:"gte=2,lte=20"`
	Email       string `json:"email" binding:"required" validate:"email"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,e164,len=13"`
	DeptID      int32  `json:"dept_id" binding:"required" validate:"min=1,number"`
	Designation string `json:"designation" binding:"required" validate:"gte=2,lte=50"`
	RankID      int32  `json:"rank_id" binding:"required" validate:"min=1,number"`
	OfficeID    int32  `json:"office_id" binding:"required" validate:"min=1,number"`
}

type AdminAddDept struct {
	DeptName        string `json:"dept_name" binding:"required" validate:"gte=2,lte=50"`
	DeptDescription string `json:"dept_description" binding:"required" validate:"gte=2,lte=100"`
}

type AdminCreateSubOffice struct {
	OfficeName     string `json:"office_name" binding:"required" validate:"min=2,max=50,alpha"`
	Region         string `json:"region" binding:"required" validate:"min=2,max=50,alpha"`
	OfficeLevel    int32  `json:"office_level" binding:"required" validate:"min=1"`
	ParentOfficeID int32  `json:"parent_office_id" binding:"required" validate:"min=1"`
}
