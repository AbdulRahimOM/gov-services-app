package ksebhandler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/AbdulRahimOM/gov-services-app/internal/gateway"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// UserChat
func (k *KsebHandler) UserChat(c *gin.Context) {
	complaintId, ok := gateway.HandleGetUrlParamsInt32(c, "complaintId")
	if !ok {
		return
	}

	userID, ok := gateway.GetUserIdFromContext(c)
	if !ok {
		return
	}

	_, err := k.agencyUserClient.CheckIfComplaintBelongsToUser(c, &pb.CheckIfComplaintBelongsToUserRequest{
		UserId:      userID,
		ComplaintId: complaintId,
	})
	if err != nil {
		gateway.HandleGrpcStatus(c, err)
		return
	}

	handleWebSocket(c.Writer, c.Request, k.ksebChatClient, userID,complaintId)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request, chatClient pb.KsebChatServiceClient, userID int32, complaintId int32) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	stream, err := chatClient.UserChat(context.Background(), &pb.UserChatRequest{
		UserId: userID,
		ComplaintId: complaintId,
	})
	if err != nil {
		log.Printf("Failed to create gRPC stream: %v", err)
		return
	}

	// Read messages from WebSocket client and send to gRPC server
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Read error: %v", err)
				break
			}

			_, err = chatClient.UserSendMessage(context.Background(), &pb.UserSendMessageRequest{
				UserId: userID,
				ComplaintId: complaintId,
				Message: string(message),
			})
			if err != nil {
				log.Printf("Failed to send gRPC message: %v", err)
				break
			}
		}
	}()
	
	// Receive messages from gRPC server and send to WebSocket client
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Receive error: %v", err)
			break
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s: %s", msg.Sender, msg.Message))); err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}
