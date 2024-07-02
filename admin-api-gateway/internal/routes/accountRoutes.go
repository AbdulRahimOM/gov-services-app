package routes

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.RouterGroup, entryHandler *handler.AccountEntryHandler, adminAccHandler *handler.AdminAccountHandler) {
	engine.Use(middleware.ClearCache)

	//routes for those who are not logged in______________________________________________
	strangersGroup := engine.Group("/admin")
	strangersGroup.Use(middleware.NotLoggedInCheck)
	{
		strangersGroup.GET("/login", entryHandler.Ping)                                  //done
		strangersGroup.POST("/login-using-password", entryHandler.AdminLoginViaPassword) //done
	}

	//routes for those who are logged in-------------------------------------------------
	authGroup := engine.Group("/admin")
	authGroup.Use(middleware.AdminAuthCheck)
	{
		profileGroup := authGroup.Group("/profile")
		{
			profileGroup.GET("/view", adminAccHandler.AdminGetProfile)                                        //done
			profileGroup.GET("/edit-page", adminAccHandler.AdminGetProfile)                                   //done
			profileGroup.POST("/update", adminAccHandler.AdminUpdateProfile)                                  //done
			profileGroup.POST("/update-password/using-old-pw", adminAccHandler.AdminUpdatePasswordUsingOldPw) //done
		}

		authGroup.GET("/view-admins", adminAccHandler.AdminGetAdmins) //unimplemented
	}
}
