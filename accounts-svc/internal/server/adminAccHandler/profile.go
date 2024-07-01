package adminAccHandler

import (
	"context"
	"log"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/shared/std-response/std-response"
	"google.golang.org/protobuf/types/known/emptypb"
)
// AdminGetProfile
func (s *AdminAccountsServer) AdminGetProfile(ctx context.Context, req *pb.AdminGetProfileRequest) (*pb.AdminGetProfileResponse, error) {
	resp, responseCode, err := s.AdminUseCase.AdminGetProfile(req.AdminId)
	if err != nil {
		log.Printf("failed to get profile: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}

	return &pb.AdminGetProfileResponse{
		FirstName:   resp.FirstName,
		LastName:    resp.LastName,
		Email:       resp.Email,
		Address:     resp.Address,
		Pincode:     resp.Pincode,
		PhoneNumber: resp.PhoneNumber,
	}, nil
}


// AdminUpdatePasswordUsingOldPw
func (s *AdminAccountsServer) AdminUpdatePasswordUsingOldPw(ctx context.Context, req *pb.AdminUpdatePasswordUsingOldPwRequest) (*emptypb.Empty, error) {
	responseCode, err := s.AdminUseCase.AdminUpdatePasswordUsingOldPw(&request.AdminUpdatePasswordUsingOldPw{
		AdminId:      req.AdminId,
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

//AdminUpdateProfile
func (s *AdminAccountsServer) AdminUpdateProfile(ctx context.Context, req *pb.AdminUpdateProfileRequest) (*emptypb.Empty, error) {
	responseCode, err := s.AdminUseCase.AdminUpdateProfile(&request.AdminUpdateProfile{
		AdminId:     req.AdminId,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Address:     req.Address,
		Pincode:     req.Pincode,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		log.Printf("failed to update profile: %v", err)
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}

	log.Println("Profile updated")
	return nil, nil
}