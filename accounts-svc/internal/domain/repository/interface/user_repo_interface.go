package repointerface

import (
	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
)

type IUserRepo interface {
	CheckIfPhoneNumberIsRegistered(phoneNumber *string) (bool, error)
	GetUserByPhoneNumber(phoneNumber *string) (*dto.LoggedInUser, error)
	GetUserWithPasswordByPhoneNumber(phoneNumber *string) (*dto.LoggedInUser, *string, error)
	GetPasswordByUserID(userID int32) (*string, error)
	UpdatePasswordByUserID(userID int32, hashedPassword *string) error
	CreateSigningUpUser(phoneNumber *string, isBlocked bool) (int32, error)

	GetPhoneNumberByUserID(userID int32) (string, error)
	UpdatePassword(userID int32, hashedPassword *string) error

	UserGetProfileByUserID(userID int32) (*dto.UserProfile, error)
	UserUpdateProfile(*request.UserUpdateProfile) error
}
