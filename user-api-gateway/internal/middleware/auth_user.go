package middleware

import (
	"log"
	"strings"

	jwttoken "github.com/AbdulRahimOM/gov-services-app/internal/jwt-token"
	respCode "github.com/AbdulRahimOM/gov-services-app/shared/std-response/response-code"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/response"

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
func UserAuthCheck(c *gin.Context) {
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
	case "user":
		c.Set("role", accInfo.Role)
		c.Set("userID", accInfo.Id)
	case "password-not-set-user":
		c.Redirect(302, "/user/profile/set-password")
	default:
		c.JSON(401, response.SRE{
			Status:       "failed",
			ResponseCode: respCode.Unauthorized,
			Error:        "Invalid role. Should be user, but is " + accInfo.Role,
		})
		c.Abort()
		return
	}

	if addlInfo != nil {
		addlInfo.SetContext(c)
	}

	c.Next()
}

func NewUserCheck(c *gin.Context) {
	tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	isTokenValid, accInfo, _, err := jwtVerifier.ValidateToken(tokenString)
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
	case "password-not-set-user":
		c.Set("role", accInfo.Role)
		c.Set("userID", accInfo.Id)
		c.Next()
	default:
		c.JSON(401, response.SRE{
			Status:       "failed",
			ResponseCode: respCode.Unauthorized,
			Error:        "Invalid role. Should be 'password-not-set-user', but is " + accInfo.Role,
		})
		c.Abort()
		return
	}

}
