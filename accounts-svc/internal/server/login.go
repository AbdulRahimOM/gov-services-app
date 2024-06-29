package server

import (
	"context"
	"fmt"
	"log"
	"strings"

	// pb "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/infrastructure/pb/generated"
	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"
	respCode "github.com/AbdulRahimOM/gov-services-app/shared/std-response/response-code"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/shared/std-response/std-response"
)

// UserLoginGetOTP
func (s *UserAccountsServer) UserLoginGetOTP(ctx context.Context, req *pb.UserLoginGetOTPRequest) (*pb.UserLoginGetOTPResponse, error) {
	fmt.Println("UserLoginGetOTP")
	//checking if code is india code
	if !strings.HasPrefix(req.PhoneNumber, "+91") {
		log.Println("Phone number must start with +91")
		return nil, stdresponse.GetGrpcStatus(respCode.ValidationError, "Phone number must start with +91")
	}

	responseCode, err := s.UserUseCase.UserLoginGetOTP(&req.PhoneNumber)
	if err != nil {
		log.Printf("failed to get OTP for login: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		log.Println("OTP sent")
		return &pb.UserLoginGetOTPResponse{
			// Message: msg,
		}, nil
	}
}

// VerifyOTPForLogin
func (s *UserAccountsServer) UserLoginVerifyOTP(ctx context.Context, req *pb.UserLoginVerifyOTPRequest) (*pb.UserLoginResponse, error) {
	//checking if code is india code
	if !strings.HasPrefix(req.PhoneNumber, "+91") {
		return nil, stdresponse.GetGrpcStatus(respCode.ValidationError, "Phone number must start with +91")
	}

	resp, responseCode, err := s.UserUseCase.VerifyOtpForLogin(&req.PhoneNumber, &req.Otp)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return &pb.UserLoginResponse{
			Token: *resp.Token,
			UserDetails: &pb.LoggingUserDetails{
				Id:          resp.User.ID,
				FirstName:   resp.User.FName,
				LastName:    resp.User.LName,
				PhoneNumber: req.PhoneNumber,
			},
		}, nil
	}
}

// UserLoginViaPassword
func (s *UserAccountsServer) UserLoginViaPassword(ctx context.Context, req *pb.UserLoginViaPasswordRequest) (*pb.UserLoginResponse, error) {
	//checking if code is india code
	if !strings.HasPrefix(req.PhoneNumber, "+91") {
		return nil, stdresponse.GetGrpcStatus(respCode.ValidationError, "Phone number must start with +91")
	}

	resp, responseCode, err := s.UserUseCase.VerifyPasswordForLogin(&req.PhoneNumber, &req.Password)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return &pb.UserLoginResponse{
			Token: *resp.Token,
			UserDetails: &pb.LoggingUserDetails{
				Id:          resp.User.ID,
				FirstName:   resp.User.FName,
				LastName:    resp.User.LName,
				PhoneNumber: req.PhoneNumber,
			},
		}, nil
	}
}
