package db

import (
	"log"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
	"gorm.io/gorm/clause"
)

func createOfficeIfNotExist(office models.Office) {
	result := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).Create(&office)

	if result.Error != nil {
		log.Fatal("couldn't create office for dept(id): ", office.DeptID)
	}
}

