package repointerface

import (
	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
)

type IAdminRepo interface {
	GetAdminWithPasswordByUsername(username *string) (*dto.LoggedInAdmin, *string, error)
	GetPasswordByAdminID(adminID int32) (*string, error)
	UpdatePasswordByAdminID(adminID int32, hashedPassword *string) error
	AdminGetProfileByAdminID(adminID int32) (*dto.AdminProfile, error)
}
