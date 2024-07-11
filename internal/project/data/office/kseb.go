package office

// var
// (
// 	KSEB_Region_Office = "Region Office"
// 	KSEB_Circle_Office = "Circle Office"
// 	KSEB_Division_Office = "Division Office"
// 	KSEB_SubDivision_Office = "SubDivision Office"
// 	KSEB_Section_Office = "Section Office"
// )

// model
//Rank 4
type KSEB_Region_Office struct {
	ID   int32  `json:"id"`
	Name string `json:"name" gorm:"column:name"`
}

//Rank 5
type KSEB_Circle_Office struct {
	ID       int32  `json:"id"`
	Name     string `json:"name" gorm:"column:name"`
	RegionID int32  `json:"regionId" gorm:"column:region_id"`
}

//Rank 6
type KSEB_Division_Office struct {
	ID       int32  `json:"id"`
	Name     string `json:"name" gorm:"column:name"`
	CircleID int32  `json:"circleId" gorm:"column:circle_id"`
}

//Rank 7
type KSEB_SubDivision_Office struct {
	ID         int32  `json:"id"`
	Name       string `json:"name" gorm:"column:name"`
	DivisionID int32  `json:"divisionId" gorm:"column:division_id"`
}

//Rank 8
type KSEB_Section_Office struct {
	ID            int32  `json:"id"`
	Name          string `json:"name" gorm:"column:name"`
	SubDivisionID int32  `json:"subDivisionId" gorm:"column:sub_division_id"`
	Code          string `json:"code" gorm:"column:code"`
}
