package adminAccHandler

import (
	"context"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"google.golang.org/protobuf/types/known/emptypb"
)

// AdminGetDepts
func (s *AdminAccountsServer) AdminGetDepts(c context.Context, req *emptypb.Empty) (*pb.AdminGetDeptsResponse, error) {
	depts, responseCode, err := s.AdminUseCase.AdminGetDepts()
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		var pbDepts []*pb.Dept = make([]*pb.Dept, len(*depts))
		for i, dept := range *depts {
			pbDepts[i] = &pb.Dept{
				Id:          dept.ID,
				Name:        dept.Name,
				Description: dept.Description,
			}
		}
		return &pb.AdminGetDeptsResponse{
			Dept: pbDepts,
		}, nil
	}
}

// AdminAddDept
func (s *AdminAccountsServer) AdminAddDept(c context.Context, req *pb.AdminAddDeptRequest) (*pb.AdminAddDeptResponse, error) {
	newDeptID, responseCode, err := s.AdminUseCase.AdminAddDept(request.AdminAddDept{
		AdminID: req.AdminId,
		NewDept: request.NewDept{
			Name:        req.NewDept.Name,
			Description: req.NewDept.Description,
		},
	})
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return &pb.AdminAddDeptResponse{
			NewDeptID: newDeptID,
		}, nil
	}

}
