package ucinterface

import (
	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/response"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
)

type IAdminUC interface {
	//login
	VerifyPasswordForLogin(username, password *string) (*response.AdminLogin, string, error)

	//profile
	AdminGetProfile(adminID int32) (*dto.AdminProfile, string, error)
	AdminUpdateProfile(req *request.AdminUpdateProfile) (string, error)
	AdminUpdatePasswordUsingOldPw(req *request.AdminUpdatePasswordUsingOldPw) (string, error)

	AdminGetOffices(*request.OfficeSearchCriteria) (*[]models.Office, string, error)

	//manage accounts
	AdminGetAdmins(adminID int32, searchCriteria *request.AdminSearchCriteria) (*[]commondto.Admin, string, error)
}

type IAppointmentUC interface {
	// AppointOfficer(req *request.AppointOfficer) (int32, string, error)
	AppointAttender(adminID int32, appointee *requests.Appointee) (int32, string, error)
	CreateChildOffice(adminID int32, proposedChildOffice *requests.ProposedOffice) (int32, string, error)
	AppointChildOfficeHead(adminID int32, childOfficeID int32, appointee *requests.Appointee) (int32, string, error)
	AppointChildOfficeDeputyHead(adminID int32, childOfficeID int32, appointee *requests.Appointee) (int32, string, error)
}

type IKsebAdminUC interface {
	RegisterSectionCode(adminId int32, req *requests.KsebRegSectionCode) (savedRecordId int32, responseCode string, err error)
}