package routes

import (
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/handler"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.RouterGroup, entryHandler *handler.AccountEntryHandler, userAccHandler *handler.UserAccountHandler) {
	engine.Use(middleware.ClearCache)

	//routes for those who are not logged in______________________________________________
	strangersGroup := engine.Group("/")
	strangersGroup.Use(middleware.NotLoggedInCheck)
	{
		strangersGroup.GET("/login", entryHandler.Ping)                          //done
		strangersGroup.POST("/login-getOTP", entryHandler.RequestOTPForLogin)    //done
		strangersGroup.POST("/login-submitOTP", entryHandler.UserLoginVerifyOTP) //done

		strangersGroup.POST("/login-using-password", entryHandler.UserLoginViaPassword) //done

		strangersGroup.GET("/register", entryHandler.Ping)                          //done
		strangersGroup.POST("/register-getOTP", entryHandler.RequestOTPForSignUp)   //done
		strangersGroup.POST("/register-submitOTP", entryHandler.SubmitOTPForSignUp) //done

	}

	//routes for those who just signed up - signup process is not complete---------------
	newUserGroup := engine.Group("/new-user")
	newUserGroup.Use(middleware.NewUserCheck)
	{
		newUserGroup.POST("/profile/set-password", userAccHandler.SignedUpUserSettingPw) //done
	}

	//routes for those who are logged in-------------------------------------------------
	authGroup := engine.Group("/user")
	authGroup.Use(middleware.UserAuthCheck)
	{
		profileGroup := authGroup.Group("/profile")
		{
			profileGroup.GET("/view", userAccHandler.UserGetProfile)                                                 //done
			profileGroup.GET("/edit-page", userAccHandler.UserGetProfile)                                            //done
			profileGroup.POST("/update", userAccHandler.UserUpdateProfile)                                           //done
			profileGroup.POST("/update-password/using-old-pw", userAccHandler.UserUpdatePasswordUsingOldPw)          //done
			profileGroup.GET("/update-password/using-otp/get-otp", userAccHandler.UserGetOTPForPwChange)             //not tested
			profileGroup.POST("/update-password/using-otp/verify-otp", userAccHandler.UserVerifyOTPForPwChange)      //not tested
			profileGroup.POST("/update-password/using-otp/set-new-pw", userAccHandler.UserSetNewPwAfterVerifyingOTP) //not tested
		}
	}
}
