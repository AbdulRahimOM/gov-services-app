package adminAccHandler

import (
	"context"

	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/shared/std-response/std-response"
)

// AdminLoginViaPassword
func (s *AdminAccountsServer) AdminLoginViaPassword(ctx context.Context, req *pb.AdminLoginViaPasswordRequest) (*pb.AdminLoginResponse, error) {
	//checking if code is india code
	resp, responseCode, err := s.AdminUseCase.VerifyPasswordForLogin(&req.Username, &req.Password)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return &pb.AdminLoginResponse{
			Token: *resp.Token,
			AdminDetails: &pb.LoggingAdminDetails{
				Id:          resp.Admin.ID,
				FirstName:   resp.Admin.FName,
				LastName:    resp.Admin.LName,
				PhoneNumber: resp.Admin.PhoneNumber,
			},
		}, nil
	}
}
