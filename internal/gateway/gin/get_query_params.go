package gateway

import (
	"fmt"

	"github.com/AbdulRahimOM/go-utils/mymath"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"github.com/gin-gonic/gin"
)

/*
HandleGetQueryParamsInt32 retrieves a query parameter from the context request and converts it to int32

Boolean indicates if the conversion was successful or not.

If the parameter is not found, it returns 0 and true.

If the parameter is found and valid, it returns the converted value and true.

If the parameter is found but invalid, it returns {0, false} and sends a JSON error response to the client.
*/
func HandleGetQueryParamsInt32(c *gin.Context, key string) (int32, bool) {
	str := c.DefaultQuery(key, "")
	if str == "" {
		return 0, true
	}

	val, err := mymath.StringToInt32(str)
	if err != nil {
		c.JSON(400, stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respcode.InvalidQueryParams,
			Error:        fmt.Sprintf("Invalid %s: %s", key, str),
		})
		return 0, false
	}

	return val, true
}

func HandleGetUrlParamsInt32(c *gin.Context, key string) (int32, bool) {
	str := c.Param(key)
	if str == "" {
		return 0, true
	}

	val, err := mymath.StringToInt32(str)
	if err != nil {
		c.JSON(400, stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respcode.InvalidUrlParams,
			Error:        fmt.Sprintf("Invalid %s: %s", key, str),
		})
		return 0, false
	}

	return val, true
}
