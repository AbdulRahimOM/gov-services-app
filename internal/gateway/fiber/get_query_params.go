package gateway

import (
	"fmt"

	"github.com/AbdulRahimOM/go-utils/mymath"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"github.com/gofiber/fiber/v2"
)

/*
HandleGetQueryParamsInt32 retrieves a query parameter from the context request and converts it to int32

Boolean indicates if the conversion was successful or not.

If the parameter is not found, it returns 0 and true.

If the parameter is found and valid, it returns the converted value and true.

If the parameter is found but invalid, it returns {0, false} and sends a JSON error response to the client.
*/
func HandleGetQueryParamsInt32Fiber(c *fiber.Ctx, key string) (int32, error) {
	str := c.Query(key)
	if str == "" {
		return 0, nil
	}

	val, err := mymath.StringToInt32(str)
	if err != nil {
		return 0, c.Status(400).JSON(stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respcode.InvalidQueryParams,
			Error:        fmt.Sprintf("Invalid %s: %s", key, str),
		})
	}

	return val, nil
}

func HandleGetUrlParamsInt32Fiber(c *fiber.Ctx, key string) (int32, error) {
	str := c.Params(key)
	if str == "" {
		return 0, nil
	}

	val, err := mymath.StringToInt32(str)
	if err != nil {
		return 0, c.Status(400).JSON(stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respcode.InvalidUrlParams,
			Error:        fmt.Sprintf("Invalid %s: %s", key, str),
		})
	}

	return val, nil
}