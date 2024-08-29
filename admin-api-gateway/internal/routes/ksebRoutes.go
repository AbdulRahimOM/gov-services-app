package routes

import (
	"time"

	ksebhanlder "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/kseb"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func RegisterKSEBAccRoutes(api fiber.Router, ksebHandler *ksebhanlder.KSEBHandler) {
	api.Use(middleware.ClearCache)

	//admin routes
	authGroup := api.Group("/admin")
	authGroup.Use(middleware.AdminAuthCheck)
	{
		authGroup.Put("/register-section-code", ksebHandler.KSEBRegisterSectionCode)
		authGroup.Get("/get-complaints", ksebHandler.AdminGetComplaints)
		authGroup.Get("/open-complaint/:complaintId", ksebHandler.AdminOpenComplaint)
		authGroup.Post("/close-complaint", ksebHandler.AdminCloseComplaint)

		//chat
		authGroup.Get("/chat/:complaintId/ws", websocket.New(ksebHandler.AdminChatWebsocket, websocket.Config{
			HandshakeTimeout: 10 * time.Second,
		}))

		authGroup2 := api.Group("/oadmin")
		{
			authGroup2.Get("/welcome", ksebHandler.Welcome)
			authGroup2.Get("/videocall/room/:complaintId", ksebHandler.VideoCallRoom)
			authGroup2.Get("/videocall/room/:complaintId/websocket", websocket.New(ksebHandler.RoomWebsocket, websocket.Config{
				HandshakeTimeout: 10 * time.Second,
			}))
		}
	}

}
