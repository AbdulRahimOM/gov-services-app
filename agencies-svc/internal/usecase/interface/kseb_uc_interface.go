package ucinterface

import (
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
)

type IKsebAgencyAdminUC interface {
	RegisterSectionCode(adminId int32, req *requests.KsebRegSectionCode) (savedRecordId int32, responseCode string, err error)
}

type IKsebAgencyUserUC interface {
	AddConsumerNumber(userID int32, consumerNumber, nickName string) (string, error)
	GetUserConsumerNumbers(userID int32) (*[]commondto.UserConsumerNumber, string, error)

	//complaint
	RaiseComplaint(userID int32, complaint *requests.KSEBComplaint) (int32, string, error)

	CheckIfComplaintBelongsToUser(userId, complaintId int32) (bool, string, error)
}
