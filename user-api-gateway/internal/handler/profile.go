package handler

import (
	"context"
	"fmt"

	"github.com/AbdulRahimOM/gov-services-app/shared/gateway"
	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/shared/std-response/my_status"
	respCode "github.com/AbdulRahimOM/gov-services-app/shared/std-response/response-code"
	"github.com/AbdulRahimOM/gov-services-app/shared/tag"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/request"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/response"
	"github.com/gin-gonic/gin"
)

func (u *UserAccountHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// GetOTPForPwChange
func (u *UserAccountHandler) GetOTPForPwChange(c *gin.Context) { //unimplemented
	userID, ok := gateway.GetUserIdFromContext(c)
	if !ok {
		return
	}

	resp, err := u.accountsClient.GetOTPForPwChange(context.Background(), &pb.GetOTPForPwChangeRequest{
		UserId: userID,
	})
	if err == nil {
		c.JSON(200, response.GetOTPResponse{
			Status:      mystatus.Success,
			Last4Digits: resp.Last4Digits,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}

}

// VerifyOTPForPwChange
func (u *UserAccountHandler) VerifyOTPForPwChange(c *gin.Context) { //unimplemented

	var req request.VerifyOTPForPwChange
	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	userID, ok := gateway.GetUserIdFromContext(c)
	if !ok {
		return
	}

	resp, err := u.accountsClient.VerifyOTPForPwChange(context.Background(), &pb.VerifyOTPForPwChangeRequest{
		UserId: userID,
		Otp:    req.Otp,
	})
	if err == nil {
		c.JSON(200, response.VerifyOTPForPwChangeResponse{
			Status:    mystatus.Success,
			Msg:       "OTP verified",
			TempToken: resp.TempToken,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

func (u *UserAccountHandler) SignedUpUserSettingPw(c *gin.Context) { //unimplemented
	var req request.SettingNewPassword

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	userID, ok := gateway.GetUserIdFromContext(c)
	if !ok {
		return
	}

	resp, err := u.accountsClient.SignedUpUserSettingPw(context.Background(), &pb.SignedUpUserSettingPwRequest{
		UserId:      userID,
		NewPassword: req.NewPassword,
	})
	if err == nil {
		c.JSON(200, response.UpdateToken{
			Status: mystatus.Success,
			Token:  resp.Token,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

// UpdateProfile
func (u *UserAccountHandler) UpdateProfile(c *gin.Context) { //unimplemented
	var req request.UpdateProfile

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	userID, ok := gateway.GetUserIdFromContext(c)
	if !ok {
		return
	}

	_, err := u.accountsClient.UpdateProfile(context.Background(), &pb.UpdateProfileRequest{
		UserId:    userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Address:   req.Address,
		Pincode:   req.Pincode,
	})
	if err == nil {
		c.JSON(200, response.SM{
			Status: mystatus.Success,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

// GetProfile
func (u *UserAccountHandler) GetProfile(c *gin.Context) { //unimplemented
	userID, ok := gateway.GetUserIdFromContext(c)
	if !ok {
		return
	}

	resp, err := u.accountsClient.GetProfile(context.Background(), &pb.GetProfileRequest{
		UserId: userID,
	})
	if err == nil {
		c.JSON(200, response.GetProfileResponse{
			Status: mystatus.Success,
			Profile: response.Profile{
				FirstName:   resp.FirstName,
				LastName:    resp.LastName,
				Email:       resp.Email,
				Address:     resp.Address,
				PhoneNumber: resp.PhoneNumber,
				Pincode:     resp.Pincode,
			},
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

// UpdatePassword
func (u *UserAccountHandler) UpdatePasswordUsingOldPw(c *gin.Context) {
	var req request.UpdatePasswordUsingOldPw

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	userID, ok := gateway.GetUserIdFromContext(c)
	if !ok {
		return
	}

	_, err := u.accountsClient.UpdatePasswordUsingOldPw(context.Background(), &pb.UpdatePasswordUsingOldPwRequest{
		UserId:      userID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err == nil {
		c.JSON(200, response.SM{
			Status: mystatus.Success,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

// SetNewPwAfterVerifyingOTP
func (u *UserAccountHandler) SetNewPwAfterVerifyingOTP(c *gin.Context) {
	purpose := c.GetString(tag.CtxPurpose)
	purposeStatus := c.GetString(tag.CtxPurposeStatus)
	if purpose == "" || purpose != tag.PwChange {
		c.JSON(400, response.SRE{
			Status:       mystatus.Failed,
			Error:        "no '" + tag.PwChange + "' purpose in token",
			ResponseCode: respCode.NotEnoughPermissionsInToken,
		})
		return
	}
	if purposeStatus == "expired" {
		c.JSON(400, response.SRE{
			Status:       mystatus.Failed,
			Error:        "Purpose '" + tag.PwChange + "' expired in token",
			ResponseCode: respCode.NotEnoughPermissionsInToken,
		})
		return
	}

	var req request.SettingNewPassword
	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	userID, ok := gateway.GetUserIdFromContext(c)
	if !ok {
		return
	}

	_, err := u.accountsClient.SetNewPwAfterVerifyingOTP(context.Background(), &pb.SetNewPwAfterVerifyingOTPRequest{
		UserId:      userID,
		NewPassword: req.NewPassword,
	})
	if err == nil {
		c.JSON(200, response.SM{
			Status: mystatus.Success,
		})
	} else {
		fmt.Println("error retuurned from grpc server in setting new password: ", err)
		gateway.HandleGrpcStatus(c, err)
	}
}
