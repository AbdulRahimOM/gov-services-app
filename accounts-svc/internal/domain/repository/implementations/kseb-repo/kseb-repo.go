package ksebrepo

import (
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
	repointerface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/interface"
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