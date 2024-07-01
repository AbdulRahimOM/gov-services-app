package repointerface

import (
	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
)

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
