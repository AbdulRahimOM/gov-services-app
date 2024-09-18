package server

import (
	"time"

	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/config"
	acchandler "github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/handler/account-handler"
	ksebhandler "github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/handler/kseb-handler"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/middleware"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/routes"
	w "github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/webrtc"
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

var (
	circuitBreaker *breaker.Breaker = breaker.New(errorThreshold, successThreshold, timeout)
)

type ServiceClients struct {
	//services in account-svc
	UserAccountsClient pb.UserAccountServiceClient
	//KsebAccClient         pb.KSEBUserAccServiceClient

	//services in agency-svc
	KSEBAgencyUserClient pb.KSEBAgencyUserServiceClient

	// services in chat-svc
	KsebChatClient pb.KsebChatServiceClient
}

func InitServiceClients() (*ServiceClients, error) {
	accountsSvcClientConn, err := grpc.NewClient(config.EnvValues.AccountsSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	agencySvcClientConn, err := grpc.NewClient(config.EnvValues.AgenciesSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	chatSvcClientConn, err := grpc.NewClient(config.EnvValues.ChatSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &ServiceClients{
		UserAccountsClient: pb.NewUserAccountServiceClient(accountsSvcClientConn),
		// KsebAccClient:         pb.NewKSEBAgencyUserServiceClient(accountsSvcClientConn),
		KSEBAgencyUserClient: pb.NewKSEBAgencyUserServiceClient(agencySvcClientConn),
		KsebChatClient:       pb.NewKsebChatServiceClient(chatSvcClientConn),
	}, nil
}

func InitRoutes(serviceClients *ServiceClients, api fiber.Router) {
	accountHandler := acchandler.NewUserAccountHandler(serviceClients.UserAccountsClient, circuitBreaker)
	// ksebAccHandler := ksebhandler.NewKsebHandler(serviceClients.KsebAccClient)

	ksebAgencyUserHandler := ksebhandler.NewKsebHandler(
		serviceClients.KSEBAgencyUserClient,
		serviceClients.KsebChatClient,
		circuitBreaker,
	)

	api.Use(middleware.ClearCache)
	routes.RegisterRoutes(api.Group("/"), accountHandler)
	routes.RegisterKsebRoutes(api.Group("/kseb"), ksebAgencyUserHandler)
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
