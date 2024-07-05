package routes

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.RouterGroup, adminAccHandler *handler.AdminAccountHandler) {
	engine.Use(middleware.ClearCache)

	//routes for those who are not logged in______________________________________________
	strangersGroup := engine.Group("/admin")
	strangersGroup.Use(middleware.NotLoggedInCheck)
	{
		strangersGroup.GET("/login",adminAccHandler.Ping)                                  //done
		strangersGroup.POST("/login-using-password",adminAccHandler.AdminLoginViaPassword) //done
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

		manageAdminsGroup := authGroup.Group("/manage-admins")
		{
			manageAdminsGroup.GET("/view", adminAccHandler.AdminGetAdmins)  //done
			manageAdminsGroup.PUT("/add", adminAccHandler.AdminAddSubAdmin) //implementing
		}

		// deptGroup := authGroup.Group("/depts")
		// {
		// deptGroup.PUT("/add", adminAccHandler.AdminAddDept) //done
		// deptGroup.GET("/view", adminAccHandler.AdminGetDepts) //done
		// }

		officeGroup := authGroup.Group("/offices")
		{
			// officeGroup.PUT("/add", adminAccHandler.AdminCreateSubOffice) //
			officeGroup.GET("/view", adminAccHandler.AdminGetOffices) //
		}

		superAdminGroup := authGroup.Group("/rank-1-role")
		superAdminGroup.Use(middleware.CheckRank(1))
		{
			// appointments := superAdminGroup.Group("/appoint")
			{
				// appointments.PUT("/SNO", adminAccHandler.AppointStateNodalOfficer)
				// appointments.PUT("/state-SDyNO", adminAccHandler.AppointStateDeputyNodalOfficer)
			}
		}

	}
}
