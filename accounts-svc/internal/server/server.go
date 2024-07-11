package server

import (
	adminrepo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/implementations/admin-repo"
	userrepo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/implementations/user-repo"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/infrastructure/db"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/adminAccHandler"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/userAccHandler"
	appointmentsHandler "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/appointmentsHandler"
	adminuc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/admin-uc/admin-account"
	useruc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/user-uc"
	appointmentsuc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/admin-uc/appointments"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
)

func InitializeServer() (
	pb.UserAccountServiceServer,
	pb.AdminAccountServiceServer,
	pb.AppointmentServiceServer,
) {
	userRepository := userrepo.NewUserRepository(db.DB)
	userUseCase := useruc.NewUserUseCase(userRepository)
	userAccSvcServer := userAccHandler.NewUserAccountsServer(userUseCase)

	adminRepository := adminrepo.NewAdminRepository(db.DB)
	adminUseCase := adminuc.NewAdminUseCase(adminRepository)
	adminAccSvcServer := adminAccHandler.NewAdminAccountsServer(adminUseCase)

	appointmentsUseCase := appointmentsuc.NewAppointmentUseCase(adminRepository)
	appointmentsServer := appointmentsHandler .NewAppointmentServer(appointmentsUseCase)

	ksebUseCase := ksebuc.NewKsebUseCase(adminRepository)
	ksebServer := ksebHandler.NewKsebServer(ksebUseCase)

	return userAccSvcServer, adminAccSvcServer, appointmentsServer
}
