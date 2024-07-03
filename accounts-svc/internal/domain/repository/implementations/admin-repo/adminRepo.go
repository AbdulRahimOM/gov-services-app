package adminrepo

import (
	"errors"
	"fmt"

	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
	repointerface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/interface"
	dberror "github.com/AbdulRahimOM/gov-services-app/internal/std-response/error/db"

	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) repointerface.IAdminRepo {
	return &AdminRepository{DB: db}
}

//CheckIfAdminUsernameExists
func (ur AdminRepository) CheckIfAdminUsernameExists(username *string) (bool, error) {
	var count int64
	result := ur.DB.Raw("SELECT COUNT(*) FROM admins WHERE username=?", *username).Scan(&count)
	if result.Error != nil {
		return false, fmt.Errorf("@db: failed to check if admin username exists: %v", result.Error)
	}
	return count > 0, nil
}

//AddSubAdmin
func (ur AdminRepository) AddSubAdmin(newSubAdmin *models.Admin) (int32, error) {

	result := ur.DB.Create(newSubAdmin)
	if result.Error != nil {
		return 0, fmt.Errorf("@db: failed to add sub admin: %v", result.Error)
	}
	return newSubAdmin.ID, nil
}

// AdminGetOffices
func (ur AdminRepository) AdminGetOffices() (*[]models.Office, error) {
	var offices []models.Office
	result := ur.DB.Find(&offices)
	if result.Error != nil {
		return nil, fmt.Errorf("@db: failed to get offices: %v", result.Error)
	}
	return &offices, nil
}

// GetDepts
func (ur AdminRepository) GetDepts() (*[]models.Department, error) {
	var depts []models.Department
	result := ur.DB.Find(&depts)
	if result.Error != nil {
		return nil, fmt.Errorf("@db: failed to get depts: %v", result.Error)
	}
	return &depts, nil
}

// CheckIfDeptNameExists
func (ur AdminRepository) CheckIfDeptNameExists(deptName *string) (bool, error) {
	var count int64
	result := ur.DB.Raw("SELECT COUNT(*) FROM departments WHERE name=?", *deptName).Scan(&count)
	if result.Error != nil {
		return false, fmt.Errorf("@db: failed to check if dept name exists: %v", result.Error)
	}
	return count > 0, nil
}

// GetPasswordByAdminID
func (ur AdminRepository) GetPasswordByAdminID(adminID int32) (*string, error) {
	var hashedPw string
	result := ur.DB.Raw("SELECT hashed_pw FROM admins WHERE id=?", adminID).Scan(&hashedPw)
	if result.Error != nil {
		return nil, fmt.Errorf("@db: failed to get password: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, dberror.ErrRecordNotFound
	}
	return &hashedPw, nil
}

// UpdatePasswordByAdminID
func (ur AdminRepository) UpdatePasswordByAdminID(adminID int32, hashedPassword *string) error {
	result := ur.DB.Exec("UPDATE admins SET hashed_pw=? WHERE id=?", *hashedPassword, adminID)
	if result.Error != nil {
		return fmt.Errorf("@db: failed to update password: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return dberror.ErrRecordNotFound
	}
	return nil
}

// AdminGetProfileByAdminID
func (ur AdminRepository) AdminGetProfileByAdminID(adminID int32) (*dto.AdminProfile, error) {
	var profile dto.AdminProfile
	result := ur.DB.Raw("SELECT f_name, l_name, email, address, pincode, phone_number FROM admins WHERE id=?", adminID).Scan(&profile)
	if result.Error != nil {
		return nil, fmt.Errorf("@db: failed to get profile: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, dberror.ErrRecordNotFound
	}
	return &profile, nil
}

func (ur AdminRepository) GetAdminWithPasswordByUsername(username *string) (*dto.LoggedInAdmin, *string, error) {
	var dtoAdmin dto.LoggedInAdmin
	var hashedPw string
	row := ur.DB.Raw("SELECT id, f_name, l_name, hashed_pw FROM admins WHERE username=?", *username).Row()
	err := row.Scan(&dtoAdmin.ID, &dtoAdmin.FName, &dtoAdmin.LName, &hashedPw)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, dberror.ErrRecordNotFound
		}
		return nil, nil, fmt.Errorf("@db: failed to get admin: %v", err)
	}
	return &dtoAdmin, &hashedPw, nil
}

func (ur AdminRepository) AdminUpdateProfile(req *request.AdminUpdateProfile) error {
	result := ur.DB.Exec("UPDATE admins SET f_name=?, l_name=?, email=?, address=?, pincode=?, phone_number=? WHERE id=?",
		req.FirstName, req.LastName, req.Email, req.Address, req.Pincode, req.PhoneNumber, req.AdminId)
	if result.Error != nil {
		return fmt.Errorf("@db: failed to update profile: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return dberror.ErrRecordNotFound
	}
	return nil
}
