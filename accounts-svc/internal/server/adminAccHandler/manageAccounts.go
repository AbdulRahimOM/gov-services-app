package adminAccHandler

import (
	"context"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
)

func (s *AdminAccountsServer) AdminGetAdmins(ctx context.Context, req *pb.AdminGetAdminsRequest) (*pb.AdminGetAdminsResponse, error) {
	admins, responseCode, err := s.AdminUseCase.AdminGetAdmins(req.AdminId, &request.AdminSearchCriteria{
		FirstName:   req.SearchCriteria.FirstName,
		LastName:    req.SearchCriteria.LastName,
		Email:       req.SearchCriteria.Email,
		PhoneNumber: req.SearchCriteria.PhoneNumber,
		Designation: req.SearchCriteria.Designation,
		OfficeId:    req.SearchCriteria.OfficeId,
	})
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		adminsResponse := []*pb.Admin{}
		for _, admin := range *admins {
			adminsResponse = append(adminsResponse, &pb.Admin{
				Id:          admin.ID,
				FirstName:   admin.FirstName,
				LastName:    admin.LastName,
				Email:       admin.Email,
				Address:     admin.Address,
				Pincode:     admin.Pincode,
				PhoneNumber: admin.PhoneNumber,
				Designation: admin.Designation,
				OfficeId:    admin.OfficeId,
			})
		}
		return &pb.AdminGetAdminsResponse{
			Admin: adminsResponse,
		}, nil
	}
}
