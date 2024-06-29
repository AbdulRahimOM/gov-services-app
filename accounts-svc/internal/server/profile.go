package server

import (
	"context"
	"log"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/shared/std-response/std-response"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetOTPForPwChange
func (s *AccountsServer) GetOTPForPwChange(ctx context.Context, req *pb.GetOTPForPwChangeRequest) (*pb.GetOTPForPwChangeResponse, error) {
	resp, responseCode, err := s.UserUseCase.GetOTPForPwChange(req.UserId)
	if err != nil {
		log.Printf("failed to get OTP for password change: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}
	log.Println("OTP sent")
	return &pb.GetOTPForPwChangeResponse{
		Last4Digits: resp.Last4Digits,
	}, nil
}

// VerifyOTPForPwChange
func (s *AccountsServer) VerifyOTPForPwChange(ctx context.Context, req *pb.VerifyOTPForPwChangeRequest) (*pb.VerifyOTPForPwChangeResponse, error) {
	resp, responseCode, err := s.UserUseCase.VerifyOTPForPwChange(req.UserId, &req.Otp)
	if err != nil {
		log.Printf("failed to verify OTP for password change: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}
	log.Println("OTP verified")
	return &pb.VerifyOTPForPwChangeResponse{
		TempToken: resp.TempToken,
	}, nil
}

//GetProfile
func (s *AccountsServer) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	resp, responseCode, err := s.UserUseCase.GetProfile(req.UserId)
	if err != nil {
		log.Printf("failed to get profile: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}
	
	return &pb.GetProfileResponse{
		FirstName: resp.FirstName,
		LastName:  resp.LastName,
		Email:    resp.Email,
		Address:  resp.Address,
		Pincode: resp.Pincode,
		PhoneNumber: resp.PhoneNumber,
	}, nil	
}

// UpdateProfile
func (s *AccountsServer) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*emptypb.Empty, error) {
	responseCode, err := s.UserUseCase.UpdateProfile(&request.UpdateProfile{
		UserId:   req.UserId,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:    req.Email,
		Address:  req.Address,
		Pincode: req.Pincode,
	})
	if err != nil {
		log.Printf("failed to update profile: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}

	log.Println("Profile updated")
	return nil, nil
}

//UpdatePasswordUsingOldPw
func (s *AccountsServer) UpdatePasswordUsingOldPw(ctx context.Context, req *pb.UpdatePasswordUsingOldPwRequest) (*emptypb.Empty, error) {
	responseCode, err := s.UserUseCase.UpdatePasswordUsingOldPw(&request.UpdatePasswordUsingOldPw{
		UserId:   req.UserId,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		log.Printf("failed to update password: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}

	log.Println("Password updated")
	return nil, nil
}

//SetNewPwAfterVerifyingOTP
func (s *AccountsServer) SetNewPwAfterVerifyingOTP(ctx context.Context, req *pb.SetNewPwAfterVerifyingOTPRequest) (*emptypb.Empty, error) {
	responseCode, err := s.UserUseCase.SetNewPwAfterVerifyingOTP(req.UserId, &req.NewPassword)
	if err != nil {
		log.Printf("failed to set new password: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}

	log.Println("Password set")
	return nil, nil
}