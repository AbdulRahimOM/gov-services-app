package ksebrepo

import (
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
	repointerface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/interface"
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

func (kr KsebRepository) AddConsumerNumber(userID int32, consumerNumber string) error {
	entry := models.UserKsebConsumerNumber{
		UserID:         userID,
		ConsumerNumber: consumerNumber,
	}
	result := kr.DB.Create(&entry)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
