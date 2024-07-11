package adminaccount

import (
	"fmt"
	"time"

	hashpassword "github.com/AbdulRahimOM/go-utils/hashPassword"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/response"
	jwttoken "github.com/AbdulRahimOM/gov-services-app/internal/jwt-token"
	responsecode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// VerifyPasswordForLogin
func (uc AdminUseCase) VerifyPasswordForLogin(username, password *string) (*response.AdminLogin, string, error) {
	//verify password
	admin, hashedPw, err := uc.adminRepo.GetAdminWithPasswordByUsername(username)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, responsecode.CorruptRequest, fmt.Errorf("no admin registered with this phoneNumber number")
		} else {
			return nil, responsecode.DBError, fmt.Errorf("at database: failed to get admin password: %v", err)
		}
	}

	err = hashpassword.CompareHashedPassword(*hashedPw, *password)
	if err != nil {
		return nil, responsecode.InvalidPassword, fmt.Errorf("pw verification failed: %v", err)
	}

	//generate jwt token
	token, err := uc.jwtClient.GenerateToken(
		jwttoken.AccountInfo{
			Role: "admin",
			Id:   admin.ID,
		}, nil, time.Minute*time.Duration(config.JWT.ExpTimeInMinutes))
	if err != nil {
		return nil, responsecode.OtherInternalError, fmt.Errorf("failed to generate token: %v", err)
	}

	return &response.AdminLogin{
		Admin: admin,
		Token: token,
	}, "", nil
}
