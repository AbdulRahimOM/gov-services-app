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
}

type UserProfile struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Pincode     string `json:"pincode"`
	PhoneNumber string `json:"phoneNumber"`
}

type AdminProfile struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Pincode     string `json:"pincode"`
	PhoneNumber string `json:"phoneNumber"`
}