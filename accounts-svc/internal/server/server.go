package server

import (
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/config"
	adminrepo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/implementations/admin-repo"
	userrepo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/implementations/user-repo"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/infrastructure/db"
	adminAccHandler "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/admin/accHandler"
	appointmentsHandler "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/admin/appointmentsHandler"
	ksebHandler "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/admin/kseb-handler"
	userAccHandler "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/user/accHandler"
	ksebUserHandler "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server/user/kseb-handler"
	adminuc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/admin-uc/admin-account"
	appointmentsuc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/admin-uc/appointments"
	ksebuc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/admin-uc/kseb"
	userAccUc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/user-uc/account"
	ksebUserUc "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/implementations/user-uc/kseb"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AgenciesClients struct {
	KSEBClient pb.KSEBAgencyAdminServiceClient
}

func InitAgenciesClients() (*AgenciesClients, error) {
	clientConn, err := grpc.NewClient(config.EnvValues.AgenciesSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &AgenciesClients{
		KSEBClient: pb.NewKSEBAgencyAdminServiceClient(clientConn),
	}, nil
}

func InitializeServer(agenciesClients AgenciesClients) (
	pb.UserAccountServiceServer,
	pb.AdminAccountServiceServer,
	pb.AppointmentServiceServer,
	pb.KSEBAdminAccServiceServer,
	pb.KSEBUserAccServiceServer,
) {
	userRepository := userrepo.NewUserRepository(db.DB)
	userUseCase := userAccUc.NewUserUseCase(userRepository)
	userAccSvcServer := userAccHandler.NewUserAccountsServer(userUseCase)

	adminRepository := adminrepo.NewAdminRepository(db.DB)
	adminUseCase := adminuc.NewAdminUseCase(adminRepository)
	adminAccSvcServer := adminAccHandler.NewAdminAccountsServer(adminUseCase)

	appointmentsUseCase := appointmentsuc.NewAppointmentUseCase(adminRepository)
	appointmentsServer := appointmentsHandler.NewAppointmentServer(appointmentsUseCase)

	ksebAdminUseCase := ksebuc.NewKsebAdminUseCase(adminRepository)
	ksebAdminAccServer := ksebHandler.NewKSEBAdminServer(ksebAdminUseCase, agenciesClients.KSEBClient)

	ksebUserUseCase := ksebUserUc.NewKsebUserUseCase(userRepository)
	ksebUserServer := ksebUserHandler.NewKSEBUserServer(ksebUserUseCase)

	return userAccSvcServer, adminAccSvcServer, appointmentsServer, ksebAdminAccServer, ksebUserServer
}
