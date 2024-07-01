package adminAccHandler

import (
	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"
	ucinterface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
)

type AdminAccountsServer struct {
	AdminUseCase ucinterface.IAdminUC
	pb.UnimplementedAdminAccountServiceServer
}

var _ pb.AdminAccountServiceServer

func NewAdminAccountsServer(adminUseCase ucinterface.IAdminUC) *AdminAccountsServer {
	return &AdminAccountsServer{
		AdminUseCase: adminUseCase,
	}
}
