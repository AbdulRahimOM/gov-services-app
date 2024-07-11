package repointerface

import (
	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
)

type IAdminRepo interface {
	//self account management
	GetAdminWithPasswordByUsername(username *string) (*dto.LoggedInAdmin, *string, error)
	GetPasswordByAdminID(adminID int32) (*string, error)
	UpdatePasswordByAdminID(adminID int32, hashedPassword *string) error
	AdminGetProfileByAdminID(adminID int32) (*dto.AdminProfile, error)
	AdminUpdateProfile(req *request.AdminUpdateProfile) error

	//office management
	AdminGetOffices(searchCriteria *request.OfficeSearchCriteria) (*[]models.Office, error)
	GetOfficeIDByAdminID(adminID int32) (officeID int32, err error)
	GetSuperiorOfficeIdByOfficeId(officeID int32)(superiorOfficeID int32,err error)
	GetOfficeDetailsByAdminID(adminID int32) (officeDetails *dto.OfficeDetails, err error)
	AddChildOffice(newOffice *models.Office) (int32, error)
	GetRankOfOffice(officeID int32) (int32, error)
	CheckOccupancyByDesignation(officeID int32,designation string) (bool, error)
	CheckIfOfficeNameExists(name *string) (bool, error)

	//account management
	AdminGetAdmins(searchCriteria *request.AdminSearchCriteria) (*[]commondto.Admin, error)
	AddSubAdmin(newSubAdmin *models.Admin) (int32, error)
	CheckIfAdminUsernameExists(username *string) (bool, error)
	
}
