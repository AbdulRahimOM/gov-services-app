package routes

import (
	ksebhanlder "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/kseb"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterKSEBAccRoutes(engine *gin.RouterGroup, ksebHandler *ksebhanlder.KSEBHandler) {
	engine.Use(middleware.ClearCache)

	//admin routes
	authGroup := engine.Group("/admin")
	authGroup.Use(middleware.AdminAuthCheck)
	{
		authGroup.PUT("/register-section-code", ksebHandler.KSEBRegisterSectionCode)

		//chat
		// authGroup.POST("/chat/send-message", ksebHandler.SendMessage)
		// authGroup.GET("/chat", ksebHandler.Chat)
	}
}
