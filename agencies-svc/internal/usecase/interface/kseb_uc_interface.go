package ucinterface

import (
	"github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/domain/dto/request"
	"github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/domain/models"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
)

type IKsebAgencyAdminUC interface {
	RegisterSectionCode(adminId int32, req *requests.KsebRegSectionCode) (savedRecordId int32, responseCode string, err error)
	GetComplaints(adminId int32, searchCriteria *request.KsebComplaintSearchCriteria) (*[]models.KsebComplaint, string, error)
	OpenComplaint(adminId, complaintId int32) (*models.KsebComplaint,string, error)
	CloseComplaint(adminId, complaintId int32, remarks string) (string, error)
}

type IKsebAgencyUserUC interface {
	AddConsumerNumber(userID int32, consumerNumber, nickName string) (string, error)
	GetUserConsumerNumbers(userID int32) (*[]commondto.UserConsumerNumber, string, error)

	//complaint
	RaiseComplaint(userID int32, complaint *requests.KSEBComplaint) (int32, string, error)

	CheckIfComplaintBelongsToUser(userId, complaintId int32) (bool, string, error)
}
