package ksebUserHandler

import (
	"context"

	ucinterface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
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
	responseCode,err := k.KsebUseCase.AddConsumerNumber(req.UserId, req.ConsumerNumber)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}else{
		return nil,nil
	}
}
