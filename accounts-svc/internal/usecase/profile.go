package usecase

import (
	"fmt"
	"time"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/config"
	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/response"
	jwttoken "github.com/AbdulRahimOM/gov-services-app/internal/jwt-token"
	hashpassword "github.com/AbdulRahimOM/gov-services-app/shared/hash-password"
	respcode "github.com/AbdulRahimOM/gov-services-app/shared/std-response/response-code"
	"github.com/AbdulRahimOM/gov-services-app/shared/tag"
)

// UserGetOTPForPwChange
func (u *UserUseCase) UserGetOTPForPwChange(userID int32) (*response.UserGetOTPForPwChange, string, error) {
	mobile, err := u.userRepo.GetMobileByUserID(userID)
	if err != nil {
		return nil, respcode.DBError, fmt.Errorf("failed to get mobile number: %v", err)
	}

	err = u.twilioOTPClient.SendOtp("+" + mobile)
	if err != nil {
		return nil, respcode.OtherInternalError, fmt.Errorf("failed to send OTP: %v", err)
	}

	return &response.UserGetOTPForPwChange{
		Last4Digits: mobile[len(mobile)-4:],
	}, "", nil
}

// UserVerifyOTPForPwChange
func (u *UserUseCase) UserVerifyOTPForPwChange(userID int32, otp *string) (*response.UserVerifyOTPForPwChange, string, error) {
	mobile, err := u.userRepo.GetMobileByUserID(userID)
	if err != nil {
		return nil, respcode.DBError, fmt.Errorf("failed to get mobile number: %v", err)
	}

	isVerified, err := u.twilioOTPClient.VerifyOtp(mobile, *otp)
	if err != nil {
		return nil, respcode.OtherInternalError, fmt.Errorf("failed to verify OTP: %v", err)
	}

	if !isVerified {
		return nil, respcode.InvalidOTP, fmt.Errorf("invalid OTP")
	}

	//create temporary token
	token, err := u.jwtClient.GenerateToken(
		jwttoken.AccountInfo{
			Role: "user",
			Id:   userID,
		},
		&jwttoken.ExtraPurposeInfo{
			Purpose:    tag.PwChange,
			ExpiryTime: time.Now().Add(time.Minute * time.Duration(config.JWT.ExpTimeInMinutes)),
		},

		time.Minute*time.Duration(config.JWT.ExpTimeInMinutes),
	)
	if err != nil {
		return nil, respcode.OtherInternalError, fmt.Errorf("failed to generate token: %v", err)
	}

	return &response.UserVerifyOTPForPwChange{
		TempToken: *token,
	}, "", nil
}

// UserGetProfile
func (u *UserUseCase) UserGetProfile(userID int32) (*dto.UserProfile, string, error) {
	profile, err := u.userRepo.UserGetProfileByUserID(userID)
	if err != nil {
		return nil, respcode.DBError, fmt.Errorf("failed to get profile: %v", err)
	}

	return profile, "", nil
}

// UserUpdateProfile
func (u *UserUseCase) UserUpdateProfile(req *request.UserUpdateProfile) (string, error) {
	err := u.userRepo.UserUpdateProfile(req)
	if err != nil {
		return respcode.DBError, fmt.Errorf("failed to update profile: %v", err)
	}

	return "", nil
}

// UserUpdatePasswordUsingOldPw
func (u *UserUseCase) UserUpdatePasswordUsingOldPw(req *request.UserUpdatePasswordUsingOldPw) (string, error) {
	//get user details
	hashedPw, err := u.userRepo.GetPasswordByUserID(req.UserId)
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
	err = u.userRepo.UpdatePasswordByUserID(req.UserId, &hashedNewPassword)
	if err != nil {
		return respcode.DBError, fmt.Errorf("failed to update password: %v", err)
	}

	return "", nil
}

// UserSetNewPwAfterVerifyingOTP
func (u *UserUseCase) UserSetNewPwAfterVerifyingOTP(userID int32, newPassword *string) (string, error) {
	//hash new password
	hashedNewPassword, err := hashpassword.Hashpassword(*newPassword)
	if err != nil {
		return respcode.OtherInternalError, fmt.Errorf("failed to hash password: %v", err)
	}

	//update password
	err = u.userRepo.UpdatePasswordByUserID(userID, &hashedNewPassword)
	if err != nil {
		return respcode.DBError, fmt.Errorf("failed to update password: %v", err)
	}

	return "", nil
}
