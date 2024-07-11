package ksebHandler

import (
	"context"

	ucinterface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	ksebpb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated/ksebpb"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"google.golang.org/protobuf/types/known/emptypb"
)

type KsebServer struct {
	KsebUseCase ucinterface.IKsebUC
	ksebpb.UnimplementedKSEBServiceServer
}

func NewKsebServer(ksebUseCase ucinterface.IKsebUC) *KsebServer {
	return &KsebServer{
		KsebUseCase: ksebUseCase,
	}
}

func (k *KsebServer) RegisterSectionCode(ctx context.Context, req *ksebpb.RegisterSectionCodeRequest) (*emptypb.Empty, error) {
	regSectionCodeReq := requests.KsebRegSectionCode{
		SectionCode: req.SectionCode,
		OfficeId:    req.OfficeId,
	}
	_, responseCode, err := k.KsebUseCase.RegisterSectionCode(req.AdminId, &regSectionCodeReq)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}else{
		return nil, nil
	}
}
