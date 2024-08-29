package server

import (
	"time"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/config"
	acchandler "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/acc-handler"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/appointments"
	ksebhanlder "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/kseb"
	w "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/webrtc"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/routes"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	"github.com/gofiber/fiber/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClients struct {
	// accounts-svc clients
	AccountsClient     pb.AdminAccountServiceClient
	AppointmentsClient pb.AppointmentServiceClient
	KsebAccClient      pb.KSEBAdminAccServiceClient

	// agencies-svc clients
	KSEBAgencyAdminClient pb.KSEBAgencyAdminServiceClient
	KSEBAgencyUserClient  pb.KSEBAgencyUserServiceClient

	// chat-svc clients
	KsebChatClient pb.KsebChatServiceClient
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

	chatSvcClientConn, err := grpc.NewClient(config.EnvValues.ChatSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &ServiceClients{
		AccountsClient:     pb.NewAdminAccountServiceClient(accountsSvcClientConn),
		AppointmentsClient: pb.NewAppointmentServiceClient(accountsSvcClientConn),
		KsebAccClient:      pb.NewKSEBAdminAccServiceClient(accountsSvcClientConn),

		KSEBAgencyAdminClient: pb.NewKSEBAgencyAdminServiceClient(agenciesSvcClientConn),
		KSEBAgencyUserClient:  pb.NewKSEBAgencyUserServiceClient(agenciesSvcClientConn),

		KsebChatClient: pb.NewKsebChatServiceClient(chatSvcClientConn),
	}, nil
}

func InitRoutes(serviceClients *ServiceClients, api *fiber.App) {
	accountHandler := acchandler.NewAdminAccountHandler(serviceClients.AccountsClient)
	appointmentHandler := appointments.NewAppointmentHandler(serviceClients.AppointmentsClient)

	ksebAccHandler := ksebhanlder.NewKsebHandler(serviceClients.KsebAccClient, serviceClients.KSEBAgencyAdminClient, serviceClients.KsebChatClient)

	routes.RegisterRoutes(api.Group("/"), accountHandler, appointmentHandler)
	routes.RegisterKSEBAccRoutes(api.Group("/kseb"), ksebAccHandler)
	w.Rooms = make(map[int32]*w.Room)
	go dispatchKeyFrames()
}

func dispatchKeyFrames() {
	for range time.NewTicker(time.Second * 3).C {
		for _, room := range w.Rooms {
			room.Peers.DispatchKeyFrame()
		}
	}
}
