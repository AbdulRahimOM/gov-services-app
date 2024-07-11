package routes

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/acc-handler"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/appointments"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.RouterGroup, adminAccHandler *acchandler.AdminAccountHandler, appointmentHandler *appointments.AppointmentHandler) {
	engine.Use(middleware.ClearCache)

	//routes for those who are not logged in______________________________________________
	strangersGroup := engine.Group("/admin")
	strangersGroup.Use(middleware.NotLoggedInCheck)
	{
		strangersGroup.GET("/login", adminAccHandler.Ping)                                  //done
		strangersGroup.POST("/login-using-password", adminAccHandler.AdminLoginViaPassword) //done
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

		authGroup.GET("/manage-admins/view", adminAccHandler.AdminGetAdmins) //done
		
		offices:= authGroup.Group("/offices")	//...................../admin/offices
		{
			offices.GET("/view", adminAccHandler.AdminGetOffices) 
			offices.PUT("/new-child-office", appointmentHandler.CreateChildOffice)
		}

		appointments := authGroup.Group("/appoint")	//...................../admin/appoint
		{
			appointments.PUT("/attender", appointmentHandler.AppointAttender) //appoint other attenders of an office
			appointments.PUT("/child-office-head", appointmentHandler.AppointChildOfficeHead)
			appointments.PUT("/child-office-deputy-head", appointmentHandler.AppointChildOfficeDeputyHead)
		}
	}

}
