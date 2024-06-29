package server

import (
	repo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/infrastructure/db"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase"
)

func InitializeServer() *AccountsServer {
	userRepository := repo.NewUserRepository(db.DB)
	userUseCase := usecase.NewUserUseCase(userRepository)

	return newAccountsServer(userUseCase)
}

func newAccountsServer(userUseCase usecase.IUserUC) *AccountsServer {
	return &AccountsServer{
		UserUseCase: userUseCase,
	}
}
