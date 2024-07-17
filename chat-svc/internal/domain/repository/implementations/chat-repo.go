package chatrepo

import (
	repointerface "github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/domain/repository/interface"
	"gorm.io/gorm"
)

type ChatRepository struct {
	DB *gorm.DB
}

func NewChatRepository(db *gorm.DB) repointerface.IChatRepo {
	return &ChatRepository{DB: db}
}
