package response

type AdminLogin struct {
	Status    string         `json:"status"`
	AdminData AdminBasicData `json:"admin"`
	Token     string         `json:"token"`
}

type UpdateToken struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

type AdminBasicData struct {
	Id          int32  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	// DeptID      int32  `json:"deptId"`
	// RankID      int32  `json:"rankId"`
	Designation string `json:"designation"`
}

type AdminSignUp struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Admin   PreliminaryAdminData `json:"admin"`
	Token   string               `json:"token"`
}

type PreliminaryAdminData struct {
	Id          int32  `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
}

type AdminVerifyOTPForPwChangeResponse struct {
	Status    string `json:"status"`
	Msg       string `json:"message"`
	TempToken string `json:"tempToken"`
}
