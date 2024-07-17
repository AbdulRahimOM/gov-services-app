package ksebuc

import (
	"fmt"

	"github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/domain/dto/request"
	"github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/domain/models"
	repo "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/domain/repository/interface"
	usecase "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/usecase/interface"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
)

type KSEBAgencyAdminUseCase struct {
	ksebRepo repo.IKsebRepo
}

func NewKSEBAgencyAdminUseCase(ksebRepo repo.IKsebRepo) usecase.IKsebAgencyAdminUC {
	return &KSEBAgencyAdminUseCase{
		ksebRepo: ksebRepo,
	}
}

// GetComplaints
func (k *KSEBAgencyAdminUseCase) GetComplaints(adminID int32, searchCriteria *request.KsebComplaintSearchCriteria) (*[]models.KsebComplaint, string, error) {
	// status := c.DefaultQuery("status", "all")       //all, open, closed, not-opened
	// attenderScope := c.DefaultQuery("scope", "all") //all, me-only
	switch searchCriteria.AttenderScope {
	case "all":
		switch searchCriteria.Status {
		case "all":
			complaints, err := k.ksebRepo.AdminGetAllComplaints()
			if err != nil {
				return nil, respcode.DBError, fmt.Errorf("@db: failed to get complaints: %v", err)
			}
			return complaints, "", nil
		default:
			complaints, err := k.ksebRepo.AdminGetAllComplaintsByStatus(adminID, searchCriteria.Status)
			if err != nil {
				return nil, respcode.DBError, fmt.Errorf("@db: failed to get open complaints: %v", err)
			}
			return complaints, "", nil
		}
	case "me-only":
		switch searchCriteria.Status {
		case "all":
			complaints, err := k.ksebRepo.AdminGetAllComplaintsAttendedByHimOrNotOpened(adminID)
			if err != nil {
				return nil, respcode.DBError, fmt.Errorf("@db: failed to get complaints: %v", err)
			}
			return complaints, "", nil
		case "not-opened":
			complaints, err := k.ksebRepo.AdminGetAllComplaintsByStatus(adminID, searchCriteria.Status)
			if err != nil {
				return nil, respcode.DBError, fmt.Errorf("@db: failed to get open complaints: %v", err)
			}
			return complaints, "", nil
		default:
			complaints, err := k.ksebRepo.AdminGetAllComplaintsAttendedByHimByStatus(adminID, searchCriteria.Status)
			if err != nil {
				return nil, respcode.DBError, fmt.Errorf("@db: failed to get open complaints: %v", err)
			}
			return complaints, "", nil
		}
	default:
		return nil, respcode.InvalidUrlParams, fmt.Errorf("invalid attender scope")
	}

}

// RegisterSectionCode
func (k *KSEBAgencyAdminUseCase) RegisterSectionCode(adminID int32, req *requests.KsebRegSectionCode) (int32, string, error) {
	//check if section code already exists
	exists, err := k.ksebRepo.CheckIfSectionCodeExists(req.SectionCode)
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

// OpenComplaint
func (k *KSEBAgencyAdminUseCase) OpenComplaint(adminID, complaintID int32) (*models.KsebComplaint, string, error) {
	complaint, err := k.ksebRepo.GetComplaintByID(complaintID)
	if err != nil {
		return nil, respcode.DBError, fmt.Errorf("@db: failed to get complaint: %v", err)
	}

	if complaint.Status != "not-opened" {
		return nil, respcode.KSEB_ComplaintAlreadyOpened, fmt.Errorf("complaint is already opened")
	}

	err = k.ksebRepo.MarkComplaintAsOpened(complaintID, adminID)
	if err != nil {
		return nil, respcode.DBError, fmt.Errorf("@db: failed to update complaint-status to 'opened'. Error: %v", err)
	}

	return complaint, "", nil
}

// CloseComplaint
func (k *KSEBAgencyAdminUseCase) CloseComplaint(adminID, complaintID int32, remarks string) (string, error) {
	complaint, err := k.ksebRepo.GetComplaintByID(complaintID)
	if err != nil {
		return respcode.DBError, fmt.Errorf("@db: failed to get complaint: %v", err)
	}

	if complaint.Status != "opened" {
		return respcode.KSEB_ComplaintNotOpened, fmt.Errorf("complaint is not opened")
	}

	err = k.ksebRepo.MarkComplaintAsClosed(complaintID, remarks)
	if err != nil {
		return respcode.DBError, fmt.Errorf("@db: failed to update complaint-status to 'closed'. Error: %v", err)
	}

	return "", nil
}

// CheckIfComplaintBeAccessibleToAdmin
func (k *KSEBAgencyAdminUseCase) CheckIfComplaintBeAccessibleToAdmin(adminID, complaintID int32) (bool, string, error) {
	attenderId, err := k.ksebRepo.GetAttenderIDOfComplaint(complaintID)
	if err != nil {
		return false, respcode.DBError, fmt.Errorf("@db: failed to get complaint: %v", err)
	}

	if attenderId != adminID {
		return false, respcode.KSEB_ComplaintNotAccessibleToAdmin, fmt.Errorf("complaint is not accessible to you")
	}

	return true, "", nil
}
