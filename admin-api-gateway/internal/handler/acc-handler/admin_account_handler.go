package acchandler

import (
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	"github.com/eapache/go-resiliency/breaker"
)

type AdminAccountHandler struct {
	accountsClient pb.AdminAccountServiceClient
	circuitBreaker *breaker.Breaker
}

func NewAdminAccountHandler(client pb.AdminAccountServiceClient, circuitBreaker *breaker.Breaker) *AdminAccountHandler {
	return &AdminAccountHandler{
		accountsClient: client,
		circuitBreaker: circuitBreaker,
	}
}