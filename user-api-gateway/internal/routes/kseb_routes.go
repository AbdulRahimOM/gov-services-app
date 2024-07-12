package routes

import (
	ksebhandler "github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/handler/kseb-handler"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterKsebRoutes(engine *gin.RouterGroup, ksebHandler *ksebhandler.KsebHandler) {
	engine.Use(middleware.ClearCache)

	authGroup := engine.Group("/user")
	authGroup.Use(middleware.UserAuthCheck)
	{
		authGroup.PUT("/add-consumer-number", ksebHandler.AddConsumerNumber)
		authGroup.GET("/get-my-consumer-numbers", ksebHandler.GetUserConsumerNumbers)

		authGroup.POST("/raise-complaint", ksebHandler.RaiseComplaint)
	}


}
