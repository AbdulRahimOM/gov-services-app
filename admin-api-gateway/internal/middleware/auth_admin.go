package middleware

import (
	"log"
	"strings"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	jwttoken "github.com/AbdulRahimOM/gov-services-app/internal/jwt-token"
	respCode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"

	"github.com/gin-gonic/gin"
)

var jwtVerifier *jwttoken.TokenVerifier

func init() {
	var err error
	jwtVerifier, err = jwttoken.NewTokenVerifier("./internal/config/public.key")
	if err != nil {
		log.Fatalf("Failed to create token verifier: %v", err)
	}
}
func AdminAuthCheck(c *gin.Context) {
	tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	isTokenValid, accInfo, addlInfo, err := jwtVerifier.ValidateToken(tokenString)
	if !isTokenValid {
		c.JSON(401, response.SRE{
			Status:       "failed",
			ResponseCode: respCode.Unauthenticated,
			Error:        err.Error(),
		})
		c.Abort()
		return
	}

	switch accInfo.Role {
	case "admin":
		c.Set("role", accInfo.Role)
		c.Set("adminID", accInfo.Id)
	case "password-not-set-admin":
		c.Redirect(302, "/admin/profile/set-password")
	default:
		c.JSON(401, response.SRE{
			Status:       "failed",
			ResponseCode: respCode.Unauthorized,
			Error:        "Invalid role. Should be admin, but is " + accInfo.Role,
		})
		c.Abort()
		return
	}

	if addlInfo != nil {
		addlInfo.SetContext(c)
	}

	c.Next()
}
