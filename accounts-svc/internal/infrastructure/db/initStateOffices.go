package db

import (
	"log"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
	"github.com/AbdulRahimOM/gov-services-app/internal/project/data"
	"github.com/AbdulRahimOM/gov-services-app/internal/project/data/dept"
	"gorm.io/gorm/clause"
)

func SeedDataToDbIfNotInitialised() {
	createStateOfficesIfNotExists()
	createStateLevelDeptOfficesIfNotExists()

	initSuperAdminIfNotInitialized()
}

func createStateOfficesIfNotExists() {

	superAdminOffice := models.Office{
		ID:      1,
		Name:    "Super Admin Office",
		DeptID:  0,
		Rank:    1,
		Address: data.StateNodalAddress,
	}

	result := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).Create(&superAdminOffice)

	if result.Error != nil {
		log.Fatal("couldn't create super admin office")
	}

	stateNodalOffice := models.Office{
		ID:               2,
		Name:             "State Nodal Office",
		DeptID:           0,
		Rank:             2,
		Address:          data.StateNodalAddress,
		SuperiorOfficeID: 1,
	}

	createOfficeIfNotExist(stateNodalOffice)
}

func createStateLevelDeptOfficesIfNotExists() {

	policeDeptStateHeadQuarters := models.Office{
		ID:               3,
		Name:             "Police Department State Head Quarters",
		DeptID:           dept.PoliceDeptID,
		Rank:             3,
		Address:          data.PoliceDeptAddress,
		SuperiorOfficeID: 2,
	}

	createOfficeIfNotExist(policeDeptStateHeadQuarters)

	fireAndSafetyDeptStateHeadQuarters := models.Office{
		ID:               4,
		Name:             "Fire and Safety Department State Head Quarters",
		DeptID:           dept.FireAndSafetyDeptID,
		Rank:             3,
		Address:          data.FireAndSafetyDeptAddress,
		SuperiorOfficeID: 2,
	}

	createOfficeIfNotExist(fireAndSafetyDeptStateHeadQuarters)

	healthDeptStateHeadQuarters := models.Office{
		ID:               5,
		Name:             "Health Department State Head Quarters",
		DeptID:           dept.HealthDeptID,
		Rank:             3,
		Address:          data.StateNodalAddress,
		SuperiorOfficeID: 2,
	}

	createOfficeIfNotExist(healthDeptStateHeadQuarters)

	emergencyDeptStateHeadQuarters := models.Office{
		ID:               6,
		Name:             "Emergency Department State Nodal Office",
		DeptID:           dept.EmergencyDeptID,
		Rank:             3,
		Address:          data.StateNodalAddress,
		SuperiorOfficeID: 2,
	}

	createOfficeIfNotExist(emergencyDeptStateHeadQuarters)

	KSEBDeptStateHeadQuarters := models.Office{
		ID:               7,
		Name:             "KSEB Department State Head Quarters",
		DeptID:           dept.KSEBDeptID,
		Rank:             3,
		Address:          data.StateNodalAddress,
		SuperiorOfficeID: 2,
	}

	createOfficeIfNotExist(KSEBDeptStateHeadQuarters)

	KWADeptStateHeadQuarters := models.Office{
		ID:               8,
		Name:             "KWA Department State Head Quarters",
		DeptID:           dept.KWADeptID,
		Rank:             3,
		Address:          data.StateNodalAddress,
		SuperiorOfficeID: 2,
	}

	createOfficeIfNotExist(KWADeptStateHeadQuarters)
}
