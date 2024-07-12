package ucinterface

import (
	dto "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/other-dto"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/response"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
)

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

type IKsebUserUC interface {
	AddConsumerNumber(userID int32, consumerNumber, nickName string) (string, error)
	GetUserConsumerNumbers(userID int32) (*[]commondto.UserConsumerNumber, string, error)

	//complaint
	RaiseComplaint(userID int32, complaint *requests.KSEBComplaint) (int32,string, error)
}
