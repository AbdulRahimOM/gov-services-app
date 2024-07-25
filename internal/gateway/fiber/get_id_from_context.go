package gateway

import (
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func GetUserIdFromContextFiber(c *fiber.Ctx) (int32, error) {
	userID := c.Locals("userID")
	if userID == nil {
		return 0, c.Status(400).JSON(stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respcode.BugNoUserInContext,
			Error:        "user ID not found in context",
		})
	}
	userIDInt, ok := userID.(int32)
	if !ok {
		return 0, c.Status(400).JSON(stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respcode.BugNoUserInContext,
			Error:        "user ID not found in context",
		})
	}

	return userIDInt, nil
}

func GetAdminIdFromContextFiber(c *fiber.Ctx) (int32, error) {
	adminID := c.Locals("adminID")
	if adminID == nil {
		return 0, c.Status(400).JSON(stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respcode.BugNoAdminInContext,
			Error:        "admin ID not found in context",
		})
	}
	adminIDInt, ok := adminID.(int32)
	if !ok {
		return 0, c.Status(400).JSON(stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respcode.BugNoAdminInContext,
			Error:        "admin ID not found in context",
		})
	}

	return adminIDInt, nil
}
func GetUserIdFromWebsocketConn(c *websocket.Conn) (int32, bool) {
	userID := c.Locals("userID")
	if userID == nil {
		return 0, false
	}
	userIDInt, ok := userID.(int32)
	if !ok {
		return 0, false
	}

	return userIDInt, true
}
func GetAdminIdFromWebsocketConn(c *websocket.Conn) (int32, bool) {
	adminID := c.Locals("adminID")
	if adminID == nil {
		return 0, false
	}
	adminIDInt, ok := adminID.(int32)
	if !ok {
		return 0, false
	}

	return adminIDInt, true
}
