package commondto

type Admin struct {
	ID          int32  `json:"id"`
	FirstName   string `json:"firstName" gorm:"column:f_name"`
	LastName    string `json:"lastName" gorm:"column:l_name"`
	Email       string `json:"email" gorm:"column:email"`
	Address     string `json:"address" gorm:"column:address"`
	Pincode     string `json:"pincode" gorm:"column:pincode"`
	PhoneNumber string `json:"phoneNumber" gorm:"column:phone_number"`
	OfficeId   int32  `json:"officeId" gorm:"column:office_id"`
	// PostID      int32  `json:"postId" gorm:"column:post_id"`
	Designation string `json:"designation" gorm:"column:designation"`
}

type Office struct {
	ID               int32  `json:"id"`
	Name             string `json:"name" gorm:"column:name"`
	DeptID           int32  `json:"deptId" gorm:"column:dept_id"`
	Rank             int32  `json:"rank" gorm:"column:hierarchy_rank"`
	Address          string `json:"address" gorm:"column:address"`
	SuperiorOfficeID int32  `json:"superiorOfficeId" gorm:"column:parent_office_id"`
}
