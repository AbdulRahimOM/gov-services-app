package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/server"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	"github.com/segmentio/kafka-go"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", config.EnvValues.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Account Service >>>>>> Listening on port: ", config.EnvValues.Port)
	}

	go kafkaReader(context.Background())

	implementedServers := server.InitializeServer()
	grpcServer := grpc.NewServer()
	pb.RegisterKsebChatServiceServer(grpcServer, implementedServers.KsebChatServiceServer)

	// err = grpcServer.Serve(lis)
	// if err != nil {
	// 	log.Fatalln("failed to serve", err)
	// }

	// Handle server shutdown gracefully
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for termination signal
	waitForShutdown()

	// Gracefully shutdown
	grpcServer.GracefulStop()
	log.Println("Server stopped")
}

func kafkaReader(ctx context.Context) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{config.EnvValues.KafkaUrl},
		Topic:     "chat-messages",
		Partition: 0,
	})
	reader.SetOffset(kafka.FirstOffset)

	defer func() {
		if err := reader.Close(); err != nil {
			log.Printf("error closing Kafka reader: %s", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("Kafka reader stopped")
			return
		default:
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("could not read message %s", err)
				continue
			}

			var chatMessage commondto.ChatMessage
			if err := json.Unmarshal(msg.Value, &chatMessage); err != nil {
				log.Printf("could not unmarshal message %s", err)
				continue
			}

			log.Printf("received: %+v", chatMessage)
			// Process the message here (e.g., save to database, send to WebSocket, etc.)
		}
	}
}
func waitForShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Println("Received termination signal. Shutting down...")
}

// func kafkaReader() {
// 	reader := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers:   []string{config.EnvValues.KafkaUrl},
// 		Topic:     "chat-messages",
// 		Partition: 0,
// 	})
// 	reader.SetOffset(kafka.FirstOffset)

// 	for {
// 		msg, err := reader.ReadMessage(context.Background())
// 		if err != nil {
// 			log.Printf("could not read message %s", err)
// 			continue
// 		}

// 		var chatMessage commondto.ChatMessage
// 		err = json.Unmarshal(msg.Value, &chatMessage)
// 		if err != nil {
// 			log.Printf("could not unmarshal message %s", err)
// 			continue
// 		}

// 		log.Printf("received: %+v", chatMessage)
// 		// Process the message here (e.g., save to database, send to WebSocket, etc.)
// 	}
// }
