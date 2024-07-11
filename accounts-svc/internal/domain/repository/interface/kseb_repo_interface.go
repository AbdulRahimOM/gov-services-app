package repointerface

import (
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
)

type IKsebRepo interface {
	RegisterSectionCode(sectionCodeReq *requests.KsebRegSectionCode) (int32, error)
	CheckIfSectionCodeExists(sectionCode string) (bool, error)
	IsSectionCodeRegistered(sectionCode string) (bool, error)
	AddConsumerNumber(userID int32, consumerNumber, nickName string) error
	GetUserConsumerNumbers(userID int32) (*[]commondto.UserConsumerNumber, error)
	CheckIfConsumerNumberRegisteredByUser(userID int32, consumerNumber string) (bool, error)
}
