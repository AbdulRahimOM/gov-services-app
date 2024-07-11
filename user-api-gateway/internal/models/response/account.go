package response

type UserLogin struct {
	Status string `json:"status"`
	UserData UserBasicData `json:"user"`
	Token    string        `json:"token"`
}
type UpdateToken struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}
type UserBasicData struct {
	Id          int32  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
}

type UserSignUp struct {
	Status  string              `json:"status"`
	Message string              `json:"message"`
	User    PreliminaryUserData `json:"user"`
	Token   string              `json:"token"`
}

type PreliminaryUserData struct {
	Id          int32  `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
}

type UserVerifyOTPForPwChangeResponse struct {
	Status    string `json:"status"`
	Msg       string `json:"message"`
	TempToken string `json:"tempToken"`
}

type GetOTPResponse struct {
	Status      string `json:"status"`
	Last4Digits string `json:"last4Digits"`
}

type UserGetProfileResponse struct {
	Status  string  `json:"status"`
	Profile Profile `json:"profile"`
}

type Profile struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Pincode     string `json:"pincode"`
	PhoneNumber string `json:"phoneNumber"`
}
