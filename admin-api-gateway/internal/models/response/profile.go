package response

type GetOTPResponse struct {
	Status      string `json:"status"`
	Last4Digits string `json:"last4Digits"`
}

type AdminGetProfileResponse struct {
	Status  string  `json:"status"`
	Profile Profile `json:"profile"`
}

type Profile struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Username	string `json:"username"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Pincode     string `json:"pincode"`
	PhoneNumber string `json:"phoneNumber"`
	DeptID      int32  `json:"deptId"`
	Designation string `json:"designation"`
	RankID      int32  `json:"rankId"`
}
