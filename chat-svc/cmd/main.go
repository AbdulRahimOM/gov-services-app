package main

import (
	"log"
	"net"

	"github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/server"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", config.EnvValues.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Account Service >>>>>> Listening on port: ", config.EnvValues.Port)
	}

	implementedServers := server.InitializeServer()
	grpcServer := grpc.NewServer()
	pb.RegisterKsebChatServiceServer(grpcServer, implementedServers.KsebChatServiceServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("failed to serve", err)
	}
}
