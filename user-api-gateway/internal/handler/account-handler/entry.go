package acchandler

import (
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/request"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/response"

	"github.com/gin-gonic/gin"
)

func (u *UserAccountHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (u *UserAccountHandler) RequestOTPForLogin(c *gin.Context) {
	var req request.UserLoginGetOTP

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	_, err := u.accountsClient.UserLoginGetOTP(c, &pb.UserLoginGetOTPRequest{
		PhoneNumber: req.PhoneNumber,
	})
	if err == nil {
		c.JSON(200, response.GetOTPResponse{
			Status:      mystatus.Success,
			Last4Digits: req.PhoneNumber[len(req.PhoneNumber)-4:],
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

func (u *UserAccountHandler) UserLoginVerifyOTP(c *gin.Context) {
	var req request.UserLoginVerifyOTP

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	resp, err := u.accountsClient.UserLoginVerifyOTP(c, &pb.UserLoginVerifyOTPRequest{
		PhoneNumber: req.PhoneNumber,
		Otp:         req.OTP,
	})

	if err == nil {
		c.JSON(200, response.UserLogin{
			Status: mystatus.Success,
			UserData: response.UserBasicData{
				Id:          resp.UserDetails.Id,
				FirstName:   resp.UserDetails.FirstName,
				LastName:    resp.UserDetails.LastName,
				PhoneNumber: resp.UserDetails.PhoneNumber,
			},
			Token: resp.Token,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

func (u *UserAccountHandler) RequestOTPForSignUp(c *gin.Context) {
	var req request.GetOTPForSignup

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	_, err := u.accountsClient.UserSignUpGetOTP(c, &pb.UserSignUpGetOTPRequest{
		PhoneNumber: req.PhoneNumber,
	})
	if err == nil {
		c.JSON(200, response.GetOTPResponse{
			Status:      mystatus.Success,
			Last4Digits: req.PhoneNumber[len(req.PhoneNumber)-4:],
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

func (u *UserAccountHandler) SubmitOTPForSignUp(c *gin.Context) {
	var req request.UserSignpViaOTP

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	resp, err := u.accountsClient.UserSignUpVerifyOTP(c, &pb.UserSignUpVerifyOTPRequest{
		PhoneNumber: req.PhoneNumber,
		Otp:         req.OTP,
	})
	if err == nil {
		c.JSON(200, response.UserSignUp{
			Status:  mystatus.Success,
			Message: "User signed up successfully",
			User: response.PreliminaryUserData{
				Id:          resp.UserDetails.Id,
				PhoneNumber: resp.UserDetails.PhoneNumber,
			},
			Token: resp.Token,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

func (u *UserAccountHandler) UserLoginViaPassword(c *gin.Context) {
	var req request.UserLoginViaPassword

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}
	resp, err := u.accountsClient.UserLoginViaPassword(c, &pb.UserLoginViaPasswordRequest{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})
	if err == nil {
		c.JSON(200, response.UserLogin{
			Status: mystatus.Success,
			UserData: response.UserBasicData{
				Id:          resp.UserDetails.Id,
				FirstName:   resp.UserDetails.FirstName,
				LastName:    resp.UserDetails.LastName,
				PhoneNumber: resp.UserDetails.PhoneNumber,
			},
			Token: resp.Token,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}
