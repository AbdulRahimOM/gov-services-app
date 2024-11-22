package userrepo

import (
	"database/sql"
	"errors"
	"fmt"

	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	repointerface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/interface"
	dberror "github.com/AbdulRahimOM/gov-services-app/internal/std-response/error/db"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repointerface.IUserRepo {
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
	result := ur.DB.Raw("SELECT f_name, l_name, email, address, pincode, phone_number FROM users WHERE id=?", userID).Scan(&profile)
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

func (ur UserRepository) GetPhoneNumberByUserID(userID int32) (string, error) {
	var phoneNumber string
	result := ur.DB.Raw("SELECT phone_number FROM users WHERE id=?", userID).Scan(&phoneNumber)
	if result.Error != nil {
		return "", fmt.Errorf("@db: failed to get phone number: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return "", dberror.ErrRecordNotFound
	}
	return phoneNumber, nil
}

func (ur UserRepository) CheckIfPhoneNumberIsRegistered(phoneNumber *string) (bool, error) {
	var exists bool
	err := ur.DB.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE phone_number = ?)", *phoneNumber).Scan(&exists).Error
	if err != nil {
		return false, fmt.Errorf("@db: failed to check if phone number is registered: %v", err)
	}
	return exists, nil
}

func (ur UserRepository) GetUserByPhoneNumber(phoneNumber *string) (*dto.LoggedInUser, error) {
	var dtoUser dto.LoggedInUser
	result := ur.DB.Raw("SELECT id, f_name, l_name FROM users WHERE phone_number=?", *phoneNumber).Scan(&dtoUser)
	if result.Error != nil {
		return nil, fmt.Errorf("@db: failed to get user: %v", result.Error)

	}
	if result.RowsAffected == 0 {
		return nil, dberror.ErrRecordNotFound
	}
	return &dtoUser, nil
}

func (ur UserRepository) GetUserWithPasswordByPhoneNumber(phoneNumber *string) (*dto.LoggedInUser, *string, error) {
	var dtoUser dto.LoggedInUser
	var hashedPw string
	row := ur.DB.Raw("SELECT id, f_name, l_name, hashed_pw FROM users WHERE phone_number=?", *phoneNumber).Row()
	err := row.Scan(&dtoUser.ID, &dtoUser.FName, &dtoUser.LName, &hashedPw)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, dberror.ErrRecordNotFound
		}
		return nil, nil, fmt.Errorf("@db: failed to get user: %v", err)
	}
	return &dtoUser, &hashedPw, nil
}

func (ur UserRepository) CreateSigningUpUser(phoneNumber *string, isBlocked bool) (int32, error) {
	var id int32
	result := ur.DB.Raw("INSERT INTO users (phone_number) VALUES (?) RETURNING id", *phoneNumber).Scan(&id)
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
