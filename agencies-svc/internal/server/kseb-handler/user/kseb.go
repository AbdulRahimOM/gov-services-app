package ksebUserHandler

import (
	"context"

	ucinterface "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/usecase/interface"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

type KSEBAgencyUserServer struct {
	KsebUseCase ucinterface.IKsebAgencyUserUC
	pb.UnimplementedKSEBAgencyUserServiceServer
	userChatStreams map[int32]chan *pb.ChatMessage
	getGrpcStatus   func(respCode string, errMsg string) error
}

func NewKSEBAgencyUserServer(ksebUseCase ucinterface.IKsebAgencyUserUC, logger *logrus.Entry) *KSEBAgencyUserServer {
	return &KSEBAgencyUserServer{
		KsebUseCase:     ksebUseCase,
		userChatStreams: make(map[int32]chan *pb.ChatMessage),
		getGrpcStatus:   stdresponse.NewGetGrpcStatusForService("agencies-svc", logger),
	}
}

func (k *KSEBAgencyUserServer) AddConsumerNumber(ctx context.Context, req *pb.AddConsumerNumberRequest) (*emptypb.Empty, error) {
	responseCode, err := k.KsebUseCase.AddConsumerNumber(req.UserId, req.ConsumerNumber, req.NickName)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return nil, nil
	}
}

func (k *KSEBAgencyUserServer) GetUserConsumerNumbers(ctx context.Context, req *pb.GetUserConsumerNumbersRequest) (*pb.GetUserConsumerNumbersResponse, error) {
	resp, responseCode, err := k.KsebUseCase.GetUserConsumerNumbers(req.UserId)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		conns := make([]*pb.ConsumerNumber, len(*resp))
		for i, v := range *resp {
			conns[i] = &pb.ConsumerNumber{
				Id:             v.Id,
				ConsumerNumber: v.ConsumerNumber,
				NickName:       v.NickName,
			}
		}

		return &pb.GetUserConsumerNumbersResponse{
			ConsumerNumbers: conns,
		}, nil
	}
}

func (k *KSEBAgencyUserServer) RaiseComplaint(ctx context.Context, req *pb.RaiseComplaintRequest) (*pb.RaiseComplaintResponse, error) {
	complaint := requests.KSEBComplaint{
		Type:           req.Complaint.Type,
		Category:       req.Complaint.Category,
		Title:          req.Complaint.Title,
		Description:    req.Complaint.Description,
		ConsumerNumber: req.Complaint.ConsumerNumber,
	}
	complaintId, responseCode, err := k.KsebUseCase.RaiseComplaint(req.UserId, &complaint)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return &pb.RaiseComplaintResponse{
			ComplaintId: complaintId,
		}, nil
	}
}

func (k *KSEBAgencyUserServer) CheckIfComplaintBelongsToUser(ctx context.Context, req *pb.CheckIfComplaintBelongsToUserRequest) (*emptypb.Empty, error) {
	_, responseCode, err := k.KsebUseCase.CheckIfComplaintBelongsToUser(req.UserId, req.ComplaintId)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	} else {
		return nil, nil
	}
}
