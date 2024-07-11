package request

type AdminSearchCriteria struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Designation string
	OfficeId    int32
	// PostID	  int32
}