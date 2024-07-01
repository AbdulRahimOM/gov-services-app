package userAccHandler

import (
	ucinterface "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"
)

type UserAccountsServer struct {
	UserUseCase ucinterface.IUserUC
	pb.UnimplementedUserAccountServiceServer
}

func NewUserAccountsServer(userUseCase ucinterface.IUserUC) *UserAccountsServer {
	return &UserAccountsServer{
		UserUseCase: userUseCase,
	}
}
