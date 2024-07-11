package main

import (
	"log"
	"net"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/config"
	db "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/infrastructure/db"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/server"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	ksebpb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated/ksebpb"

	"google.golang.org/grpc"
)

func main() {
	db.SeedDataToDbIfNotInitialised()
	// if err!=nil{
	// 	log.Fatal("database seed")
	// }
	lis, err := net.Listen("tcp", config.EnvValues.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Account Service >>>>>> Listening on port: ", config.EnvValues.Port)
	}

	userAccSvcServer, adminAccSvcServer, adminAppointmentsSvcServer, ksebAdminServer,ksebUserServer := server.InitializeServer()
	grpcServer := grpc.NewServer()
	pb.RegisterUserAccountServiceServer(grpcServer, userAccSvcServer)
	pb.RegisterAdminAccountServiceServer(grpcServer, adminAccSvcServer)
	pb.RegisterAppointmentServiceServer(grpcServer, adminAppointmentsSvcServer)
	ksebpb.RegisterKSEBAdminServiceServer(grpcServer, ksebAdminServer)
	ksebpb.RegisterKSEBUserServiceServer(grpcServer, ksebUserServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("failed to serve", err)
	}
}
