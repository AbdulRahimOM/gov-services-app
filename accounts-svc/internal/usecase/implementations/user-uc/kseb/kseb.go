package ksebUserUc

import (
	"fmt"

	repo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/interface"
	usecase "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
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
func (u *KsebUserUseCase) AddConsumerNumber(userID int32, consumerNumber, nickName string) (string, error) {
	// 1. Check if consumerNumber is valid
	// 2. Check if consumerNumber is already registered by the same user

	isConsumerNumberValid, err := u.ksebRepo.IsSectionCodeRegistered(consumerNumber[2:6])
	if err != nil {
		return respcode.DBError, fmt.Errorf("@db error while checking if section code is valid: %v", err)
	}
	if !isConsumerNumberValid {
		return respcode.KSEB_ConsumerNumberInvalid, fmt.Errorf("invalid consumer number. section code part is invalid")
	}

	isConsumerNumberRegistered, err := u.ksebRepo.CheckIfConsumerNumberRegisteredByUser(userID, consumerNumber)
	if err != nil {
		return respcode.DBError, fmt.Errorf("@db error while checking if consumer number is already registered: %v", err)
	}
	if isConsumerNumberRegistered {
		return respcode.KSEB_ConsumerNumberAlreadyRegistered, fmt.Errorf("consumer number is already added by user")
	}

	if nickName == "" {
		nickName = "KSEB_" + consumerNumber[7:13]
	}

	err = u.ksebRepo.AddConsumerNumber(userID, consumerNumber, nickName)
	if err != nil {
		return respcode.DBError, fmt.Errorf("@db error while adding consumer number: %v", err)
	}

	return "", nil
}

// GetUserConsumerNumbers
func (u *KsebUserUseCase) GetUserConsumerNumbers(userID int32) (*[]commondto.UserConsumerNumber, string, error) {
	consumerNumbers, err := u.ksebRepo.GetUserConsumerNumbers(userID)
	if err != nil {
		return nil, respcode.DBError, fmt.Errorf("@db error while getting user consumer numbers: %v", err)
	}

	return consumerNumbers, "", nil
}
