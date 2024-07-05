package handler

import (
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
)

type AdminAccountHandler struct {
	accountsClient pb.AdminAccountServiceClient
}

func NewAdminAccountHandler(client pb.AdminAccountServiceClient) *AdminAccountHandler {
	return &AdminAccountHandler{
		accountsClient: client,
	}
}

type AccountHandler struct {
	accountsClient pb.AdminAccountServiceClient
}

func NewEntryHandler(client pb.AdminAccountServiceClient) *AccountHandler {
	return &AccountHandler{
		accountsClient: client,
	}
}
