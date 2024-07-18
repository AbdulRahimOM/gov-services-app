package ksebHandler

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/domain/models"
	ucinterface "github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/usecase/interface"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	"gorm.io/gorm"
)

const msgChannSize = 100

type KsebChatServer struct {
	ChatUseCase ucinterface.IKsebChatUC
	pb.UnimplementedKsebChatServiceServer
	UserChatStreams  map[int32]chan *pb.ChatMessage
	AdminChatStreams map[int32]chan *pb.ChatMessage
	mutex            sync.Mutex
	db               *gorm.DB
}

func NewKsebChatServer(chatUseCase ucinterface.IKsebChatUC,db *gorm.DB) *KsebChatServer {
	return &KsebChatServer{
		ChatUseCase:      chatUseCase,
		UserChatStreams:  make(map[int32]chan *pb.ChatMessage),
		AdminChatStreams: make(map[int32]chan *pb.ChatMessage),
		db: 			 db,
	}
}
func (s *KsebChatServer) UserChat(req *pb.UserChatRequest, stream pb.KsebChatService_UserChatServer) error {
	complaintId := req.GetComplaintId()

	s.mutex.Lock()
	if _, ok := s.AdminChatStreams[complaintId]; !ok {
		msgChan := make(chan *pb.ChatMessage, msgChannSize)
		defer close(msgChan)
		s.AdminChatStreams[complaintId] = msgChan
	}
	if _, ok := s.UserChatStreams[complaintId]; !ok {
		msgChan := make(chan *pb.ChatMessage, msgChannSize)
		defer close(msgChan)
		s.UserChatStreams[complaintId] = msgChan
	}
	s.mutex.Unlock()

	defer func() {
		s.mutex.Lock()
		delete(s.UserChatStreams, complaintId)
		delete(s.AdminChatStreams, complaintId)
		s.mutex.Unlock()
	}()

	for msg := range s.AdminChatStreams[complaintId] {
		if err := stream.Send(msg); err != nil {
			log.Printf("Error sending message to user: %v", err)
			return err
		}
	}
	return nil
}
func (s *KsebChatServer) UserSendMessage(ctx context.Context, req *pb.UserSendMessageRequest) (*pb.SendMessageResponse, error) {
	userID := req.GetUserId()
	complaintId := req.GetComplaintId()
	message := req.GetMessage()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	msg := &pb.ChatMessage{
		Sender:  fmt.Sprintf("User %d", userID),
		Message: message,
	}

	s.UserChatStreams[complaintId] <- msg

	msgModel := &models.ChatMessage{
		ComplaintId: complaintId,
		SenderId:    userID,
		SenderType:  "user",
		Content:     message,
		CreatedAt:   time.Now(),
	}
	s.SaveChatToDB(msgModel)

	return &pb.SendMessageResponse{Success: true}, nil
}

func (s *KsebChatServer) AdminChat(req *pb.AdminChatRequest, stream pb.KsebChatService_AdminChatServer) error {
	complaintId := req.GetComplaintId()

	s.mutex.Lock()
	if _, ok := s.AdminChatStreams[complaintId]; !ok {
		msgChan := make(chan *pb.ChatMessage, msgChannSize)
		defer close(msgChan)
		s.AdminChatStreams[complaintId] = msgChan

	}
	if _, ok := s.UserChatStreams[complaintId]; !ok {
		msgChan := make(chan *pb.ChatMessage, msgChannSize)
		defer close(msgChan)
		s.UserChatStreams[complaintId] = msgChan
	}
	s.mutex.Unlock()

	defer func() {
		s.mutex.Lock()
		delete(s.UserChatStreams, complaintId)
		delete(s.AdminChatStreams, complaintId)
		s.mutex.Unlock()
	}()

	for msg := range s.UserChatStreams[complaintId] {
		if err := stream.Send(msg); err != nil {
			log.Printf("Error sending message to admin: %v", err)
			return err
		}
	}
	return nil
}

func (s *KsebChatServer) AdminSendMessage(ctx context.Context, req *pb.AdminSendMessageRequest) (*pb.SendMessageResponse, error) {
	fmt.Println("req: ", req)
	adminID := req.GetAdminId()
	complaintId := req.GetComplaintId()
	message := req.GetMessage()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	msg := &pb.ChatMessage{
		Sender:  fmt.Sprintf("Admin %d", adminID),
		Message: message,
	}

	s.AdminChatStreams[complaintId] <- msg

	msgModel := &models.ChatMessage{
		ComplaintId: complaintId,
		SenderId:    adminID,
		SenderType:  "admin",
		Content:     message,
		CreatedAt:   time.Now(),
	}
	s.SaveChatToDB(msgModel)

	return &pb.SendMessageResponse{Success: true}, nil
}

func (s *KsebChatServer) SaveChatToDB(entry *models.ChatMessage) error {
	result := s.db.Create(entry)
	if result.Error != nil {
		return result.Error
	}
	return nil
}