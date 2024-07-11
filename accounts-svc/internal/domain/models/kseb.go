package models

type KsebSectionCode struct {
	ID              int32  `gorm:"primaryKey" json:"id"`
	SectionOfficeID int32  `json:"sec_office_id"`
	SectionCode     string `json:"sec_code"`
}