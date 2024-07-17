package ksebUserUc

import (
	repo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/interface"
	usecase "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
)

type KsebUserUseCase struct {
	userRepo repo.IUserRepo
}

func NewKsebUserUseCase(userRepo repo.IUserRepo) usecase.IKsebUserUC {
	return &KsebUserUseCase{
		userRepo: userRepo,
	}
}
