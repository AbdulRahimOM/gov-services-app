package models

type Post struct {
	ID                     int32  `gorm:"primaryKey"`
	PostName               string `gorm:"post_name;default:''"`
	OfficeID               int32  `gorm:"office_id;default:0"`
	MaxNumberOfPosts       int32
	CanAppointSubordinates bool
	Rank                   int32 `gorm:"rank"`
}

type Office struct {
	ID         int32  `gorm:"primaryKey"`
	DeptID     int32  `gorm:"dept_id;default:0"`
	Rank       int32  `gorm:"hierarchy_rank;default:0"`
	RegionName string `gorm:"region_name;default:''"`
	// HeadOfficerID    int32  `gorm:"head_officer;default:0"`
	Address          string `gorm:"address;default:''"`
	SuperiorOfficeID int32  `gorm:"parent_office_id;default:0"`
}
