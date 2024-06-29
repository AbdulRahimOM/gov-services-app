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
		defer func() {
			lis.Close()
			log.Println("Listener closed")
		}()
	}

	server := server.InitializeServer()
	grpcServer := grpc.NewServer()
	defer func() {
		grpcServer.GracefulStop()
		log.Println("Server gracefully stopped")
	}()
	pb.RegisterUserAccountServiceServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("failed to serve", err)
	}
}

// func main() {
// 	lis, err := net.Listen("tcp", config.EnvValues.Port)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}else{
// 		log.Println("Listening on port: ", config.EnvValues.Port)
// 	}
// 	defer lis.Close()

// 	server := server.InitializeServer()
// 	grpcServer := grpc.NewServer()
// 	pb.RegisterAccountServiceServer(grpcServer, server)
// 	err = grpcServer.Serve(lis)
// 	if err != nil {
// 		log.Fatalln("failed to serve", err)
// 	}

// }
