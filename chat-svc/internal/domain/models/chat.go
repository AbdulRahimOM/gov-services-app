package models

import "time"

type ChatMessage struct {
	MessageId   int32    `json:"message_id" gorm:"primaryKey"`
	ComplaintId int32     `json:"complaint_id" gorm:"not null"`
	SenderId    int32     `json:"author_id" gorm:"not null"`
	SenderType  string    `json:"author_type" gorm:"not null"`
	Content     string    `json:"content" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	HasRead     bool      `json:"has_read" gorm:"default:false"`
}
