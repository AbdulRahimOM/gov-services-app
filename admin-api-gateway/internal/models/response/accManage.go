package response

import commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"

type AdminGetAdminsResponse struct {
	Status string             `json:"status"`
	Admins []*commondto.Admin `json:"admins"`
}

type AdminAddSubAdmin struct {
	Status       string `json:"status"`
	AddedAdminID int32  `json:"added_admin_id"`
}

type AdminAddDept struct {
	Status      string `json:"status"`
	AddedDeptID int32  `json:"added_dept_id"`
}

type AdminGetDepts struct {
	Status string                  `json:"status"`
	Depts  []*commondto.Department `json:"depts"`
}

type AdminGetOffices struct {
	Status  string              `json:"status"`
	Offices []*commondto.Office `json:"offices"`
}
