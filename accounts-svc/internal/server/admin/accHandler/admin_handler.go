package adminAccHandler

import (
	ucinterface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"github.com/sirupsen/logrus"
)

type AdminAccountsServer struct {
	AdminUseCase ucinterface.IAdminUC
	pb.UnimplementedAdminAccountServiceServer
	getGrpcStatus func(respCode string, errMsg string) error
}

func NewAdminAccountsServer(adminUseCase ucinterface.IAdminUC, logger *logrus.Entry) *AdminAccountsServer {
	return &AdminAccountsServer{
		AdminUseCase:  adminUseCase,
		getGrpcStatus: stdresponse.NewGetGrpcStatusForService("accounts-svc", logger),
	}
}
