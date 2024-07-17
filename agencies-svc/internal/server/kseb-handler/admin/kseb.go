package ksebAdminHandler

import (
	"context"
	"time"

	"github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/domain/dto/request"
	ucinterface "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/usecase/interface"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"google.golang.org/protobuf/types/known/emptypb"
)

type KSEBAgencyAdminServer struct {
	KsebUseCase ucinterface.IKsebAgencyAdminUC
	pb.UnimplementedKSEBAgencyAdminServiceServer
	// messages map[int32][]*pb.ChatResponse
	// mu       sync.Mutex
}

func NewKSEBAgencyAdminServer(ksebUseCase ucinterface.IKsebAgencyAdminUC) *KSEBAgencyAdminServer {
	return &KSEBAgencyAdminServer{
		KsebUseCase: ksebUseCase,
	}
}

func (k *KSEBAgencyAdminServer) RegisterSectionCode(ctx context.Context, req *pb.RegisterSectionCodeRequest) (*emptypb.Empty, error) {
	regSectionCodeReq := requests.KsebRegSectionCode{
		SectionCode: req.SectionCode,
		OfficeId:    req.OfficeId,
	}
	_, responseCode, err := k.KsebUseCase.RegisterSectionCode(req.AdminId, &regSectionCodeReq)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return nil, nil
	}
}

func (k *KSEBAgencyAdminServer) GetComplaints(ctx context.Context, req *pb.GetComplaintsRequest) (*pb.GetComplaintsResponse, error) {
	searchCriteria := request.KsebComplaintSearchCriteria{
		Status:        req.SearchCriteria.Status,
		AttenderScope: req.SearchCriteria.AttenderScope,
	}
	complaints, responseCode, err := k.KsebUseCase.GetComplaints(req.AdminId, &searchCriteria)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		complaintsResponse := make([]*pb.KsebComplaint, len(*complaints))
		for i, v := range *complaints {
			complaintsResponse[i] = &pb.KsebComplaint{
				ID:             v.ID,
				UserID:         v.UserID,
				Type:           v.Type,
				Title:          v.Title,
				Description:    v.Description,
				ConsumerNumber: v.ConsumerNumber,
				AttenderID:     v.AttenderID,
				Status:         v.Status,
				CreatedAt:      v.CreatedAt.Format(time.RFC3339),
				Remarks:        v.Remarks,
				ClosedAt:       v.ClosedAt.Format(time.RFC3339),
			}
		}
		return &pb.GetComplaintsResponse{
			Complaints: complaintsResponse,
		}, nil
	}
}

func (k *KSEBAgencyAdminServer) OpenComplaint(ctx context.Context, req *pb.OpenComplaintRequest) (*pb.KsebComplaint, error) {
	complaint, responseCode, err := k.KsebUseCase.OpenComplaint(req.AdminId, req.ComplaintId)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return &pb.KsebComplaint{
			ID:             complaint.ID,
			UserID:         complaint.UserID,
			Type:           complaint.Type,
			Title:          complaint.Title,
			Description:    complaint.Description,
			ConsumerNumber: complaint.ConsumerNumber,
			AttenderID:     complaint.AttenderID,
			Status:         complaint.Status,
			CreatedAt:      complaint.CreatedAt.Format(time.RFC3339),
			Remarks:        complaint.Remarks,
			ClosedAt:       complaint.ClosedAt.Format(time.RFC3339),
		}, nil
	}
}

func (k *KSEBAgencyAdminServer) CloseComplaint(ctx context.Context, req *pb.CloseComplaintRequest) (*emptypb.Empty, error) {
	responseCode, err := k.KsebUseCase.CloseComplaint(req.AdminId, req.ComplaintId, req.Remarks)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return nil, nil
	}
}

func (k *KSEBAgencyAdminServer) CheckIfComplaintAccessibleToAdmin(ctx context.Context, req *pb.CheckIfComplaintAccessibleToAdminRequest) (*emptypb.Empty, error) {
	_, responseCode, err := k.KsebUseCase.CheckIfComplaintBeAccessibleToAdmin(req.AdminId, req.ComplaintId)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return nil, nil
	}
}
