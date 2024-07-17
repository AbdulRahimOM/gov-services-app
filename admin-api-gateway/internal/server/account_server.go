package server

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/config"
	acchandler "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/acc-handler"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/appointments"
	ksebhanlder "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/kseb"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/routes"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClients struct {
	AccountsClient     pb.AdminAccountServiceClient
	AppointmentsClient pb.AppointmentServiceClient
	KsebAccClient      pb.KSEBAdminAccServiceClient

	KSEBAgencyAdminClient pb.KSEBAgencyAdminServiceClient
	KSEBAgencyUserClient pb.KSEBAgencyUserServiceClient
}

func InitServiceClients() (*ServiceClients, error) {
	accountsSvcClientConn, err := grpc.NewClient(config.EnvValues.AccountsSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	agenciesSvcClientConn, err := grpc.NewClient(config.EnvValues.AgenciesSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &ServiceClients{
		AccountsClient:     pb.NewAdminAccountServiceClient(accountsSvcClientConn),
		AppointmentsClient: pb.NewAppointmentServiceClient(accountsSvcClientConn),
		KsebAccClient:      pb.NewKSEBAdminAccServiceClient(accountsSvcClientConn),

		KSEBAgencyAdminClient: pb.NewKSEBAgencyAdminServiceClient(agenciesSvcClientConn),
		KSEBAgencyUserClient:  pb.NewKSEBAgencyUserServiceClient(agenciesSvcClientConn),
	}, nil
}

func InitRoutes(serviceClients *ServiceClients, engine *gin.Engine) {
	accountHandler := acchandler.NewAdminAccountHandler(serviceClients.AccountsClient)
	appointmentHandler := appointments.NewAppointmentHandler(serviceClients.AppointmentsClient)

	ksebAccHandler := ksebhanlder.NewKsebHandler(serviceClients.KsebAccClient, serviceClients.KSEBAgencyAdminClient)

	routes.RegisterRoutes(engine.Group("/"), accountHandler, appointmentHandler)
	routes.RegisterKSEBAccRoutes(engine.Group("/kseb"), ksebAccHandler)
}
