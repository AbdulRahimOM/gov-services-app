package ksebhanlder


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
		AdminId:      adminID,
		ComplaintId: complaintId,
	})
	if err != nil {
		gateway.HandleGrpcStatus(c, err)
		return
	}

	handleWebSocket(c.Writer, c.Request, k.ksebChatClient, adminID,complaintId)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request, chatClient pb.KsebChatServiceClient, adminID int32, complaintId int32) {
	fmt.Println("in handleWebSocket")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()
	defer fmt.Println("Connection closed")
	fmt.Println("++++1")
	stream, err := chatClient.AdminChat(context.Background(), &pb.AdminChatRequest{
		AdminId: adminID,
		ComplaintId: complaintId,
	})
	if err != nil {
		log.Printf("Failed to create gRPC stream: %v", err)
		return
	}
	fmt.Println("~~~~")

	// Read messages from WebSocket client and send to gRPC server
	go func() {
		for {
			fmt.Println("---.")
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Read error: %v", err)
				break
			}

			_, err = chatClient.AdminSendMessage(context.Background(), &pb.AdminSendMessageRequest{
				AdminId: adminID,
				ComplaintId: complaintId,
				Message: string(message),
			})
			if err != nil {
				log.Printf("Failed to send gRPC message: %v", err)
				break
			}
		}
	}()
	fmt.Println("????")
	// Receive messages from gRPC server and send to WebSocket client
	for {
		fmt.Println("===")
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Receive error: %v", err)
			break
		}
		fmt.Println("====2")

		if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s: %s", msg.Sender, msg.Message))); err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}
