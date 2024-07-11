package server

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/acc-handler"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/appointments"
	ksebhanlder "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/kseb"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/routes"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	ksebpb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated/ksebpb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClients struct {
	AccountsClient     pb.AdminAccountServiceClient
	AppointmentsClient pb.AppointmentServiceClient
	KsebClient         ksebpb.KSEBAdminServiceClient
}

func InitServiceClients() (*ServiceClients, error) {
	clientConn, err := grpc.NewClient(config.EnvValues.AccountsSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &ServiceClients{
		AccountsClient:     pb.NewAdminAccountServiceClient(clientConn),
		AppointmentsClient: pb.NewAppointmentServiceClient(clientConn),
		KsebClient:         ksebpb.NewKSEBAdminServiceClient(clientConn),
	}, nil
}

func InitRoutes(serviceClients *ServiceClients, engine *gin.Engine) {
	accountHandler := acchandler.NewAdminAccountHandler(serviceClients.AccountsClient)
	appointmentHandler := appointments.NewAppointmentHandler(serviceClients.AppointmentsClient)

	ksebHandler := ksebhanlder.NewAppointmentHandler(serviceClients.KsebClient)

	routes.RegisterRoutes(engine.Group("/"), accountHandler, appointmentHandler)
	routes.RegisterKSEBRoutes(engine.Group("/kseb"), ksebHandler)
}
