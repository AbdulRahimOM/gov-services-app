package server

import (
	adminrepo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/implementations/admin-repo"
	ksebrepo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/implementations/kseb-repo"
	userrepo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/implementations/user-repo"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/infrastructure/db"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/admin/accHandler"
	appointmentsHandler "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/admin/appointmentsHandler"
	ksebHandler "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/admin/kseb-handler"
	ksebUserHandler "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/user/kseb-handler"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/userAccHandler"
	adminuc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/admin-uc/admin-account"
	appointmentsuc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/admin-uc/appointments"
	ksebuc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/admin-uc/kseb"
	userAccUc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/user-uc/account"
	ksebUserUc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/user-uc/kseb"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	ksebpb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated/ksebpb"
)

func InitializeServer() (
	pb.UserAccountServiceServer,
	pb.AdminAccountServiceServer,
	pb.AppointmentServiceServer,
	ksebpb.KSEBAdminServiceServer,
	ksebpb.KSEBUserServiceServer,
) {
	userRepository := userrepo.NewUserRepository(db.DB)
	userUseCase := userAccUc.NewUserUseCase(userRepository)
	userAccSvcServer := userAccHandler.NewUserAccountsServer(userUseCase)

	adminRepository := adminrepo.NewAdminRepository(db.DB)
	adminUseCase := adminuc.NewAdminUseCase(adminRepository)
	adminAccSvcServer := adminAccHandler.NewAdminAccountsServer(adminUseCase)

	appointmentsUseCase := appointmentsuc.NewAppointmentUseCase(adminRepository)
	appointmentsServer := appointmentsHandler.NewAppointmentServer(appointmentsUseCase)

	ksebRepository := ksebrepo.NewKsebRepository(db.DB)

	ksebAdminUseCase := ksebuc.NewKsebAdminUseCase(adminRepository, ksebRepository)
	ksebAdminServer := ksebHandler.NewKSEBAdminServer(ksebAdminUseCase)

	ksebUserUseCase := ksebUserUc.NewKsebUserUseCase(userRepository, ksebRepository)
	ksebUserServer := ksebUserHandler.NewKSEBUserServer(ksebUserUseCase)

	return userAccSvcServer, adminAccSvcServer, appointmentsServer, ksebAdminServer, ksebUserServer
}
