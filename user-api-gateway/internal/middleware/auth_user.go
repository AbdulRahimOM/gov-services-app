package middleware

import (
	"log"
	"strings"

	jwttoken "github.com/AbdulRahimOM/gov-services-app/internal/jwt-token"
	"github.com/AbdulRahimOM/gov-services-app/internal/logs"
	respCode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/response"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var (
	jwtVerifier *jwttoken.TokenVerifier
	logger      *logrus.Entry
)

func init() {
	var err error
	jwtVerifier, err = jwttoken.NewTokenVerifier("./internal/config/public.key")
	if err != nil {
		log.Fatalf("Failed to create token verifier: %v", err)
	}

	logger = logs.NewLoggerWithServiceName("user-api-gateway-middleware")
}

func UserAuthCheck(c *fiber.Ctx) error {
	logger.WithField("method", "UserAuthCheck")

	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	isTokenValid, accInfo, addlInfo, err := jwtVerifier.ValidateToken(tokenString)
	if !isTokenValid {

		logger.Info("Invalid token attempted access")
		return c.Status(401).JSON(response.SRE{
			Status:       "failed",
			ResponseCode: respCode.Unauthenticated,
			Error:        err.Error(),
		})
	}
	switch accInfo.Role {
	case "user":
		c.Locals("role", accInfo.Role)
		c.Locals("userID", accInfo.Id)
	case "password-not-set-user":
		logger.Warn("Password not set user attempting to access other routes")
		return c.Redirect("/user/profile/set-password")
	default:
		logger.Info("Invalid role attempting to access other routes")

		return c.Status(401).JSON(response.SRE{
			Status:       "failed",
			ResponseCode: respCode.Unauthorized,
			Error:        "Invalid role. Should be user, but is " + accInfo.Role,
		})
	}

	if addlInfo != nil {
		addlInfo.SetContextFiber(c)
	}

	return c.Next()
}

func NewUserCheck(c *fiber.Ctx) error {
	logger.WithFields(logrus.Fields{
		"method": "NewUserCheck",
	})
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	isTokenValid, accInfo, _, err := jwtVerifier.ValidateToken(tokenString)
	if !isTokenValid {
		logger.Info("Unauthenticated user attempting to access other routes")
		return c.Status(401).JSON(response.SRE{
			Status:       "failed",
			ResponseCode: respCode.Unauthenticated,
			Error:        err.Error(),
		})
	}
	switch accInfo.Role {
	case "password-not-set-user":
		c.Locals("role", accInfo.Role)
		c.Locals("userID", accInfo.Id)
		return c.Next()
	default:
		logger.Info("Invalid role attempting to access other routes")
		return c.Status(401).JSON(response.SRE{
			Status:       "failed",
			ResponseCode: respCode.Unauthorized,
			Error:        "Invalid role. Should be 'password-not-set-user', but is " + accInfo.Role,
		})
	}
}
