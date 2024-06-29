package gateway

import (
	mystatus "github.com/AbdulRahimOM/gov-services-app/shared/std-response/my_status"
	respcode "github.com/AbdulRahimOM/gov-services-app/shared/std-response/response-code"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/shared/std-response/std-response"
	"github.com/gin-gonic/gin"
)

func GetUserIdFromContext(c *gin.Context) (int32, bool) {
	userID, found := c.Get("userID")
	if !found {
		c.JSON(400, stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respcode.BugNoUserInContext,
			Error:        "user ID not found in context",
		})
		return 0, false
	}
	
	return userID.(int32), true
}
