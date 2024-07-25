package middleware

import (
	"fmt"
	"log"
	"strings"

	jwttoken "github.com/AbdulRahimOM/gov-services-app/internal/jwt-token"
	respCode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/response"
	"github.com/gofiber/fiber/v2"
)

var jwtVerifier *jwttoken.TokenVerifier

func init() {
	var err error
	jwtVerifier, err = jwttoken.NewTokenVerifier("./internal/config/public.key")
	if err != nil {
		log.Fatalf("Failed to create token verifier: %v", err)
	}
}

func UserAuthCheck(c *fiber.Ctx) error {
	fmt.Println("====UserAuthCheck====")
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	isTokenValid, accInfo, addlInfo, err := jwtVerifier.ValidateToken(tokenString)
	if !isTokenValid {
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
		return c.Redirect("/user/profile/set-password")
	default:
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
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	isTokenValid, accInfo, _, err := jwtVerifier.ValidateToken(tokenString)
	if !isTokenValid {
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
		return c.Status(401).JSON(response.SRE{
			Status:       "failed",
			ResponseCode: respCode.Unauthorized,
			Error:        "Invalid role. Should be 'password-not-set-user', but is " + accInfo.Role,
		})
	}
}
