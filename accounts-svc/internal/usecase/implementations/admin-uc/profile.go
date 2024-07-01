package adminuc

import (
	"fmt"

	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	hashpassword "github.com/AbdulRahimOM/go-utils/hashPassword"
	respcode "github.com/AbdulRahimOM/gov-services-app/shared/std-response/response-code"
)

// AdminGetProfile
func (u *AdminUseCase) AdminGetProfile(adminID int32) (*dto.AdminProfile, string, error) {
	profile, err := u.adminRepo.AdminGetProfileByAdminID(adminID)
	if err != nil {
		return nil, respcode.DBError, fmt.Errorf("failed to get profile: %v", err)
	}

	return profile, "", nil
}

// AdminUpdatePasswordUsingOldPw
func (u *AdminUseCase) AdminUpdatePasswordUsingOldPw(req *request.AdminUpdatePasswordUsingOldPw) (string, error) {
	//get admin details
	hashedPw, err := u.adminRepo.GetPasswordByAdminID(req.AdminId)
	if err != nil {
		return respcode.DBError, fmt.Errorf("@db error: %v", err)
	}

	//verify old password
	err = hashpassword.CompareHashedPassword(*hashedPw, req.OldPassword)
	if err != nil {
		return respcode.InvalidPassword, fmt.Errorf("old password verification failed: %v", err)
	}

	//hash new password
	hashedNewPassword, err := hashpassword.Hashpassword(req.NewPassword)
	if err != nil {
		return respcode.OtherInternalError, fmt.Errorf("failed to hash password: %v", err)
	}

	//update password
	err = u.adminRepo.UpdatePasswordByAdminID(req.AdminId, &hashedNewPassword)
	if err != nil {
		return respcode.DBError, fmt.Errorf("failed to update password: %v", err)
	}

	return "", nil
}

// AdminUpdateProfile
func (u *AdminUseCase) AdminUpdateProfile(req *request.AdminUpdateProfile) (string, error) {
	err := u.adminRepo.AdminUpdateProfile(req)
	if err != nil {
		return respcode.DBError, fmt.Errorf("failed to update profile: %v", err)
	}

	return "", nil
}