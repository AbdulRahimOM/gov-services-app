package ksebAdminHandler

import (
	"context"
	"sync"

	ucinterface "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/usecase/interface"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"google.golang.org/protobuf/types/known/emptypb"
)

type KSEBAgencyAdminServer struct {
	KsebUseCase ucinterface.IKsebAgencyAdminUC
	pb.UnimplementedKSEBAgencyAdminServiceServer
	messages map[int32][]*pb.ChatResponse
	mu       sync.Mutex
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