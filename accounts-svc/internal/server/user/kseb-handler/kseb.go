package ksebUserHandler

import (
	ucinterface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"github.com/sirupsen/logrus"
)

type KSEBUserServer struct {
	KsebUseCase ucinterface.IKsebUserUC
	pb.UnimplementedKSEBUserAccServiceServer
	getGrpcStatus func(respCode string, errMsg string) error
}

func NewKSEBUserServer(ksebUseCase ucinterface.IKsebUserUC, logger *logrus.Entry) *KSEBUserServer {
	return &KSEBUserServer{
		KsebUseCase:   ksebUseCase,
		getGrpcStatus: stdresponse.NewGetGrpcStatusForService("accounts-svc", logger),
	}
}
