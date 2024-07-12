package ksebrepo

import (
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
)

func (kr KsebRepository) RaiseComplaint(userID int32, ksebComplaint *models.KsebComplaint) (int32, error) {

	result := kr.DB.Create(&ksebComplaint)
	if result.Error != nil {
		return 0, result.Error
	}
	return ksebComplaint.ID, nil
}
