package adminuc

import (
	"fmt"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
)

// AdminGetDepts
func (a *AdminUseCase) AdminGetDepts() (*[]models.Department, string, error) {
	depts, err := a.adminRepo.GetDepts()
	if err != nil {
		return nil, respcode.DBError, fmt.Errorf("failed to get departments: %v", err)
	}

	return depts, "", nil
}

// AdminAddDept
func (a *AdminUseCase) AdminAddDept(req request.AdminAddDept) (int32, string, error) {
	//check if the admin is a super admin or state head
	rankId, err := a.adminRepo.GetRankByAdminID(req.AdminID)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("failed to get rank: %v", err)
	}
	if rankId != 1 && rankId != 2 {
		return 0, respcode.Unauthorized, fmt.Errorf("this admin(by rank) is not authorized to add a department")
	}

	//check if the department already exists
	nameExists, err := a.adminRepo.CheckIfDeptNameExists(&req.NewDept.Name)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("failed to get department name: %v", err)
	}
	if nameExists {
		return 0, respcode.AlreadyExists, fmt.Errorf("department name already exists")
	}

	//add the department
	newDeptID, err := a.adminRepo.AddDept(req.NewDept)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("failed to add department: %v", err)
	}

	return newDeptID, "", nil
}
