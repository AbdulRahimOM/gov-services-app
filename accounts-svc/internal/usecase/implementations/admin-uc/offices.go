package adminuc

import (
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
)

// AdminGetOffices
func (a *AdminUseCase) AdminGetOffices() (*[]models.Office, string, error) {
	offices, err := a.adminRepo.AdminGetOffices()
	if err != nil {
		return nil, respcode.DBError, err
	} else {
		return offices, "", nil
	}
}
