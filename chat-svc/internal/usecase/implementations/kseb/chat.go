package ksebChatUC

import (
	repo "github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/domain/repository/interface"
	usecase "github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/usecase/interface"
)

type KsebChatUseCase struct {
	chatRepo repo.IChatRepo
}

func NewKsebChatUseCase(chatRepo repo.IChatRepo) usecase.IKsebChatUC {
	return &KsebChatUseCase{
		chatRepo: chatRepo,
	}
}
