package ksebUserHandler

import (
	ucinterface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
)

type KSEBUserServer struct {
	KsebUseCase ucinterface.IKsebUserUC
	pb.UnimplementedKSEBUserAccServiceServer
}

func NewKSEBUserServer(ksebUseCase ucinterface.IKsebUserUC) *KSEBUserServer {
	return &KSEBUserServer{
		KsebUseCase: ksebUseCase,
	}
}
