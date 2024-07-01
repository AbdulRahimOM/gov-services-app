package server

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/routes"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClients struct {
	AccountsClient pb.AdminAccountServiceClient
}

func InitServiceClients() (*ServiceClients, error) {
	clientConn, err := grpc.NewClient(config.EnvValues.AccountsSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &ServiceClients{
		AccountsClient: pb.NewAdminAccountServiceClient(clientConn),
	}, nil
}

func InitRoutes(serviceClients *ServiceClients, engine *gin.Engine) {
	entryHandler := handler.NewEntryHandler(serviceClients.AccountsClient)
	accountHandler := handler.NewAdminAccountHandler(serviceClients.AccountsClient)

	routes.RegisterRoutes(engine.Group("/"), entryHandler, accountHandler)
}
