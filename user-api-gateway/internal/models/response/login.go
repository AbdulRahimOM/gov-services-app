package response

type UserLogin struct {
	Status string `json:"status"`
	// Msg      string        `json:"message"`
	UserData UserBasicData `json:"user"`
	Token    string        `json:"token"`
}
type UpdateToken struct {
	Status string `json:"status"`
	Token string `json:"token"`
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

// message UserSignUpViaOTPResponse {
//     string message = 1;
//     string token = 2;
//     SignedUpUserDetails userDetails = 3;
// }
// message SignedUpUserDetails {
//     int32 id = 1;
//     string phoneNumber = 4;
// }

type VerifyOTPForPwChangeResponse struct {
	Status string `json:"status"`
	Msg    string `json:"message"`
	TempToken string `json:"tempToken"`
}