package requests

type Appointee struct {
	FirstName   string `json:"first_name" binding:"required" validate:"gte=2,lte=20"`
	LastName    string `json:"last_name" binding:"required" validate:"gte=2,lte=20"`
	Email       string `json:"email" binding:"required" validate:"email"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,e164,len=13"`
}

type ProposedOffice struct {
	Name             string `json:"name" binding:"required" validate:"gte=2,lte=50"`
	Address          string `json:"address" binding:"required" validate:"gte=2,lte=100"`
}

type AppointChildOfficeHead struct {
	Appointee Appointee `json:"appointee" binding:"required"`
	ChildOfficeID int32 `json:"child_office_id" binding:"required"`
}



type AppointChildOfficeDeputyHead AppointChildOfficeHead