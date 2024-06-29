package handler

import (
	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"
)

type UserAccountHandler struct {
	accountsClient pb.UserAccountServiceClient
}

func NewUserAccountHandler(client pb.UserAccountServiceClient) *UserAccountHandler {
	return &UserAccountHandler{
		accountsClient: client,
	}
}

type AccountEntryHandler struct {
	accountsClient pb.UserAccountServiceClient
}

func NewEntryHandler(client pb.UserAccountServiceClient) *AccountEntryHandler {
	return &AccountEntryHandler{
		accountsClient: client,
	}
}
