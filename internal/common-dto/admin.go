package commondto

type Admin struct {
	ID          int32  `json:"id"`
	FirstName   string `json:"firstName" gorm:"column:f_name"`
	LastName    string `json:"lastName" gorm:"column:l_name"`
	Email       string `json:"email" gorm:"column:email"`
	Address     string `json:"address" gorm:"column:address"`
	Pincode     string `json:"pincode" gorm:"column:pincode"`
	PhoneNumber string `json:"phoneNumber" gorm:"column:phone_number"`
	DeptID      int32  `json:"deptId" gorm:"column:dept_id"`
	Designation string `json:"designation" gorm:"column:designation"`
	RankID      int32  `json:"rankId" gorm:"column:rank_id"`
}