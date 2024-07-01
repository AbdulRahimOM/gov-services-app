package userAccHandler

import (
	"context"
	"log"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/shared/std-response/std-response"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UserGetOTPForPwChange
func (s *UserAccountsServer) UserGetOTPForPwChange(ctx context.Context, req *pb.UserGetOTPForPwChangeRequest) (*pb.UserGetOTPForPwChangeResponse, error) {
	resp, responseCode, err := s.UserUseCase.UserGetOTPForPwChange(req.UserId)
	if err != nil {
		log.Printf("failed to get OTP for password change: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}
	log.Println("OTP sent")
	return &pb.UserGetOTPForPwChangeResponse{
		Last4Digits: resp.Last4Digits,
	}, nil
}

// UserVerifyOTPForPwChange
func (s *UserAccountsServer) UserVerifyOTPForPwChange(ctx context.Context, req *pb.UserVerifyOTPForPwChangeRequest) (*pb.UserVerifyOTPForPwChangeResponse, error) {
	resp, responseCode, err := s.UserUseCase.UserVerifyOTPForPwChange(req.UserId, &req.Otp)
	if err != nil {
		log.Printf("failed to verify OTP for password change: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}
	log.Println("OTP verified")
	return &pb.UserVerifyOTPForPwChangeResponse{
		TempToken: resp.TempToken,
	}, nil
}

// UserGetProfile
func (s *UserAccountsServer) UserGetProfile(ctx context.Context, req *pb.UserGetProfileRequest) (*pb.UserGetProfileResponse, error) {
	resp, responseCode, err := s.UserUseCase.UserGetProfile(req.UserId)
	if err != nil {
		log.Printf("failed to get profile: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}

	return &pb.UserGetProfileResponse{
		FirstName:   resp.FirstName,
		LastName:    resp.LastName,
		Email:       resp.Email,
		Address:     resp.Address,
		Pincode:     resp.Pincode,
		PhoneNumber: resp.PhoneNumber,
	}, nil
}

// UserUpdateProfile
func (s *UserAccountsServer) UserUpdateProfile(ctx context.Context, req *pb.UserUpdateProfileRequest) (*emptypb.Empty, error) {
	responseCode, err := s.UserUseCase.UserUpdateProfile(&request.UserUpdateProfile{
		UserId:    req.UserId,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Address:   req.Address,
		Pincode:   req.Pincode,
	})
	if err != nil {
		log.Printf("failed to update profile: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}

	log.Println("Profile updated")
	return nil, nil
}

// UserUpdatePasswordUsingOldPw
func (s *UserAccountsServer) UserUpdatePasswordUsingOldPw(ctx context.Context, req *pb.UserUpdatePasswordUsingOldPwRequest) (*emptypb.Empty, error) {
	responseCode, err := s.UserUseCase.UserUpdatePasswordUsingOldPw(&request.UserUpdatePasswordUsingOldPw{
		UserId:      req.UserId,
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

// UserSetNewPwAfterVerifyingOTP
func (s *UserAccountsServer) UserSetNewPwAfterVerifyingOTP(ctx context.Context, req *pb.UserSetNewPwAfterVerifyingOTPRequest) (*emptypb.Empty, error) {
	responseCode, err := s.UserUseCase.UserSetNewPwAfterVerifyingOTP(req.UserId, &req.NewPassword)
	if err != nil {
		log.Printf("failed to set new password: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}

	log.Println("Password set")
	return nil, nil
}
