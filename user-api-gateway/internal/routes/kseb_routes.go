package routes

import (
	"time"

	ksebhandler "github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/handler/kseb-handler"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func RegisterKsebRoutes(api fiber.Router, ksebHandler *ksebhandler.KsebHandler) {
	api.Use(middleware.ClearCache)

	authGroup := api.Group("/user")
	authGroup.Use(middleware.UserAuthCheck)
	{
		authGroup.Put("/add-consumer-number", ksebHandler.AddConsumerNumber)
		authGroup.Get("/get-my-consumer-numbers", ksebHandler.GetUserConsumerNumbers)

		authGroup.Post("/raise-complaint", ksebHandler.RaiseComplaint)

		//chat
		authGroup.Get("/chat/:complaintId/websocket", websocket.New(ksebHandler.UserChatWebsocket, websocket.Config{
			HandshakeTimeout: 10 * time.Second,
		}))
	}

}
