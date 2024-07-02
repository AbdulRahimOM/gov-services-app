package ucinterface

import (
	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/response"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
)

type IAdminUC interface {
	//login
	VerifyPasswordForLogin(username, password *string) (*response.AdminLogin, string, error)

	//profile
	AdminGetProfile(adminID int32) (*dto.AdminProfile, string, error)
	AdminUpdateProfile(req *request.AdminUpdateProfile) (string, error)
	AdminUpdatePasswordUsingOldPw(req *request.AdminUpdatePasswordUsingOldPw) (string, error)

	//manage accounts
	AdminGetAdmins(adminID int32, searchCriteria *request.AdminSearchCriteria) (*[]commondto.Admin, string, error)
}
