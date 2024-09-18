package ksebAdminHandler

import (
	"context"

	ucinterface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

type KSEBAdminServer struct {
	KsebUseCase ucinterface.IKsebAdminUC
	pb.UnimplementedKSEBAdminAccServiceServer
	KSEBAgencyClient pb.KSEBAgencyAdminServiceClient
	getGrpcStatus    func(respCode string, errMsg string) error
}

func NewKSEBAdminServer(
	ksebUseCase ucinterface.IKsebAdminUC,
	ksebAgencySvcClient pb.KSEBAgencyAdminServiceClient,
	logger *logrus.Entry,
) *KSEBAdminServer {
	return &KSEBAdminServer{
		KsebUseCase:      ksebUseCase,
		KSEBAgencyClient: ksebAgencySvcClient,
		getGrpcStatus:    stdresponse.NewGetGrpcStatusForService("accounts-svc", logger),
	}
}

func (k *KSEBAdminServer) RegisterSectionCode(ctx context.Context, req *pb.RegisterSectionCodeRequest) (*emptypb.Empty, error) {
	regSectionCodeReq := requests.KsebRegSectionCode{
		SectionCode: req.SectionCode,
		OfficeId:    req.OfficeId,
	}
	_, responseCode, err := k.KsebUseCase.CheckIfAdminCanRegisterSectionCode(req.AdminId, &regSectionCodeReq)
	if err != nil {
		return nil, stdresponse.GetGrpcStatus(responseCode, err.Error())
	}

	return k.KSEBAgencyClient.RegisterSectionCode(ctx, req)
}
