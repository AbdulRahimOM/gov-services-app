package adminAccHandler

import (
	ucinterface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
)

type AdminAccountsServer struct {
	AdminUseCase ucinterface.IAdminUC
	pb.UnimplementedAdminAccountServiceServer
}

func NewAdminAccountsServer(adminUseCase ucinterface.IAdminUC) *AdminAccountsServer {
	return &AdminAccountsServer{
		AdminUseCase: adminUseCase,
	}
}
