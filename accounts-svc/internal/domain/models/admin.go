package models

type Admin struct {
	ID          int32  `gorm:"primaryKey;autoIncrement"`
	FName       string `gorm:"f_name;default:''"`
	LName       string `gorm:"l_name;default:''"`
	Username    string `gorm:"username;default:''"`
	HashedPW    string `gorm:"hashed_pw;default:''"`
	Email       string `gorm:"email;default:''"`
	Address     string `gorm:"address;default:''"`
	Pincode     string `gorm:"pincode;default:''"`
	PhoneNumber string `gorm:"phoneNumber;default:''"`
	DeptID      int32  `gorm:"dept_id;default:0"`
	Designation string `gorm:"role_id;default:''"`
	RankID      int32  `gorm:"rank_id;default:0"`
	OfficeID    int32  `gorm:"office_id;default:0"`

	IsActive bool `gorm:"is_active;default:true"`
}

type Department struct {
	ID          int32  `gorm:"primaryKey"`
	Name        string `gorm:"name;default:''"`
	Description string `gorm:"description;default:''"`
}

type Office struct {
	ID               int32  `gorm:"primaryKey"`
	DeptID           int32  `gorm:"dept_id;default:0"`
	HierarchyRank    int32  `gorm:"hierarchy_rank;default:0"`
	RegionName       string `gorm:"region_name;default:''"`
	HeadOfficerID    int32  `gorm:"head_officer;default:0"`
	OfficeLocation   string `gorm:"office_location;default:''"`
	SuperiorOfficeID int32  `gorm:"parent_office_id;default:0"`
}