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

// GetOTPForLogin
func (s *AccountsServer) GetOTPForSignUp(ctx context.Context, req *pb.GetOTPForSignUpRequest) (*pb.GetOTPForSignUpResponse, error) {
	fmt.Println("GetOTPForSignUp")
	//checking if code is india code
	if !strings.HasPrefix(req.PhoneNumber, "+91") {
		log.Println("Phone number must start with +91")
		return nil, stdresponse.GetGrpcStatus(respCode.ValidationError, "Phone number must start with +91")
	}

	msg, err := s.UserUseCase.GetOTPForSignUp(&req.PhoneNumber)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(respCode.OtherInternalError, err.Error())
	} else {
		log.Println("OTP sent")
		return &pb.GetOTPForSignUpResponse{
			Message: msg,
		}, nil
	}
}

// VerifyOTPForSignUp
func (s *AccountsServer) UserSignUpViaOTP(ctx context.Context, req *pb.UserSignUpViaOTPRequest) (*pb.UserSignUpViaOTPResponse, error) {
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
		return &pb.UserSignUpViaOTPResponse{
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
