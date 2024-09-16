package gateway

import (
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func GetUserIdFromContextFiber(c *fiber.Ctx) (int32,bool, error) {
	userLogger.WithField("method", "GetUserIdFromContextFiber")
	userID := c.Locals("userID")
	if userID == nil {
		userLogger.Debug("user ID not found in context")
		return 0, false,c.Status(400).JSON(stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respcode.BugNoUserInContext,
			Error:        "user ID not found in context",
		})
	}
	userIDInt, ok := userID.(int32)
	if !ok {
		userLogger.Debug("userID in context is not int32")
		return 0, false,c.Status(400).JSON(stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respcode.BugNoUserInContext,
			Error:        "user ID not found in context",
		})
	}

	return userIDInt, true,nil
}

func GetAdminIdFromContextFiber(c *fiber.Ctx) (int32, bool,error) {
	adminLogger.WithField("method", "GetAdminIdFromContext")

	adminID := c.Locals("adminID")
	if adminID == nil {
		adminLogger.Debug("admin ID not found in context")
		return 0,false, c.Status(400).JSON(stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respcode.BugNoAdminInContext,
			Error:        "admin ID not found in context",
		})
	}
	adminIDInt, ok := adminID.(int32)
	if !ok {
		adminLogger.Debug("adminID in context is not int32")
		return 0,false, c.Status(400).JSON(stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respcode.BugNoAdminInContext,
			Error:        "admin ID not found in context",
		})
	}

	return adminIDInt,true, nil
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
