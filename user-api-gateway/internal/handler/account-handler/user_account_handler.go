package acchandler

import (
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	"github.com/eapache/go-resiliency/breaker"
)

type UserAccountHandler struct {
	accountsClient pb.UserAccountServiceClient
	circuitBreaker *breaker.Breaker
}

func NewUserAccountHandler(client pb.UserAccountServiceClient,circuitBreaker *breaker.Breaker) *UserAccountHandler {
	return &UserAccountHandler{
		accountsClient: client,
		circuitBreaker: circuitBreaker,
	}
}
