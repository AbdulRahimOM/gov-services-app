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

type Department struct {
	ID          int32  `json:"id"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

type Office struct {
	ID               int32  `json:"id"`
	DeptID           int32  `json:"deptId" gorm:"column:dept_id"`
	HierarchyRank    int32  `json:"hierarchyRank" gorm:"column:hierarchy_rank"`
	RegionName       string `json:"regionName" gorm:"column:region_name"`
	HeadOfficerID    int32  `json:"headOfficerId" gorm:"column:head_officer"`
	OfficeLocation   string `json:"officeLocation" gorm:"column:office_location"`
	SuperiorOfficeID int32  `json:"superiorOfficeId" gorm:"column:parent_office_id"`
}