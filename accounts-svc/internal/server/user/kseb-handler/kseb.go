package ksebUserHandler

import (
	"context"

	ucinterface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	ksebpb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated/ksebpb"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"google.golang.org/protobuf/types/known/emptypb"
)

type KSEBUserServer struct {
	KsebUseCase ucinterface.IKsebUserUC
	ksebpb.UnimplementedKSEBUserServiceServer
}

func NewKSEBUserServer(ksebUseCase ucinterface.IKsebUserUC) *KSEBUserServer {
	return &KSEBUserServer{
		KsebUseCase: ksebUseCase,
	}
}

func (k *KSEBUserServer) AddConsumerNumber(ctx context.Context, req *ksebpb.AddConsumerNumberRequest) (*emptypb.Empty, error) {
	responseCode, err := k.KsebUseCase.AddConsumerNumber(req.UserId, req.ConsumerNumber, req.NickName)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return nil, nil
	}
}

func (k *KSEBUserServer) GetUserConsumerNumbers(ctx context.Context, req *ksebpb.GetUserConsumerNumbersRequest) (*ksebpb.GetUserConsumerNumbersResponse, error) {
	resp, responseCode, err := k.KsebUseCase.GetUserConsumerNumbers(req.UserId)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		conns := make([]*ksebpb.ConsumerNumber, len(*resp))
		for i, v := range *resp {
			conns[i] = &ksebpb.ConsumerNumber{
				Id:             v.Id,
				ConsumerNumber: v.ConsumerNumber,
				NickName:       v.NickName,
			}
		}

		return &ksebpb.GetUserConsumerNumbersResponse{
			ConsumerNumbers: conns,
		}, nil
	}
}

func (k *KSEBUserServer) RaiseComplaint(ctx context.Context, req *ksebpb.RaiseComplaintRequest) (*ksebpb.RaiseComplaintResponse, error) {
	complaint:=requests.KSEBComplaint{
		Type: req.Complaint.Type,
		Category: req.Complaint.Category,
		Title: req.Complaint.Title,
		Description: req.Complaint.Description,
		ConsumerNumber: req.Complaint.ConsumerNumber,
	}
	complaintId,responseCode, err := k.KsebUseCase.RaiseComplaint(req.UserId, &complaint)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return &ksebpb.RaiseComplaintResponse{
			ComplaintId: complaintId,
		}, nil
	}
}