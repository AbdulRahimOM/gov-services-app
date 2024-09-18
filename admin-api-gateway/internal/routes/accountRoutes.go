package routes

import (
	acchandler "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/acc-handler"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/handler/appointments"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router, adminAccHandler *acchandler.AdminAccountHandler, appointmentHandler *appointments.AppointmentHandler) {

	//routes for those who are not logged in______________________________________________
	strangersGroup := api.Group("/admin/entry")
	strangersGroup.Use(middleware.NotLoggedInCheck)
	{
		strangersGroup.Get("/login", adminAccHandler.Ping)                                  //done
		strangersGroup.Post("/login-using-password", adminAccHandler.AdminLoginViaPassword) //done
	}

	//routes for those who are logged in-------------------------------------------------
	authGroup := api.Group("/admin/auth")
	authGroup.Use(middleware.AdminAuthCheck)
	{
		profileGroup := authGroup.Group("/profile")
		{
			profileGroup.Get("/view", adminAccHandler.AdminGetProfile)                                        //done
			profileGroup.Get("/edit-page", adminAccHandler.AdminGetProfile)                                   //done
			profileGroup.Post("/update", adminAccHandler.AdminUpdateProfile)                                  //done
			profileGroup.Post("/update-password/using-old-pw", adminAccHandler.AdminUpdatePasswordUsingOldPw) //done
		}

		authGroup.Get("/manage-admins/view", adminAccHandler.AdminGetAdmins) //done

		offices := authGroup.Group("/offices") //...................../admin/offices
		{
			offices.Get("/view", adminAccHandler.AdminGetOffices)
			offices.Put("/new-child-office", appointmentHandler.CreateChildOffice)
		}

		appointments := authGroup.Group("/appoint") //...................../admin/appoint
		{
			appointments.Put("/attender", appointmentHandler.AppointAttender) //appoint other attenders of an office
			appointments.Put("/child-office-head", appointmentHandler.AppointChildOfficeHead)
			appointments.Put("/child-office-deputy-head", appointmentHandler.AppointChildOfficeDeputyHead)
		}
	}

}
