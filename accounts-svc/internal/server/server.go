package server

import (
	adminrepo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/implementations/admin-repo"
	userrepo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/implementations/user-repo"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/infrastructure/db"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/adminAccHandler"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/userAccHandler"
	adminuc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/admin-uc"
	useruc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/user-uc"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
)

func InitializeServer() (pb.UserAccountServiceServer, pb.AdminAccountServiceServer) {
	userRepository := userrepo.NewUserRepository(db.DB)
	userUseCase := useruc.NewUserUseCase(userRepository)
	userAccSvcServer := userAccHandler.NewUserAccountsServer(userUseCase)

	adminRepository := adminrepo.NewAdminRepository(db.DB)
	adminUseCase := adminuc.NewAdminUseCase(adminRepository)
	adminAccSvcServer := adminAccHandler.NewAdminAccountsServer(adminUseCase)

	return userAccSvcServer, adminAccSvcServer
}
