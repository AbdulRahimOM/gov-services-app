package adminAccHandler

import (
	"context"

	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
)

// AdminGetOffices
func (s *AdminAccountsServer) AdminGetOffices(c context.Context, req *pb.AdminGetOfficesRequest) (*pb.AdminGetOfficesResponse, error) {
	offices, responseCode, err := s.AdminUseCase.AdminGetOffices()
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		var pbOffices []*pb.Office = make([]*pb.Office, len(*offices))
		for i, office := range *offices {
			pbOffices[i] = &pb.Office{
				Id:               office.ID,
				DeptId:           office.DeptID,
				HierarchyRank:    office.HierarchyRank,
				RegionName:       office.RegionName,
				HeadOfficerId:    office.HeadOfficerID,
				OfficeLocation:   office.OfficeLocation,
				SuperiorOfficeId: office.SuperiorOfficeID,
			}
		}
		return &pb.AdminGetOfficesResponse{
			Office: pbOffices,
		}, nil
	}
}
