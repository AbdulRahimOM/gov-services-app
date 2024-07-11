package requests

type KsebRegSectionCode struct {
	SectionCode string `json:"section_code" binding:"required" validate:"len=4,numeric"`
	OfficeId int32 `json:"office_id" binding:"required"`
}