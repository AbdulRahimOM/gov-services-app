package usecase

import (
	"fmt"
	"time"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/response"
	jwttoken "github.com/AbdulRahimOM/gov-services-app/internal/jwt-token"
	hashpassword "github.com/AbdulRahimOM/gov-services-app/shared/hash-password"
	respCode "github.com/AbdulRahimOM/gov-services-app/shared/std-response/response-code"
)

func (uc UserUseCase) UserSignUpGetOTP(phone *string) (string, error) {
	//check if mobile number is registered
	doUserExists, err := uc.userRepo.CheckIfMobileIsRegistered(phone)
	if err != nil {
		return respCode.DBError, fmt.Errorf("at database error: failed to check if mobile is registered: %v", err)
	}
	if doUserExists {
		return respCode.MobileAlreadyRegistered, fmt.Errorf("mobile number already registered")
	}

	//send otp
	err = uc.twilioOTPClient.SendOtp(*phone)
	if err != nil {
		return respCode.OtherInternalError, fmt.Errorf("failed to send OTP: %v", err)
	}
	return "", nil
}

func (uc UserUseCase) VerifyOtpForSignUp(phone, otp *string) (*response.UserSignUp, string, error) {
	//check if mobile number is registered
	doUserExists, err := uc.userRepo.CheckIfMobileIsRegistered(phone)
	if err != nil {
		return nil, respCode.DBError, fmt.Errorf("at database error: failed to check if mobile is registered: %v", err)
	}
	if doUserExists {
		return nil, respCode.MobileAlreadyRegistered, fmt.Errorf("mobile number already registered")
	}

	//verify otp
	isVerified, err := uc.twilioOTPClient.VerifyOtp(*phone, *otp)
	if err != nil {
		return nil, respCode.OtherInternalError, fmt.Errorf("twilio error, failed to send OTP: %v", err)
	}

	if !isVerified {
		return nil, respCode.InvalidOTP, fmt.Errorf("invalid OTP")
	}

	//otp verified, so create temporary signing up user with only phone number
	userId, err := uc.userRepo.CreateSigningUpUser(phone, false)
	if err != nil {
		return nil, respCode.DBError, fmt.Errorf("at database: failed to create signing up user: %v", err)
	}

	//generate jwt token
	token, err := uc.jwtClient.GenerateToken(
		jwttoken.AccountInfo{
			Role: "password-not-set-user",
			Id:   userId,
		}, nil, time.Minute*time.Duration(config.JWT.ExpTimeInMinutes))
	if err != nil {
		return nil, respCode.OtherInternalError, fmt.Errorf("failed to generate token: %v", err)
	}

	return &response.UserSignUp{
		UserId: userId,
		Token:  token,
	}, "", nil
}

// SettingPwForNewUser
func (uc UserUseCase) SettingPwForNewUser(userID int32, newPassword *string) (*response.NewUserSetPw, string, error) {
	// //check if user exists	//no need as userID is taken from jwt token
	// doUserExists, err := uc.userRepo.CheckIfUserExists(&userID)
	// if err != nil {
	// 	log.Printf("at database error: failed to check if user exists: %v", err)
	// 	return nil, fmt.Errorf("at database error: failed to check if user exists: %v", err)
	// }
	// if !doUserExists {
	// 	log.Printf("user does not exist")
	// 	return nil, fmt.Errorf("user does not exist")
	// }

	hashedPassword, err := hashpassword.Hashpassword(*newPassword)
	if err != nil {
		return nil, respCode.OtherInternalError, fmt.Errorf("failed to hash password: %v", err)
	}

	//update password
	err = uc.userRepo.UpdatePassword(userID, &hashedPassword)
	if err != nil {
		return nil, respCode.DBError, fmt.Errorf("at database error: failed to update password: %v", err)
	}

	token, err := uc.jwtClient.GenerateToken(
		jwttoken.AccountInfo{
			Role: "user",
			Id:   userID,
		}, nil, time.Minute*time.Duration(config.JWT.ExpTimeInMinutes))
	if err != nil {
		return nil, respCode.OtherInternalError, fmt.Errorf("failed to generate token: %v", err)
	}

	return &response.NewUserSetPw{
		Token: token,
	}, "", nil
}
