package server

import (
	ksebrepo "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/domain/repository/implementations/kseb-repo"
	"github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/infrastructure/db"
	ksebAdminHandler "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/server/kseb-handler/admin"
	ksebUserHandler "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/server/kseb-handler/user"
	ksebuc "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/usecase/implementations/kseb/admin-uc"
	ksebUserUc "github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/usecase/implementations/kseb/user-uc"
	"github.com/AbdulRahimOM/gov-services-app/internal/logs"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
)

type Servers struct {
	KSEBAgencyAdminServiceServer pb.KSEBAgencyAdminServiceServer
	KSEBAgencyUserServiceServer  pb.KSEBAgencyUserServiceServer
}

var logger = logs.NewLoggerWithServiceName("agencies-svc")

func InitializeServer() *Servers {

	ksebRepository := ksebrepo.NewKsebRepository(db.DB)

	ksebAgencyAdminUseCase := ksebuc.NewKSEBAgencyAdminUseCase(ksebRepository)
	ksebAgencyAdminServer := ksebAdminHandler.NewKSEBAgencyAdminServer(ksebAgencyAdminUseCase, logger)

	ksebAgencyUserUseCase := ksebUserUc.NewKsebAgencyUserUseCase(ksebRepository)
	ksebAgencyUserServer := ksebUserHandler.NewKSEBAgencyUserServer(ksebAgencyUserUseCase, logger)

	return &Servers{
		KSEBAgencyAdminServiceServer: ksebAgencyAdminServer,
		KSEBAgencyUserServiceServer:  ksebAgencyUserServer,
	}
}
