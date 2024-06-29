package server

import (
	repo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/infrastructure/db"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase"
)

func InitializeServer() *UserAccountsServer {
	userRepository := repo.NewUserRepository(db.DB)
	userUseCase := usecase.NewUserUseCase(userRepository)

	return newUserAccountsServer(userUseCase)
}

func newUserAccountsServer(userUseCase usecase.IUserUC) *UserAccountsServer {
	return &UserAccountsServer{
		UserUseCase: userUseCase,
	}
}
