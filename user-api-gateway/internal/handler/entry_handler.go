package handler

import (
	"github.com/AbdulRahimOM/gov-services-app/shared/gateway"
	mystatus "github.com/AbdulRahimOM/gov-services-app/shared/std-response/my_status"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/request"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/response"
	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"

	"github.com/gin-gonic/gin"
)

func (u *AccountEntryHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (u *AccountEntryHandler) RequestOTPForLogin(c *gin.Context) {
	var req request.GetOTPForLogin

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	_, err := u.accountsClient.GetOTPForLogin(c, &pb.GetOTPForLoginRequest{
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

func (u *AccountEntryHandler) UserLoginViaOTP(c *gin.Context) {
	var req request.UserLoginViaOTP

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	resp, err := u.accountsClient.UserLoginViaOTP(c, &pb.UserLoginViaOTPRequest{
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

func (u *AccountEntryHandler) RequestOTPForSignUp(c *gin.Context) {
	var req request.GetOTPForSignup

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	_, err := u.accountsClient.GetOTPForSignUp(c, &pb.GetOTPForSignUpRequest{
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

func (u *AccountEntryHandler) SubmitOTPForSignUp(c *gin.Context) {
	var req request.UserSignpViaOTP

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	resp, err := u.accountsClient.UserSignUpViaOTP(c, &pb.UserSignUpViaOTPRequest{
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

func (u *AccountEntryHandler) UserLoginViaPassword(c *gin.Context) {
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
