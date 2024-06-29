package server

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase"
	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"
	respCode "github.com/AbdulRahimOM/gov-services-app/shared/std-response/response-code"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/shared/std-response/std-response"
)

type AccountsServer struct {
	UserUseCase usecase.IUserUC
	pb.UnimplementedAccountServiceServer
}

// UserLoginGetOTP
func (s *AccountsServer) UserSignUpGetOTP(ctx context.Context, req *pb.UserSignUpGetOTPRequest) (*pb.UserSignUpGetOTPResponse, error) {
	fmt.Println("UserSignUpGetOTP")
	//checking if code is india code
	if !strings.HasPrefix(req.PhoneNumber, "+91") {
		log.Println("Phone number must start with +91")
		return nil, stdresponse.GetGrpcStatus(respCode.ValidationError, "Phone number must start with +91")
	}

	msg, err := s.UserUseCase.UserSignUpGetOTP(&req.PhoneNumber)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(respCode.OtherInternalError, err.Error())
	} else {
		log.Println("OTP sent")
		return &pb.UserSignUpGetOTPResponse{
			Message: msg,
		}, nil
	}
}

// VerifyOTPForSignUp
func (s *AccountsServer) UserSignUpVerifyOTP(ctx context.Context, req *pb.UserSignUpVerifyOTPRequest) (*pb.UserSignUpVerifyOTPResponse, error) {
	//checking if code is india code
	if !strings.HasPrefix(req.PhoneNumber, "+91") {
		log.Println("Phone number must start with +91")
		return nil, stdresponse.GetGrpcStatus(respCode.ValidationError, "Phone number must start with +91")
	}
	resp, responseCode, err := s.UserUseCase.VerifyOtpForSignUp(&req.PhoneNumber, &req.Otp)
	if err != nil {
		log.Printf("failed to verify OTP for sign up: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		log.Println("OTP verified")
		return &pb.UserSignUpVerifyOTPResponse{
			Message: "User signed up successfully",
			Token:   *resp.Token,
			UserDetails: &pb.SignedUpUserDetails{
				Id:          resp.UserId,
				PhoneNumber: req.PhoneNumber,
			},
		}, nil
	}
}

// SignedUpUserSettingPw
func (s *AccountsServer) SignedUpUserSettingPw(ctx context.Context, req *pb.SignedUpUserSettingPwRequest) (*pb.SignedUpUserSettingPwResponse, error) {
	resp, responseCode, err := s.UserUseCase.SettingPwForNewUser(req.UserId, &req.NewPassword)
	if err != nil {
		log.Printf("failed to set password for signed up user: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		log.Println("Password set")
		return &pb.SignedUpUserSettingPwResponse{
			Token: *resp.Token,
		}, nil
	}
}
