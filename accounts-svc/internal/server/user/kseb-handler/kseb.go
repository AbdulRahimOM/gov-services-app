package ksebUserHandler

import (
	ucinterface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	ksebpb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated/ksebpb"
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
