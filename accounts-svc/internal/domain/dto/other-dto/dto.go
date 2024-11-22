package dto

type LoggedInUser struct {
	ID     int32
	FName  string
	LName  string
}
type LoggedInAdmin struct {
	ID     int32
	FName  string
	LName  string
	PhoneNumber string
	// DeptID int32
	// RankID int32
	Designation string
}

type UserProfile struct {
	FirstName   string `json:"firstName" gorm:"column:f_name"`
	LastName    string `json:"lastName" gorm:"column:l_name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Pincode     string `json:"pincode"`
	PhoneNumber string `json:"phoneNumber" gorm:"column:phone_number"`
}

type AdminProfile struct {
	FirstName   string `json:"firstName" gorm:"column:f_name"`
	LastName    string `json:"lastName" gorm:"column:l_name"`
	Email       string `json:"email" gorm:"column:email"`
	Address     string `json:"address" gorm:"column:address"`
	Pincode     string `json:"pincode" gorm:"column:pincode"`
	PhoneNumber string `json:"phoneNumber" gorm:"column:phone_number"`
}

type OfficeDetails struct {
	ID               int32  `json:"id"`
	Rank			 int32  `json:"rank"`
	DeptID 		 int32  `json:"deptId"`
}
