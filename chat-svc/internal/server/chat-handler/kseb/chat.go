package ksebHandler

import (
	"context"
	"fmt"
	"log"
	"sync"

	ucinterface "github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/usecase/interface"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
)

type KsebChatServer struct {
	ChatUseCase ucinterface.IKsebChatUC
	pb.UnimplementedKsebChatServiceServer
	KsebChatStreams map[int32]chan *pb.ChatMessage
	mutex           sync.Mutex
}

func NewKsebChatServer(chatUseCase ucinterface.IKsebChatUC) *KsebChatServer {
	return &KsebChatServer{
		ChatUseCase:     chatUseCase,
		KsebChatStreams: make(map[int32]chan *pb.ChatMessage),
	}
}
func (s *KsebChatServer) UserChat(req *pb.UserChatRequest, stream pb.KsebChatService_UserChatServer) error {
	userID := req.GetUserId()
	msgChan := make(chan *pb.ChatMessage, 100)

	s.mutex.Lock()
	s.KsebChatStreams[userID] = msgChan
	s.mutex.Unlock()

	defer func() {
		s.mutex.Lock()
		delete(s.KsebChatStreams, userID)
		s.mutex.Unlock()
		close(msgChan)
	}()

	// Receive messages from the channel and send them to the client
	for msg := range msgChan {
		fmt.Println("msg here: ", msg)
		if err := stream.Send(msg); err != nil {
			log.Printf("Error sending message to user %d: %v", userID, err)
			return err
		}
	}
	return nil
}
func (s *KsebChatServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	fmt.Println("req: ", req)
	userID := req.GetUserId()
	message := req.GetMessage()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	msg := &pb.ChatMessage{
		Sender:  fmt.Sprintf("User %d", userID),
		Message: message,
	}

	var sendTo int32
	switch userID {
	case 14:
		sendTo = 15
	case 15:
		sendTo = 14
	}
	// Send the message to all connected users
	for usrId, ch := range s.KsebChatStreams {
		if usrId == sendTo {
			ch <- msg
		}
	}

	// // Create a reply message with "hello" appended
	// replyMessage := &pb.ChatMessage{
	// 	Sender:  "Server",
	// 	Message: fmt.Sprintf("hello %s", message),
	// }

	// // Send the reply message back to the originating user
	// if ch, ok := s.KsebChatStreams[userID]; ok {
	// 	ch <- replyMessage
	// }

	return &pb.SendMessageResponse{Success: true}, nil
}
