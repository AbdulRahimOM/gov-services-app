package repointerface

import (
	"github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/domain/models"
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

	//complaint
	RaiseComplaint(userID int32, ksebComplaint *models.KsebComplaint) (int32, error)
	GetUserIdByComplaintId(complaintID int32) (userId int32, err error)
	AdminGetAllComplaints() (*[]models.KsebComplaint, error)
	AdminGetAllComplaintsByStatus(adminID int32, status string) (*[]models.KsebComplaint, error)
	AdminGetAllComplaintsAttendedByHimOrNotOpened(adminID int32) (*[]models.KsebComplaint, error)
	AdminGetAllComplaintsAttendedByHimByStatus(adminID int32, status string) (*[]models.KsebComplaint, error)

	GetComplaintByID(complaintID int32) (*models.KsebComplaint, error)
	MarkComplaintAsOpened(complaintID, adminID int32) error
	MarkComplaintAsClosed(complaintID int32, remarks string) error
}
