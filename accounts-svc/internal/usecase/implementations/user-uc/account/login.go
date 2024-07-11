package userAccUC

import (
	"fmt"
	"time"

	hashpassword "github.com/AbdulRahimOM/go-utils/hashPassword"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/response"
	jwttoken "github.com/AbdulRahimOM/gov-services-app/internal/jwt-token"
	responsecode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (uc UserUseCase) UserLoginGetOTP(phone *string) (string, error) {
	//check if phoneNumber number is registered
	doUserExists, err := uc.userRepo.CheckIfPhoneNumberIsRegistered(phone)
	if err != nil {
		return responsecode.DBError, fmt.Errorf("at database error: failed to check if phoneNumber is registered: %v", err)
	}

	if !doUserExists {
		return responsecode.PhoneNumberNotRegistered, fmt.Errorf("phoneNumber number not registered")
	}

	//send otp
	err = uc.twilioOTPClient.SendOtp(*phone)
	if err != nil {
		return responsecode.OtherInternalError, fmt.Errorf("failed to send OTP: %v", err)
	}

	return "", nil
}

func (uc UserUseCase) VerifyOtpForLogin(phone, otp *string) (*response.UserLogin, string, error) {
	//verify otp
	isVerified, err := uc.twilioOTPClient.VerifyOtp(*phone, *otp)
	if err != nil {
		return nil, responsecode.OtherInternalError, fmt.Errorf("twilio error, failed to send OTP: %v", err)
	}

	if !isVerified {
		return nil, responsecode.InvalidOTP, fmt.Errorf("invalid OTP")
	}

	//get user details
	user, err := uc.userRepo.GetUserByPhoneNumber(phone)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, responsecode.CorruptRequest, fmt.Errorf("user not registered with this phoneNumber number, but attempted to verify otp for login")
		} else {
			return nil, responsecode.DBError, fmt.Errorf("at database: failed to get user details: %v", err)
		}
	}

	//generate jwt token
	token, err := uc.jwtClient.GenerateToken(
		jwttoken.AccountInfo{
			Role: "user",
			Id:   user.ID,
		}, nil, time.Minute*time.Duration(config.JWT.ExpTimeInMinutes))
	if err != nil {
		return nil, responsecode.OtherInternalError, fmt.Errorf("failed to generate token: %v", err)
	}

	return &response.UserLogin{
		User:  user,
		Token: token,
	}, "", nil
}

// VerifyPasswordForLogin
func (uc UserUseCase) VerifyPasswordForLogin(phone, password *string) (*response.UserLogin, string, error) {
	//verify password
	user, hashedPw, err := uc.userRepo.GetUserWithPasswordByPhoneNumber(phone)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, responsecode.CorruptRequest, fmt.Errorf("no user registered with this phoneNumber number")
		} else {
			return nil, responsecode.DBError, fmt.Errorf("at database: failed to get user password: %v", err)
		}
	}

	err = hashpassword.CompareHashedPassword(*hashedPw, *password)
	if err != nil {
		return nil, responsecode.InvalidPassword, fmt.Errorf("pw verification failed: %v", err)
	}

	//generate jwt token
	token, err := uc.jwtClient.GenerateToken(
		jwttoken.AccountInfo{
			Role: "user",
			Id:   user.ID,
		}, nil, time.Minute*time.Duration(config.JWT.ExpTimeInMinutes))
	if err != nil {
		return nil, responsecode.OtherInternalError, fmt.Errorf("failed to generate token: %v", err)
	}

	return &response.UserLogin{
		User:  user,
		Token: token,
	}, "", nil
}
