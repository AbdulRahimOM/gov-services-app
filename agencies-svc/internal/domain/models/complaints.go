package models

import "time"

type KsebComplaint struct {
	ID     int32 `json:"id" gorm:"primaryKey" `
	UserID int32 `json:"user_id" gorm:"not null"`
	// Category	   string `json:"category" gorm:"not null"`
	Type           string    `json:"type" gorm:"not null"`
	Title          string    `json:"title" gorm:"not null"`
	Description    string    `json:"description" gorm:"not null"`
	ConsumerNumber string    `json:"consumer_number" gorm:"default:''"`
	AttenderID     int32     `json:"attender_id" gorm:"default:0"`
	Status         string    `json:"status" gorm:"default:'not-opened'"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	Remarks        string    `json:"remarks" gorm:"default:''"`
	ClosedAt       time.Time `json:"closed_at" gorm:"default:null"`
}
