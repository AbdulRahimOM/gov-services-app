package repointerface

import (
	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
)

type IAdminRepo interface {
	//self account management
	GetAdminWithPasswordByUsername(username *string) (*dto.LoggedInAdmin, *string, error)
	GetPasswordByAdminID(adminID int32) (*string, error)
	UpdatePasswordByAdminID(adminID int32, hashedPassword *string) error
	AdminGetProfileByAdminID(adminID int32) (*dto.AdminProfile, error)
	AdminUpdateProfile(req *request.AdminUpdateProfile) error

	//account management
	AdminGetAdmins(searchCriteria *request.AdminSearchCriteria) (*[]commondto.Admin, error)
}
