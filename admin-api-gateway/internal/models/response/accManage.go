package response

import commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"

type AdminGetAdminsResponse struct {
	Status string             `json:"status"`
	Admins []*commondto.Admin `json:"admins"`
}

type AdminGetOffices struct {
	Status  string              `json:"status"`
	Offices []*commondto.Office `json:"offices"`
}
