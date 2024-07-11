package models

import "time"

type Admin struct {
	ID          int32     `gorm:"primaryKey;autoIncrement"`
	FName       string    `gorm:"f_name;default:''"`
	LName       string    `gorm:"l_name;default:''"`
	Username    string    `gorm:"username;default:''"`
	HashedPW    string    `gorm:"hashed_pw;default:''"`
	Email       string    `gorm:"email;default:''"`
	Address     string    `gorm:"address;default:''"`
	Pincode     string    `gorm:"pincode;default:''"`
	PhoneNumber string    `gorm:"phoneNumber;default:''"`
	OfficeID    int32     `gorm:"office_id;default:0"`
	Designation string    `gorm:"designation;default:''"`
	CreatedBy   int32     `gorm:"created_by;default:0"`
	CreatedAt   time.Time `gorm:"created_at;default:CURRENT_TIMESTAMP"`
	IsActive    bool      `gorm:"is_active;default:true"`
}

type Office struct {
	ID               int32  `gorm:"primaryKey"`
	Name             string `gorm:"name;default:''"`
	DeptID           int32  `gorm:"dept_id;default:0"`
	Rank             int32  `gorm:"rank;default:0"`
	Address          string `gorm:"address;default:''"`
	SuperiorOfficeID int32  `gorm:"parent_office_id;default:0"`
}
