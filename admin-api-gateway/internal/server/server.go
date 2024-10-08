package server

import (
	"time"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/config"
	acchandler "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/acc-handler"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/appointments"
	ksebhanlder "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/kseb"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/middleware"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/routes"
	w "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/webrtc"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	"github.com/eapache/go-resiliency/breaker"
	"github.com/gofiber/fiber/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	//for circuit breaker configuration
	errorThreshold   = 3
	successThreshold = 1
	timeout          = 5 * time.Second
)

var circuitBreaker *breaker.Breaker = breaker.New(errorThreshold, successThreshold, timeout)

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

func InitRoutes(serviceClients *ServiceClients, api fiber.Router) {
	accountHandler := acchandler.NewAdminAccountHandler(serviceClients.AccountsClient, circuitBreaker)
	appointmentHandler := appointments.NewAppointmentHandler(serviceClients.AppointmentsClient, circuitBreaker)

	ksebAccHandler := ksebhanlder.NewKsebHandler(
		serviceClients.KsebAccClient,
		serviceClients.KSEBAgencyAdminClient,
		serviceClients.KsebChatClient,
		circuitBreaker,
	)

	adminGroup := api.Group("/admin")
	adminGroup.Use(middleware.ClearCache)
	routes.RegisterRoutes(adminGroup, accountHandler, appointmentHandler)
	routes.RegisterKSEBAccRoutes(adminGroup, ksebAccHandler)
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
