package ksebhanlder

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/segmentio/kafka-go"
)

// AdminChat
func (k *KSEBHandler) AdminChat(c *gin.Context) {
	complaintId, ok := gateway.HandleGetUrlParamsInt32(c, "complaintId")
	if !ok {
		return
	}

	adminID, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	_, err := k.agencyAdminClient.CheckIfComplaintAccessibleToAdmin(c, &pb.CheckIfComplaintAccessibleToAdminRequest{
		AdminId:     adminID,
		ComplaintId: complaintId,
	})
	if err != nil {
		gateway.HandleGrpcStatus(c, err)
		return
	}

	handleWebSocket(c.Writer, c.Request, k.ksebChatClient, adminID, complaintId)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request, chatClient pb.KsebChatServiceClient, adminID int32, complaintId int32) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	stream, err := chatClient.AdminChat(context.Background(), &pb.AdminChatRequest{
		AdminId:     adminID,
		ComplaintId: complaintId,
	})
	if err == nil {
		// Handle incoming messages from WebSocket clients
		// Read messages from WebSocket client and send to gRPC server
		go func() {
			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					log.Printf("Read error: %v", err)
					break
				}

				_, err = chatClient.AdminSendMessage(context.Background(), &pb.AdminSendMessageRequest{
					AdminId:     adminID,
					ComplaintId: complaintId,
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
		go websocketToKafka(conn, adminID, complaintId)
	}
	// Handle incoming messages from Kafka
	kafkaReader(context.Background(), conn)
}
func websocketToKafka(conn *websocket.Conn, adminID int32, complaintId int32) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "admin-messages",
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
			SenderId: adminID,
			SenderType: "admin",
			Content:  string(message),
		}
		msgBytes, err := json.Marshal(chatMessage)
		if err != nil {
			log.Printf("Failed to marshal chat message: %v", err)
			continue
		}

		err = writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("%d", adminID)),
				Value: msgBytes,
			},
		)
		if err != nil {
			log.Printf("Failed to send message to Kafka: %v", err)
			continue
		}
	}
}

func grpcReader(stream pb.KsebChatService_AdminChatClient, conn *websocket.Conn) {
	for {
		// Receive messages from gRPC stream
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Receive error from gRPC: %v", err)
			break
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s: %s", msg.Sender, msg.Message))); err != nil {
			log.Printf("Write error to WebSocket: %v", err)
			// continue
		}
	}
}
func kafkaReader(ctx context.Context, conn *websocket.Conn) { // Read messages from Kafka and send to WebSocket client
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "user-messages",
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
