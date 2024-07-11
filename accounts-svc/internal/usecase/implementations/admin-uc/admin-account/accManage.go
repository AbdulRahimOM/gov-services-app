package adminaccount

import (
	"fmt"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
)

// AdminGetAdmins
func (a *AdminUseCase) AdminGetAdmins(adminID int32, searchCriteria *request.AdminSearchCriteria) (*[]commondto.Admin, string, error) {
	admins, err := a.adminRepo.AdminGetAdmins(searchCriteria)
	if err != nil {
		return nil, respcode.DBError, fmt.Errorf("@db: failed to get admins: %v", err)
	}
	return admins, "", nil
}
