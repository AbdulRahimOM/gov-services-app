package repo

import (
	"errors"
	"fmt"

	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	dberror "github.com/AbdulRahimOM/gov-services-app/shared/std-response/error/db"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type IUserRepo interface {
	CheckIfMobileIsRegistered(mobile *string) (bool, error)
	GetUserByMobile(mobile *string) (*dto.LoggedInUser, error)
	GetUserWithPasswordByMobile(mobile *string) (*dto.LoggedInUser, *string, error)
	GetPasswordByUserID(userID int32) (*string, error)
	UpdatePasswordByUserID(userID int32, hashedPassword *string) error
	CreateSigningUpUser(mobile *string, isBlocked bool) (int32, error)

	GetMobileByUserID(userID int32) (string, error)
	UpdatePassword(userID int32, hashedPassword *string) error

	UserGetProfileByUserID(userID int32) (*dto.UserProfile, error)
	UserUpdateProfile(*request.UserUpdateProfile) error
}

func NewUserRepository(db *gorm.DB) IUserRepo {
	return &UserRepository{DB: db}
}

// GetPasswordByUserID
func (ur UserRepository) GetPasswordByUserID(userID int32) (*string, error) {
	var hashedPw string
	result := ur.DB.Raw("SELECT hashed_pw FROM users WHERE id=?", userID).Scan(&hashedPw)
	if result.Error != nil {
		return nil, fmt.Errorf("@db: failed to get password: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, dberror.ErrRecordNotFound
	}
	return &hashedPw, nil
}

// UpdatePasswordByUserID
func (ur UserRepository) UpdatePasswordByUserID(userID int32, hashedPassword *string) error {
	result := ur.DB.Exec("UPDATE users SET hashed_pw=? WHERE id=?", *hashedPassword, userID)
	if result.Error != nil {
		return fmt.Errorf("@db: failed to update password: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return dberror.ErrRecordNotFound
	}
	return nil
}

// UserGetProfileByUserID
func (ur UserRepository) UserGetProfileByUserID(userID int32) (*dto.UserProfile, error) {
	var profile dto.UserProfile
	result := ur.DB.Raw("SELECT f_name, l_name, email, address, pincode FROM users WHERE id=?", userID).Scan(&profile)
	if result.Error != nil {
		return nil, fmt.Errorf("@db: failed to get profile: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, dberror.ErrRecordNotFound
	}
	return &profile, nil
}

func (ur UserRepository) UserUpdateProfile(req *request.UserUpdateProfile) error {
	result := ur.DB.Exec("UPDATE users SET f_name=?, l_name=?, email=?, address=?, pincode=? WHERE id=?", req.FirstName, req.LastName, req.Email, req.Address, req.Pincode, req.UserId)
	if result.Error != nil {
		return fmt.Errorf("@db: failed to update profile: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return dberror.ErrRecordNotFound
	}
	return nil
}

func (ur UserRepository) GetMobileByUserID(userID int32) (string, error) {
	var mobile string
	result := ur.DB.Raw("SELECT mobile FROM users WHERE id=?", userID).Scan(&mobile)
	if result.Error != nil {
		return "", fmt.Errorf("@db: failed to get mobile: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return "", dberror.ErrRecordNotFound
	}
	return mobile, nil
}

func (ur UserRepository) CheckIfMobileIsRegistered(mobile *string) (bool, error) {
	var exists bool
	err := ur.DB.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE mobile = ?)", *mobile).Scan(&exists).Error
	if err != nil {
		return false, fmt.Errorf("@db: failed to check if mobile is registered: %v", err)
	}
	return exists, nil
}

func (ur UserRepository) GetUserByMobile(mobile *string) (*dto.LoggedInUser, error) {
	var dtoUser dto.LoggedInUser
	result := ur.DB.Raw("SELECT id, f_name, l_name FROM users WHERE mobile=?", *mobile).Scan(&dtoUser)
	if result.Error != nil {
		return nil, fmt.Errorf("@db: failed to get user: %v", result.Error)

	}
	if result.RowsAffected == 0 {
		return nil, dberror.ErrRecordNotFound
	}
	return &dtoUser, nil
}

func (ur UserRepository) GetUserWithPasswordByMobile(mobile *string) (*dto.LoggedInUser, *string, error) {
	var dtoUser dto.LoggedInUser
	var hashedPw string
	row := ur.DB.Raw("SELECT id, f_name, l_name, hashed_pw FROM users WHERE mobile=?", *mobile).Row()
	err := row.Scan(&dtoUser.ID, &dtoUser.FName, &dtoUser.LName, &hashedPw)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, dberror.ErrRecordNotFound
		}
		return nil, nil, fmt.Errorf("@db: failed to get user: %v", err)
	}
	return &dtoUser, &hashedPw, nil
}

func (ur UserRepository) CreateSigningUpUser(mobile *string, isBlocked bool) (int32, error) {
	var id int32
	result := ur.DB.Raw("INSERT INTO users (mobile) VALUES (?) RETURNING id", *mobile).Scan(&id)
	if result.Error != nil {
		return 0, fmt.Errorf("@db: failed to create signing up user: %v", result.Error)
	}
	return id, nil
}

func (ur UserRepository) UpdatePassword(userID int32, hashedPassword *string) error {
	result := ur.DB.Exec("UPDATE users SET hashed_pw=? WHERE id=?", *hashedPassword, userID)
	if result.Error != nil {
		return fmt.Errorf("@db: failed to update password: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return dberror.ErrRecordNotFound
	}
	return nil
}
