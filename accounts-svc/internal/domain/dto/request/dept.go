package request

type AdminAddDept struct {
	AdminID 	   int32  `json:"admin_id" binding:"required" validate:"min=1,number"`
	NewDept NewDept `json:"new_dept" binding:"required"`
}

type NewDept struct {
	Name        string `json:"dept_name" binding:"required" validate:"min=2,max=50,alpha"`
	Description string `json:"dept_description" binding:"required" validate:"min=2,max=100,alpha"`
}

type OfficeSearchCriteria struct {
	Name             string `json:"name"`
	Address          string `json:"address"`
	Id               int32  `json:"id"`
	DeptID           int32  `json:"dept_id"`
	Rank             int32  `json:"rank"`
	SuperiorOfficeID int32  `json:"superior_office_id"`
}