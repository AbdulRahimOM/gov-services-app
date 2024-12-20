package ksebuc

import (
	"fmt"

	repo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/interface"
	usecase "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	"github.com/AbdulRahimOM/gov-services-app/internal/project/data"
	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
)

const (
	ksebSubDivOfficeRank  = 7
	ksebSectionOfficeRank = 8
)

type KsebAdminUseCase struct {
	adminRepo repo.IAdminRepo
}

func NewKsebAdminUseCase(adminRepo repo.IAdminRepo) usecase.IKsebAdminUC {
	return &KsebAdminUseCase{
		adminRepo: adminRepo,
	}
}

// RegisterSectionCode
func (k *KsebAdminUseCase) CheckIfAdminCanRegisterSectionCode(adminID int32, req *requests.KsebRegSectionCode) (int32, string, error) {
	//1. check if office exists with the given section office id
	//2. check if it is section office (by rank(8) and dept id)
	//get office details
	sectionOffice, err := k.adminRepo.GetOfficeDetailsByOfficeID(req.OfficeId)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: failed to get office details: %v", err)
	}
	if sectionOffice.Rank != ksebSectionOfficeRank {
		return 0, respcode.KSEB_SectionOfficeNotValid, fmt.Errorf("invalid section office id")
	}

	if sectionOffice.DeptID != data.DeptID_KSEB {
		return 0, respcode.KSEB_SectionOfficeNotValid, fmt.Errorf("office id doesn't belongs to KSEB department")
	}

	//check if admin belongs to sub division office
	adminOfficeID, err := k.adminRepo.GetOfficeIDByAdminID(adminID)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: failed to get admin office id: %v", err)
	}
	// get superior office id
	subDivisionOfficeID, err := k.adminRepo.GetSuperiorOfficeIdByOfficeId(sectionOffice.ID)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: failed to get superior office of section office: %v", err)
	}
	if adminOfficeID != subDivisionOfficeID {
		return 0, respcode.Unauthorized, fmt.Errorf("admin not authorized to register section code. Admin should belong to sub division office")
	}

	// get office rank
	officeRank, err := k.adminRepo.GetRankOfOffice(adminOfficeID)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: failed to get office rank: %v", err)
	}
	if officeRank != ksebSubDivOfficeRank {
		return 0, respcode.Unauthorized, fmt.Errorf("admin not authorized to register section code. Admin should belong to sub division office(rank 7)")
	}

	//get office's dept id
	officeDeptID, err := k.adminRepo.GetDeptIDByOfficeID(adminOfficeID)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: failed to get office dept id: %v", err)
	}
	if officeDeptID != data.DeptID_KSEB {
		return 0, respcode.Unauthorized, fmt.Errorf("admin not authorized to register section code. Admin should belong to KSEB department")
	}

	//check if desination is head or deputy head
	desig, err := k.adminRepo.GetDesignationByAdminID(adminID)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: failed to get admin designation: %v", err)
	}
	if desig != data.Designation_OfficeHead && desig != data.Designation_DeputyOfficeHead {
		return 0, respcode.Unauthorized, fmt.Errorf("admin not authorized to register section code. Admin should be head or deputy head")
	}

	return 0, "", nil
}
