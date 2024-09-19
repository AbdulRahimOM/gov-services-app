package adminrepo

import (
	"fmt"

	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
	dberror "github.com/AbdulRahimOM/gov-services-app/internal/std-response/error/db"
)

func (ur AdminRepository) CheckIfOfficeExists(officeID int32) (bool, error) {
	var count int64
	result := ur.DB.Raw("SELECT COUNT(*) FROM offices WHERE id=?", officeID).Scan(&count)
	if result.Error != nil {
		return false, fmt.Errorf("@db: failed to check if office exists: %v", result.Error)
	}
	return count > 0, nil
}

func (ur AdminRepository) GetRankOfOffice(officeID int32) (int32, error) {
	var rank int32
	fmt.Println("officeID: ", officeID)
	result := ur.DB.Raw("SELECT rank FROM offices WHERE id=?", officeID).Scan(&rank)
	if result.Error != nil {
		fmt.Println("error: ", result.Error)
		return 0, fmt.Errorf("@db: failed to get rank of office: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		fmt.Println("error: ", dberror.ErrRecordNotFound)
		return 0, dberror.ErrRecordNotFound
	}
	return rank, nil
}

func (ur AdminRepository) GetSuperiorOfficeIdByOfficeId(officeID int32) (int32, error) {
	var superiorOfficeID int32
	result := ur.DB.Raw("SELECT superior_office_id FROM offices WHERE id=?", officeID).Scan(&superiorOfficeID)
	if result.Error != nil {
		return 0, fmt.Errorf("@db: failed to get superior office id: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return 0, dberror.ErrRecordNotFound
	}
	return superiorOfficeID, nil
}

func (ur AdminRepository) GetOfficeDetailsByAdminID(adminID int32) (*dto.OfficeDetails, error) {
	var officeDetails dto.OfficeDetails//, rank, dept_id
	var officeId int32
	result := ur.DB.Raw("SELECT office_id FROM admins WHERE id=?", adminID).Scan(&officeId)
	if result.Error != nil {
		return nil, fmt.Errorf("@db: failed to get office details: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, dberror.ErrRecordNotFound
	}

	result = ur.DB.Raw("SELECT rank, dept_id FROM offices WHERE id=?", officeId).Scan(&officeDetails)
	if result.Error != nil {
		return nil, fmt.Errorf("@db: failed to get office details: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, dberror.ErrRecordNotFound
	}
	officeDetails.ID = officeId
	return &officeDetails, nil
}

func (ur AdminRepository) GetOfficeDetailsByOfficeID(officeID int32) (*dto.OfficeDetails, error) {
	var officeDetails dto.OfficeDetails
	result := ur.DB.Raw("SELECT id, rank, dept_id FROM offices WHERE id=?", officeID).Scan(&officeDetails)
	if result.Error != nil {
		return nil, fmt.Errorf("@db: failed to get office details: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, dberror.ErrRecordNotFound
	}
	return &officeDetails, nil
}

func (ur AdminRepository) AddChildOffice(newOffice *models.Office) (int32, error) {
	result := ur.DB.Create(newOffice)
	if result.Error != nil {
		return 0, fmt.Errorf("@db: failed to add child office: %v", result.Error)
	}
	return newOffice.ID, nil
}

func (ur AdminRepository) GetOfficeIDByAdminID(adminID int32) (int32, error) {
	//first get post id
	var officeID int32
	result := ur.DB.Raw("SELECT office_id FROM admins WHERE id=?", adminID).Scan(&officeID)
	if result.Error != nil {
		return 0, fmt.Errorf("@db: failed to get post id: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return 0, dberror.ErrRecordNotFound
	}

	return officeID, nil
}

func (ur AdminRepository) AdminGetOffices(sc *request.OfficeSearchCriteria) (*[]models.Office, error) {
	var offices []models.Office
	query := ur.DB.Model(&models.Office{})
	if sc.Name != "" {
		query = query.Where("name ILIKE ?", "%"+sc.Name+"%")
	}
	if sc.Address != "" {
		query = query.Where("address ILIKE ?", "%"+sc.Address+"%")
	}
	if sc.Id != 0 {
		query = query.Where("id = ?", sc.Id)
	}
	if sc.DeptID != 0 {
		query = query.Where("dept_id = ?", sc.DeptID)
	}
	if sc.Rank != 0 {
		query = query.Where("rank = ?", sc.Rank)
	}
	if sc.SuperiorOfficeID != 0 {
		query = query.Where("superior_office_id = ?", sc.SuperiorOfficeID)
	}

	result := query.Find(&offices)
	if result.Error != nil {
		return nil, fmt.Errorf("@db: failed to get offices: %v", result.Error)
	}
	return &offices, nil
}

func (ur AdminRepository) CheckOccupancyByDesignation(officeID int32, designation string) (bool, error) {
	var count int64
	result := ur.DB.Raw("SELECT COUNT(*) FROM admins WHERE office_id=? AND designation=?", officeID, designation).Scan(&count)
	if result.Error != nil {
		return false, fmt.Errorf("@db: failed to check occupancy by designation: %v", result.Error)
	}
	return count > 0, nil
}
