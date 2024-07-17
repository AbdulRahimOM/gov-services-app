package ksebrepo

import (
	"time"

	"github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/domain/models"
	repointerface "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/domain/repository/interface"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	"gorm.io/gorm"
)

type KsebRepository struct {
	DB *gorm.DB
}

func NewKsebRepository(db *gorm.DB) repointerface.IKsebRepo {
	return &KsebRepository{DB: db}
}

func (kr KsebRepository) GetUserIdByComplaintId(complaintID int32) (int32, error) {
	var userId int32
	result := kr.DB.Raw("SELECT user_id FROM kseb_complaints WHERE id=?", complaintID).Scan(&userId)
	if result.Error != nil {
		return 0, result.Error
	}
	return userId, nil
}

func (kr KsebRepository) RegisterSectionCode(req *requests.KsebRegSectionCode) (int32, error) {
	entry := models.KsebSectionCode{
		SectionOfficeID: req.OfficeId,
		SectionCode:     req.SectionCode,
	}
	result := kr.DB.Create(&entry)
	if result.Error != nil {
		return 0, result.Error
	}
	return entry.ID, nil
}

func (kr KsebRepository) CheckIfSectionCodeExists(sectionCode string) (bool, error) {
	var count int64
	result := kr.DB.Raw("SELECT COUNT(*) FROM kseb_section_codes WHERE section_code=?", sectionCode).Scan(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (kr KsebRepository) IsSectionCodeRegistered(sectionCode string) (bool, error) {
	var count int64
	result := kr.DB.Raw("SELECT COUNT(*) FROM kseb_section_codes WHERE section_code=?", sectionCode).Scan(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (kr KsebRepository) AddConsumerNumber(userID int32, consumerNumber, nickName string) error {
	entry := models.UserKsebConsumerNumber{
		UserID:         userID,
		ConsumerNumber: consumerNumber,
		NickName:       nickName,
	}
	result := kr.DB.Create(&entry)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (kr KsebRepository) GetUserConsumerNumbers(userID int32) (*[]commondto.UserConsumerNumber, error) {
	var consumerNumbers []commondto.UserConsumerNumber
	result := kr.DB.Table("user_kseb_consumer_numbers").Select("id, consumer_number, nick_name").Where("user_id=?", userID).Scan(&consumerNumbers)
	if result.Error != nil {
		return nil, result.Error
	}
	return &consumerNumbers, nil
}

func (kr KsebRepository) CheckIfConsumerNumberRegisteredByUser(userID int32, consumerNumber string) (bool, error) {
	var count int64
	result := kr.DB.Raw("SELECT COUNT(*) FROM user_kseb_consumer_numbers WHERE user_id=? AND consumer_number=?", userID, consumerNumber).Scan(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (kr KsebRepository) RaiseComplaint(userID int32, ksebComplaint *models.KsebComplaint) (int32, error) {

	result := kr.DB.Create(&ksebComplaint)
	if result.Error != nil {
		return 0, result.Error
	}
	return ksebComplaint.ID, nil
}

func (kr KsebRepository) AdminGetAllComplaints() (*[]models.KsebComplaint, error) {
	var complaints []models.KsebComplaint
	result := kr.DB.Find(&complaints)
	if result.Error != nil {
		return nil, result.Error
	}
	return &complaints, nil
}

func (kr KsebRepository) AdminGetAllComplaintsByStatus(adminID int32, status string) (*[]models.KsebComplaint, error) {
	var complaints []models.KsebComplaint
	result := kr.DB.Where("status=?", status).Find(&complaints)
	if result.Error != nil {
		return nil, result.Error
	}
	return &complaints, nil
}

func (kr KsebRepository) AdminGetAllComplaintsAttendedByHimOrNotOpened(adminID int32) (*[]models.KsebComplaint, error) {
	var complaints []models.KsebComplaint
	result := kr.DB.Where("attendeder_id=? OR status='not-opened'", adminID).Find(&complaints)
	if result.Error != nil {
		return nil, result.Error
	}
	return &complaints, nil
}

func (kr KsebRepository) AdminGetAllComplaintsAttendedByHimByStatus(adminID int32, status string) (*[]models.KsebComplaint, error) {
	var complaints []models.KsebComplaint
	result := kr.DB.Where("attendeder_id=? AND status=?", adminID, status).Find(&complaints)
	if result.Error != nil {
		return nil, result.Error
	}
	return &complaints, nil
}

func (kr KsebRepository) GetComplaintByID(complaintID int32) (*models.KsebComplaint, error) {
	var complaint models.KsebComplaint
	result := kr.DB.Where("id=?", complaintID).First(&complaint)
	if result.Error != nil {
		return nil, result.Error
	}
	return &complaint, nil
}

func (kr KsebRepository) MarkComplaintAsOpened(complaintID, adminID int32) error {
	result := kr.DB.Model(&models.KsebComplaint{}).Where("id=?", complaintID).Update("status", "opened").Update("attendeder_id", adminID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (kr KsebRepository) MarkComplaintAsClosed(complaintID int32, remarks string) error {
	result := kr.DB.Model(&models.KsebComplaint{}).Where("id=?", complaintID).
	Update("status", "closed").
	Update("remarks", remarks).
	Update("closed_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	return nil
}