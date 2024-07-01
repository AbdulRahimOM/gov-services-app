package models

type User struct {
	ID          int32  `gorm:"primaryKey"`
	FName       string `gorm:"f_name;default:''"`
	LName       string `gorm:"l_name;default:''"`
	HashedPW    string `gorm:"hashed_pw;default:''"`
	Email       string `gorm:"email;default:''"`
	Address     string `gorm:"address;default:''"`
	Pincode     string `gorm:"pincode;default:''"`
	PhoneNumber string `gorm:"phone_number;default:''"`
	IsBlocked   bool   `gorm:"is_blocked;default:false"`
}
