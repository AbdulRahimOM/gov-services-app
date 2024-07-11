package middleware

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	respCode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	"github.com/gin-gonic/gin"
)

// RankCheck
func CheckRank(minRank int) gin.HandlerFunc {
	return func(c *gin.Context) {
		rank:= c.GetInt("rank")
		if rank<=0 {
			c.JSON(401, response.SRE{
				Status:       "failed",
				ResponseCode: respCode.Unauthorized,
				Error:        "Invalid rank",
			})
			c.Abort()
			return
		}else if rank<minRank{
			c.JSON(401, response.SRE{
				Status:       "failed",
				ResponseCode: respCode.Unauthorized,
				Error:        "Insufficient rank",
			})
			c.Abort()
			return
		}
	}
}
