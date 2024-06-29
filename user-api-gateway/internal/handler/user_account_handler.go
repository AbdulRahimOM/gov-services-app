package handler

import (
	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"
)

type UserAccountHandler struct {
	accountsClient pb.AccountServiceClient
}

func NewUserAccountHandler(client pb.AccountServiceClient) *UserAccountHandler {
	return &UserAccountHandler{
		accountsClient: client,
	}
}

type AccountEntryHandler struct {
	accountsClient pb.AccountServiceClient
}

func NewEntryHandler(client pb.AccountServiceClient) *AccountEntryHandler {
	return &AccountEntryHandler{
		accountsClient: client,
	}
}
