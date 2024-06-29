package usecase

import (
	"log"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/config"
	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/response"
	repo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository"
	jwttoken "github.com/AbdulRahimOM/gov-services-app/internal/jwt-token"
	twilioOTP "github.com/AbdulRahimOM/gov-services-app/shared/twilio"
)

// var (
// 	jwtClient    *jwttoken.TokenGenerator
// 	twilioOTPClient *twilioOTP.TwilioClient
// )

// func init() {
// 	var err error
// 	jwtClient, err = jwttoken.NewTokenGenerator("./internal/config/private.key", time.Minute*time.Duration(config.JWT.ExpTimeInMinutes))
// 	if err != nil {
// 		log.Fatalf("Failed to create token generator: %v", err)
// 	}

// 	twilioOTPClient = twilioOTP.NewTwilioClient(config.Twilio.AccountSid, config.Twilio.AuthToken, config.Twilio.ServiceSid)
// }

type IUserUC interface {
	//login
	UserLoginGetOTP(phone *string) (string, error)
	VerifyOtpForLogin(phone, otp *string) (*response.UserLogin, string, error)
	VerifyPasswordForLogin(phone, password *string) (*response.UserLogin, string, error)

	//signup
	UserSignUpGetOTP(phone *string) (string, error)
	VerifyOtpForSignUp(phone, otp *string) (*response.UserSignUp, string, error)
	SettingPwForNewUser(userID int32, newPassword *string) (*response.NewUserSetPw, string, error)

	//profile
	UserGetOTPForPwChange(userID int32) (*response.UserGetOTPForPwChange, string, error)
	UserVerifyOTPForPwChange(userID int32, otp *string) (*response.UserVerifyOTPForPwChange, string, error)
	UserSetNewPwAfterVerifyingOTP(userID int32, newPassword *string) (string, error)
	UserGetProfile(userID int32) (*dto.UserProfile, string, error)
	UserUpdateProfile(*request.UserUpdateProfile) (string, error)
	UserUpdatePasswordUsingOldPw(req *request.UserUpdatePasswordUsingOldPw) (string, error)
}

type UserUseCase struct {
	userRepo        repo.IUserRepo
	jwtClient       *jwttoken.TokenGenerator
	twilioOTPClient *twilioOTP.TwilioClient
}

func NewUserUseCase(userRepo repo.IUserRepo) IUserUC {
	jwtClient, err := jwttoken.NewTokenGenerator("./internal/config/private.key")
	if err != nil {
		log.Fatalf("Failed to create token generator: %v", err)
	}

	twilioOTPClient := twilioOTP.NewTwilioClient(config.Twilio.AccountSid, config.Twilio.AuthToken, config.Twilio.ServiceSid)
	return &UserUseCase{
		userRepo:        userRepo,
		jwtClient:       jwtClient,
		twilioOTPClient: twilioOTPClient,
	}
}
