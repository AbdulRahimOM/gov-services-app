package server

import (
	chatrepo "github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/domain/repository/implementations"
	"github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/infrastructure/db"
	ksebHandler "github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/server/chat-handler/kseb"
	ksebUc "github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/usecase/implementations/kseb"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
)

type Servers struct {
	KsebChatServiceServer pb.KsebChatServiceServer
}

func InitializeServer() *Servers {

	chatRepository := chatrepo.NewChatRepository(db.DB)

	KsebChatUseCase := ksebUc.NewKsebChatUseCase(chatRepository)
	KsebChatServer := ksebHandler.NewKsebChatServer(KsebChatUseCase,db.DB)


	return &Servers{
		KsebChatServiceServer: KsebChatServer,
	}
}
