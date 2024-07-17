package ksebuc

import (
	"fmt"

	repo "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/domain/repository/interface"
	usecase "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/usecase/interface"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
)

type KSEBAgencyAdminUseCase struct {
	ksebRepo  repo.IKsebRepo
}

func NewKSEBAgencyAdminUseCase(ksebRepo repo.IKsebRepo) usecase.IKsebAgencyAdminUC {
	return &KSEBAgencyAdminUseCase{
		ksebRepo:  ksebRepo,
	}
}

// RegisterSectionCode
func (k *KSEBAgencyAdminUseCase) RegisterSectionCode(adminID int32, req *requests.KsebRegSectionCode) (int32, string, error) {
	//check if section code already exists
	exists,err:=k.ksebRepo.CheckIfSectionCodeExists(req.SectionCode)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: failed to check if section code exists: %v", err)
	}
	if exists {
		return 0, respcode.KSEB_SectionCodeExists, fmt.Errorf("section code already registered")
	}

	regId, err := k.ksebRepo.RegisterSectionCode(req)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: failed to register section code: %v", err)
	}
	return regId, "", nil
}
