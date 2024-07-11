package adminaccount

import (
	"log"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/config"
	repo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/interface"
	usecase "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	jwttoken "github.com/AbdulRahimOM/gov-services-app/internal/jwt-token"
	twilioOTP "github.com/AbdulRahimOM/gov-services-app/internal/twilio"
)

type AdminUseCase struct {
	adminRepo       repo.IAdminRepo
	jwtClient       *jwttoken.TokenGenerator
	twilioOTPClient *twilioOTP.TwilioClient
}

func NewAdminUseCase(adminRepo repo.IAdminRepo) usecase.IAdminUC {
	jwtClient, err := jwttoken.NewTokenGenerator("./internal/config/private.key")
	if err != nil {
		log.Fatalf("Failed to create token generator: %v", err)
	}

	twilioOTPClient := twilioOTP.NewTwilioClient(config.Twilio.AccountSid, config.Twilio.AuthToken, config.Twilio.ServiceSid)
	return &AdminUseCase{
		adminRepo:       adminRepo,
		jwtClient:       jwtClient,
		twilioOTPClient: twilioOTPClient,
	}
}
