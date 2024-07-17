package ksebUserUc

import (
	"fmt"
	"time"

	"github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/domain/models"
	repo "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/domain/repository/interface"
	usecase "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/usecase/interface"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
)

type KsebUserUseCase struct {
	ksebRepo repo.IKsebRepo
}

func NewKsebAgencyUserUseCase(ksebRepo repo.IKsebRepo) usecase.IKsebAgencyUserUC {
	return &KsebUserUseCase{
		ksebRepo: ksebRepo,
	}
}

//CheckIfComplaintBelongsToUser
func (u *KsebUserUseCase) CheckIfComplaintBelongsToUser(userId, complaintId int32) (bool, string, error) {
	userIdOfComplaint, err := u.ksebRepo.GetUserIdByComplaintId(complaintId)
	if err != nil {
		return false, respcode.DBError, fmt.Errorf("@db error while checking if complaint belongs to user: %v", err)
	}
	if userIdOfComplaint != userId {
		return false, respcode.KSEB_ComplaintNotBelongsToUser, fmt.Errorf("complaint does not belong to user")
	}

	return true, "", nil
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

// RaiseComplaint
func (u *KsebUserUseCase) RaiseComplaint(userID int32, complaint *requests.KSEBComplaint) (int32, string, error) {
	entry := models.KsebComplaint{
		UserID:         userID,
		Type:           complaint.Type,
		Title:          complaint.Title,
		Description:    complaint.Description,
		ConsumerNumber: complaint.ConsumerNumber,
		Status:         "not opened",
		CreatedAt:      time.Now(),
	}
	complaintID, err := u.ksebRepo.RaiseComplaint(userID, &entry)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db error while raising complaint: %v", err)
	}

	return complaintID, "", nil
}
