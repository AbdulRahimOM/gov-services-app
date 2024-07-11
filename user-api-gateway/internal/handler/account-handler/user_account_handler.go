package acchandler

import pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"

type UserAccountHandler struct {
	accountsClient pb.UserAccountServiceClient
}

func NewUserAccountHandler(client pb.UserAccountServiceClient) *UserAccountHandler {
	return &UserAccountHandler{
		accountsClient: client,
	}
}
