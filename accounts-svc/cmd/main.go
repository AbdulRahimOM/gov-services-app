package main

import (
	"log"
	"net"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/config"
	db "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/infrastructure/db"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"

	"google.golang.org/grpc"
)

func main() {
	db.SeedDataToDbIfNotInitialised()
	
	lis, err := net.Listen("tcp", config.EnvValues.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Account Service >>>>>> Listening on port: ", config.EnvValues.Port)
	}

	serviceClients, err := server.InitAgenciesClients()
	if err != nil {
		log.Fatal("error occured while initializing service clients, error:", err)
	}

	userAccSvcServer, adminAccSvcServer, adminAppointmentsSvcServer, ksebAdminAccServer, ksebUserServer := server.InitializeServer(*serviceClients)
	grpcServer := grpc.NewServer()
	pb.RegisterUserAccountServiceServer(grpcServer, userAccSvcServer)
	pb.RegisterAdminAccountServiceServer(grpcServer, adminAccSvcServer)
	pb.RegisterAppointmentServiceServer(grpcServer, adminAppointmentsSvcServer)
	pb.RegisterKSEBAdminAccServiceServer(grpcServer, ksebAdminAccServer)
	pb.RegisterKSEBUserAccServiceServer(grpcServer, ksebUserServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("failed to serve", err)
	}
}
