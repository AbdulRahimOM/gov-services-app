package routes

import (
	acchandler "github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/handler/account-handler"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router, userAccHandler *acchandler.UserAccountHandler) {
	api.Use(middleware.ClearCache)

	//routes for those who are not logged in______________________________________________
	strangersGroup := api.Group("/entry")
	strangersGroup.Use(middleware.NotLoggedInCheck)
	{
		strangersGroup.Get("/login", userAccHandler.Ping)                          //done
		strangersGroup.Post("/login-getOTP", userAccHandler.RequestOTPForLogin)    //done
		strangersGroup.Post("/login-submitOTP", userAccHandler.UserLoginVerifyOTP) //done

		strangersGroup.Post("/login-using-password", userAccHandler.UserLoginViaPassword) //done

		strangersGroup.Get("/register", userAccHandler.Ping)                          //done
		strangersGroup.Post("/register-getOTP", userAccHandler.RequestOTPForSignUp)   //done
		strangersGroup.Post("/register-submitOTP", userAccHandler.SubmitOTPForSignUp) //done

	}

	//routes for those who just signed up - signup process is not complete---------------
	newUserGroup := api.Group("/new-user")
	newUserGroup.Use(middleware.NewUserCheck)
	{
		newUserGroup.Post("/profile/set-password", userAccHandler.SignedUpUserSettingPw) //done
	}

	//routes for those who are logged in-------------------------------------------------
	authGroup := api.Group("/user")
	authGroup.Use(middleware.UserAuthCheck)
	{
		profileGroup := authGroup.Group("/profile")
		{
			profileGroup.Get("/view", userAccHandler.UserGetProfile)                                           
			profileGroup.Get("/edit-page", userAccHandler.UserGetProfile)                                            //done
			profileGroup.Post("/update", userAccHandler.UserUpdateProfile)                                           //done
			profileGroup.Post("/update-password/using-old-pw", userAccHandler.UserUpdatePasswordUsingOldPw)          //done
			profileGroup.Get("/update-password/using-otp/get-otp", userAccHandler.UserGetOTPForPwChange)             //not tested
			profileGroup.Post("/update-password/using-otp/verify-otp", userAccHandler.UserVerifyOTPForPwChange)      //not tested
			profileGroup.Post("/update-password/using-otp/set-new-pw", userAccHandler.UserSetNewPwAfterVerifyingOTP) //not tested
		}
	}
}
