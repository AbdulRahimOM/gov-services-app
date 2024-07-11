package requests

type KsebRegSectionCode struct {
	SectionCode string `json:"section_code" binding:"required" validate:"len=4,numeric"`
	OfficeId int32 `json:"office_id" binding:"required"`
}

type UserAddConsumerNumber struct {
	ConsumerNumber string `json:"consumer_number" binding:"required" validate:"len=11,numeric"`
}