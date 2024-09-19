package userAccUC

import (
	"log"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/config"
	repo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/interface"
	usecase "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	jwttoken "github.com/AbdulRahimOM/gov-services-app/internal/jwt-token"
	twilioOTP "github.com/AbdulRahimOM/gov-services-app/internal/twilio"
)

type UserUseCase struct {
	userRepo        repo.IUserRepo
	jwtClient       *jwttoken.TokenGenerator
	twilioOTPClient *twilioOTP.TwilioClient
}

func NewUserUseCase(userRepo repo.IUserRepo) usecase.IUserUC {
	jwtClient, err := jwttoken.NewTokenGenerator("./internal/config/private.key")
	if err != nil {
		log.Fatalf("Failed to create token generator: %v", err)
	}

	twilioOTPClient := twilioOTP.NewTwilioClient(config.Twilio.AccountSid, config.Twilio.AuthToken, config.Twilio.ServiceSid, config.DevMode.ByPassTwilio)
	return &UserUseCase{
		userRepo:        userRepo,
		jwtClient:       jwtClient,
		twilioOTPClient: twilioOTPClient,
	}
}
