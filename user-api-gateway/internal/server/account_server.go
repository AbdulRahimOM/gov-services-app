package server

import (
	ksebhandler "github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/handler/kseb-handler"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	"github.com/AbdulRahimOM/gov-services-app/internal/pb/generated/ksebpb"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/config"
	acchandler "github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/handler/account-handler"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/routes"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClients struct {
	UserAccountsClient pb.UserAccountServiceClient
	KsebClient         ksebpb.KSEBUserServiceClient
}

func InitServiceClients() (*ServiceClients, error) {
	clientConn, err := grpc.NewClient(config.EnvValues.AccountsSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &ServiceClients{
		UserAccountsClient: pb.NewUserAccountServiceClient(clientConn),
		KsebClient:         ksebpb.NewKSEBUserServiceClient(clientConn),
	}, nil
}

func InitRoutes(serviceClients *ServiceClients, engine *gin.Engine) {
	accountHandler := acchandler.NewUserAccountHandler(serviceClients.UserAccountsClient)
	ksebHandler := ksebhandler.NewKsebHandler(serviceClients.KsebClient)

	routes.RegisterRoutes(engine.Group("/"), accountHandler)
	routes.RegisterKsebRoutes(engine.Group("/kseb"), ksebHandler)
}
