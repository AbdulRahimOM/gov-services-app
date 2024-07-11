package models

type KsebSectionCode struct {
	ID              int32  `json:"id" gorm:"primaryKey" `
	SectionOfficeID int32  `json:"sec_office_id" gorm:"not null"`
	SectionCode     string `json:"sec_code" gorm:"not null"`
}

type UserKsebConsumerNumber struct {
	ID             int32  `json:"id" gorm:"primaryKey" `
	UserID         int32  `json:"user_id" gorm:"not null"`
	ConsumerNumber string `json:"consumer_number" gorm:"not null"`
	NickName       string `json:"nick_name" gorm:"default:''"`
}
