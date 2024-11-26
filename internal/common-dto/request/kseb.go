package requests

type KsebRegSectionCode struct {
	SectionCode string `json:"section_code" binding:"required" validate:"len=4,numeric"`
	OfficeId    int32  `json:"office_id" binding:"required"`
}

type UserAddConsumerNumber struct {
	ConsumerNumber string `json:"consumer_number" binding:"required" validate:"len=13,numeric"`
	NickName       string `json:"nick_name" `
}

type KSEBComplaint struct {
	Type           string `json:"type" binding:"required" validate:"oneof=standard custom"`
	Category       string `json:"category" binding:"required" validate:"oneof=kseb"`
	Title          string `json:"title" binding:"required" validate:"gte=5,lte=50"`
	Description    string `json:"description" binding:"required" validate:"gte=5,lte=500"`
	ConsumerNumber string `json:"consumer_number" binding:"required" validate:"len=13,numeric"`
}

type SendMessage struct {
	Message string `json:"message" binding:"required" validate:"gte=1,lte=500"`
}

type KsebCloseComplaint struct {
	ComplaintId int32  `json:"complaint_id" binding:"required"`
	Remarks     string `json:"remarks" binding:"required" validate:"gte=1,lte=500"`
}
