package models

type Admin struct {
	ID          int32  `gorm:"primaryKey"`
	FName       string `gorm:"f_name;default:''"`
	LName       string `gorm:"l_name;default:''"`
	Username    string `gorm:"username;default:''"`
	HashedPW    string `gorm:"hashed_pw;default:''"`
	Email       string `gorm:"email;default:''"`
	Address     string `gorm:"address;default:''"`
	Pincode     string `gorm:"pincode;default:''"`
	Mobile      string `gorm:"mobile;default:''"`
	DeptID      int32  `gorm:"dept_id;default:0"`
	Resignation string `gorm:"role_id;default:0"`
	RankID      int32  `gorm:"rank_id;default:0"`

	IsActive bool `gorm:"is_active;default:true"`
}
