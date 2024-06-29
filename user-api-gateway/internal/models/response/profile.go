package response

type GetOTPResponse struct {
	Status      string `json:"status"`
	Last4Digits string `json:"last4Digits"`
}

type GetProfileResponse struct {
	Status  string  `json:"status"`
	Profile Profile `json:"profile"`
}

type Profile struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Pincode     string  `json:"pincode"`
	PhoneNumber string `json:"phoneNumber"`
}
