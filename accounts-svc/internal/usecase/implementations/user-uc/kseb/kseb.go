package ksebUserUc

import (
	"fmt"

	repo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/interface"
	usecase "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
)

type KsebUserUseCase struct {
	userRepo repo.IUserRepo
	ksebRepo repo.IKsebRepo
}

func NewKsebUserUseCase(userRepo repo.IUserRepo, ksebRepo repo.IKsebRepo) usecase.IKsebUserUC {
	return &KsebUserUseCase{
		userRepo: userRepo,
		ksebRepo: ksebRepo,
	}
}

// AddConsumerNumber
func (u *KsebUserUseCase) AddConsumerNumber(userID int32, consumerNumber string) (string, error) {
	// 1. Check if consumerNumber is valid
	// 2. Add consumerNumber to user record

	isConsumerNumberValid, err := u.ksebRepo.IsSectionCodeRegistered(consumerNumber[2:6])
	if err != nil {
		return respcode.DBError, fmt.Errorf("@db error while checking if section code is valid: %v", err)
	}
	if !isConsumerNumberValid {
		return respcode.KSEB_ConsumerNumberInvalid, fmt.Errorf("invalid consumer number. section code part is invalid")
	}

	err = u.ksebRepo.AddConsumerNumber(userID, consumerNumber)
	if err != nil {
		return respcode.DBError, fmt.Errorf("@db error while adding consumer number: %v", err)
	}

	return "", nil
}
