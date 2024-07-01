package main

import (
	"log"
	"net"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server"
	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", config.EnvValues.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Account Service >>>>>> Listening on port: ", config.EnvValues.Port)
	}

	

	userAccSvcServer,adminAccSvcServer := server.InitializeServer()
	grpcServer := grpc.NewServer()
	pb.RegisterUserAccountServiceServer(grpcServer, userAccSvcServer)
	pb.RegisterAdminAccountServiceServer(grpcServer, adminAccSvcServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("failed to serve", err)
	}
}