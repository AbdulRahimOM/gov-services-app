package ksebhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	gateway "github.com/AbdulRahimOM/gov-services-app/internal/gateway/fiber"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/config"
	"github.com/gofiber/websocket/v2"
	"github.com/segmentio/kafka-go"
)

func (k *KsebHandler) UserChatWebsocket(conn *websocket.Conn) {
	defer conn.Close()
	complaintId := conn.Params("complaintId")

	userID, ok := gateway.GetUserIdFromWebsocketConn(conn)
	if !ok {
		conn.WriteMessage(websocket.TextMessage, []byte("Unauthorized"))
		return
	}
	fmt.Println("userID", userID)

	// Convert complaintId to int32
	complaintIdInt, err := strconv.Atoi(complaintId)
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("Invalid complaintId"))
		return
	}
	complaintIdInt32 := int32(complaintIdInt)

	_, err = k.agencyUserClient.CheckIfComplaintBelongsToUser(context.Background(), &pb.CheckIfComplaintBelongsToUserRequest{
		UserId:      userID,
		ComplaintId: complaintIdInt32,
	})
	if err != nil {
		log.Printf("gRPC error: %v", err)
		conn.WriteMessage(websocket.TextMessage, []byte("Error: Could not verify complaint ownership"))
		return
	}

	stream, err := k.ksebChatClient.UserChat(context.Background(), &pb.UserChatRequest{
		UserId:      userID,
		ComplaintId: complaintIdInt32,
	})
	if err == nil {
		// Handle incoming messages from WebSocket clients
		go func() {
			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					log.Printf("Read error: %v", err)
					break
				}

				_, err = k.ksebChatClient.UserSendMessage(context.Background(), &pb.UserSendMessageRequest{
					UserId:      userID,
					ComplaintId: complaintIdInt32,
					Message:     string(message),
				})
				if err != nil {
					log.Printf("Failed to send gRPC message: %v", err)
					break
				}
			}
		}()

		// Handle incoming messages from gRPC server
		go grpcReader(stream, conn)
	} else {
		log.Printf("Failed to create gRPC stream: %v", err)

		// handler incoming messages from WebSocket clients and send to Kafka (instead of gRPC)
		go websocketToKafka(conn, userID, complaintIdInt32)
	}

	// Handle incoming messages from Kafka
	kafkaReader(context.Background(), conn)
}

func grpcReader(stream pb.KsebChatService_UserChatClient, conn *websocket.Conn) {
	for {
		// Receive messages from gRPC stream
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Receive error from gRPC: %v", err)
			break
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s: %s", msg.Sender, msg.Message))); err != nil {
			log.Printf("Write error to WebSocket: %v", err)
			continue
		}
	}
}

func websocketToKafka(conn *websocket.Conn, userID int32, complaintId int32) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{config.EnvValues.KafkaUrl},
		Topic:   "user-messages",
	})

	defer func() {
		if err := writer.Close(); err != nil {
			log.Printf("error closing Kafka writer: %s", err)
		}
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		chatMessage := commondto.ChatMessage{
			ComplaintId: complaintId,
			SenderId:    userID,
			SenderType:  "user",
			Content:     string(message),
		}
		msgBytes, err := json.Marshal(chatMessage)
		if err != nil {
			log.Printf("Failed to marshal chat message: %v", err)
			continue
		}

		err = writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("%d", userID)),
				Value: msgBytes,
			},
		)
		if err != nil {
			log.Printf("Failed to send message to Kafka: %v", err)
			continue
		}
	}
}

func kafkaReader(ctx context.Context, conn *websocket.Conn) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{config.EnvValues.KafkaUrl},
		Topic:     "admin-messages",
		Partition: 0,
	})
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

			log.Printf("received from Kafka: %+v", chatMessage)

			// Send Kafka message to WebSocket client
			if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d: %s", chatMessage.SenderId, chatMessage.Content))); err != nil {
				log.Printf("Write error to WebSocket: %v", err)
				continue
			}
		}
	}
}
