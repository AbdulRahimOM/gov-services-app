package adminrepo

import (
	"fmt"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
	repointerface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/interface"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"

	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) repointerface.IAdminRepo {
	return &AdminRepository{DB: db}
}

func (ur AdminRepository) GetDesignationByAdminID(adminID int32) (string, error) {
	var designation string
	result := ur.DB.Raw("SELECT designation FROM admins WHERE id=?", adminID).Scan(&designation)
	if result.Error != nil {
		return "", result.Error
	}
	return designation, nil
}

func (ur AdminRepository) CheckIfOfficeNameExists(name *string) (bool, error) {
	var count int64
	result := ur.DB.Raw("SELECT COUNT(*) FROM offices WHERE name=?", *name).Scan(&count)
	if result.Error != nil {
		return false, fmt.Errorf("@db: failed to check if office name exists: %v", result.Error)
	}
	return count > 0, nil
}

func (ur AdminRepository) CheckIfAdminUsernameExists(username *string) (bool, error) {
	var count int64
	result := ur.DB.Raw("SELECT COUNT(*) FROM admins WHERE username=?", *username).Scan(&count)
	if result.Error != nil {
		return false, fmt.Errorf("@db: failed to check if admin username exists: %v", result.Error)
	}
	return count > 0, nil
}

func (ur AdminRepository) AddSubAdmin(newSubAdmin *models.Admin) (int32, error) {

	result := ur.DB.Create(newSubAdmin)
	if result.Error != nil {
		return 0, result.Error
	}
	return newSubAdmin.ID, nil
}

func (ur AdminRepository) AdminGetAdmins(sc *request.AdminSearchCriteria) (*[]commondto.Admin, error) {
	var admins []commondto.Admin
	firstName := "%"
	if sc.FirstName != "" {
		firstName = "%" + sc.FirstName + "%"
	}
	lastName := "%"
	if sc.LastName != "" {
		lastName = "%" + sc.LastName + "%"
	}
	email := "%"
	if sc.Email != "" {
		email = "%" + sc.Email + "%"
	}
	phoneNumber := "%"
	if sc.PhoneNumber != "" {
		phoneNumber = "%" + sc.PhoneNumber + "%"
	}
	designation := "%"
	if sc.Designation != "" {
		designation = "%" + sc.Designation + "%"
	}
	result := ur.DB.Raw(`
					SELECT 
						id, f_name, l_name, email, address, pincode, phone_number, office_id, designation
					FROM 
						admins
					WHERE
						(f_name ILIKE ?) AND
						(l_name ILIKE ?) AND
						(email ILIKE ?) AND
						(phone_number ILIKE ?) AND
						(designation ILIKE ?) AND
						(office_id = ? OR ? = 0)
					`,
		firstName,
		lastName,
		email,
		phoneNumber,
		designation,
		sc.OfficeId,sc.OfficeId,
	).Scan(&admins)
	if result.Error != nil {
		return nil, fmt.Errorf("@db: failed to get admins: %v", result.Error)
	}
	return &admins, nil
}
