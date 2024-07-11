package adminAccHandler

import (
	"context"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
)

// AdminGetOffices
func (s *AdminAccountsServer) AdminGetOffices(c context.Context, req *pb.OfficeSearchCriteria) (*pb.AdminGetOfficesResponse, error) {
	searchCriteria := request.OfficeSearchCriteria{
		Name:             req.Name,
		Address:          req.Address,
		Id:               req.Id,
		DeptID:           req.DeptID,
		Rank:             req.Rank,
		SuperiorOfficeID: req.SuperiorOfficeID,
	}
	offices, responseCode, err := s.AdminUseCase.AdminGetOffices(&searchCriteria)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		var pbOffices []*pb.Office = make([]*pb.Office, len(*offices))
		for i, office := range *offices {
			pbOffices[i] = &pb.Office{
				Id:               office.ID,
				Name:             office.Name,
				DeptId:           office.DeptID,
				Rank:             office.Rank,
				Address:          office.Address,
				SuperiorOfficeId: office.SuperiorOfficeID,
			}
		}
		return &pb.AdminGetOfficesResponse{
			Office: pbOffices,
		}, nil
	}
}
