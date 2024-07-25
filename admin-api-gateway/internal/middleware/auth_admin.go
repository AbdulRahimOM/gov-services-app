package middleware

import (
	"fmt"
	"log"
	"strings"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	jwttoken "github.com/AbdulRahimOM/gov-services-app/internal/jwt-token"
	respCode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
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

func AdminAuthCheck(c *fiber.Ctx) error {
	fmt.Println("====AdminAuthCheck====")
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
	case "admin":
		c.Locals("role", accInfo.Role)
		c.Locals("adminID", accInfo.Id)
	case "password-not-set-admin":
		return c.Redirect("/admin/profile/set-password")
	default:
		return c.Status(401).JSON(response.SRE{
			Status:       "failed",
			ResponseCode: respCode.Unauthorized,
			Error:        "Invalid role. Should be admin, but is " + accInfo.Role,
		})
	}

	if addlInfo != nil {
		addlInfo.SetContextFiber(c)
	}

	return c.Next()
}
