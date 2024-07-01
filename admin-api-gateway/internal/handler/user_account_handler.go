package handler

import (
	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"
)

type AdminAccountHandler struct {
	accountsClient pb.AdminAccountServiceClient
}

func NewAdminAccountHandler(client pb.AdminAccountServiceClient) *AdminAccountHandler {
	return &AdminAccountHandler{
		accountsClient: client,
	}
}

type AccountEntryHandler struct {
	accountsClient pb.AdminAccountServiceClient
}

func NewEntryHandler(client pb.AdminAccountServiceClient) *AccountEntryHandler {
	return &AccountEntryHandler{
		accountsClient: client,
	}
}
