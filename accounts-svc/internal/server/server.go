package server

import (
	adminrepo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/implementations/admin-repo"
	ksebrepo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/implementations/kseb-repo"
	userrepo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/implementations/user-repo"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/infrastructure/db"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/adminAccHandler"
	appointmentsHandler "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/appointmentsHandler"
	ksebHandler "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/kseb-handler"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/userAccHandler"
	adminuc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/admin-uc/admin-account"
	appointmentsuc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/admin-uc/appointments"
	ksebuc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/admin-uc/kseb"
	useruc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/user-uc"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	ksebpb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated/ksebpb"
)

func InitializeServer() (
	pb.UserAccountServiceServer,
	pb.AdminAccountServiceServer,
	pb.AppointmentServiceServer,
	ksebpb.KSEBServiceServer,
) {
	userRepository := userrepo.NewUserRepository(db.DB)
	userUseCase := useruc.NewUserUseCase(userRepository)
	userAccSvcServer := userAccHandler.NewUserAccountsServer(userUseCase)

	adminRepository := adminrepo.NewAdminRepository(db.DB)
	adminUseCase := adminuc.NewAdminUseCase(adminRepository)
	adminAccSvcServer := adminAccHandler.NewAdminAccountsServer(adminUseCase)

	appointmentsUseCase := appointmentsuc.NewAppointmentUseCase(adminRepository)
	appointmentsServer := appointmentsHandler .NewAppointmentServer(appointmentsUseCase)

	ksebRepository:=ksebrepo.NewKsebRepository(db.DB)
	ksebUseCase := ksebuc.NewKsebUseCase(adminRepository,ksebRepository)
	ksebServer := ksebHandler.NewKsebServer(ksebUseCase)

	return userAccSvcServer, adminAccSvcServer, appointmentsServer, ksebServer
}
