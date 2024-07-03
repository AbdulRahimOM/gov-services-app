package adminAccHandler

import (
	"context"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
)

// AdminAddSubAdmin
func (s *AdminAccountsServer) AdminAddSubAdmin(c context.Context, req *pb.AdminAddSubAdminRequest) (*pb.AdminAddSubAdminResponse, error) {
	newSubAdminID, responseCode, err := s.AdminUseCase.AdminAddSubAdmin(&request.AdminAddSubAdmin{
		AdminID: req.AdminId,
		NewSubAdmin: request.NewSubAdmin{
			FirstName:   req.NewSubAdmin.FirstName,
			LastName:    req.NewSubAdmin.LastName,
			Email:       req.NewSubAdmin.Email,
			PhoneNumber: req.NewSubAdmin.PhoneNumber,
			DeptID:      req.NewSubAdmin.DeptId,
			Designation: req.NewSubAdmin.Designation,
			RankID:      req.NewSubAdmin.RankId,
			OfficeID:    req.NewSubAdmin.OfficeId,
		},
	})
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return &pb.AdminAddSubAdminResponse{
			NewSubAdminID: newSubAdminID,
		}, nil
	}
}
