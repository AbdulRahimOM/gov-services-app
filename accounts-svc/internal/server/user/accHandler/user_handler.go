package userAccHandler

import (
	ucinterface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"github.com/sirupsen/logrus"
)

type UserAccountsServer struct {
	UserUseCase ucinterface.IUserUC
	pb.UnimplementedUserAccountServiceServer
	getGrpcStatus func(respCode string, errMsg string) error
}

func NewUserAccountsServer(userUseCase ucinterface.IUserUC, logger *logrus.Entry) *UserAccountsServer {
	return &UserAccountsServer{
		UserUseCase:   userUseCase,
		getGrpcStatus: stdresponse.NewGetGrpcStatusForService("accounts-svc", logger),
	}
}
