package acchandler

import (
	"fmt"
	"time"

	gateway "github.com/AbdulRahimOM/gov-services-app/internal/gateway/fiber"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	respCode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	"github.com/AbdulRahimOM/gov-services-app/internal/tag"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/request"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/response"
	"github.com/gofiber/fiber/v2"
)

func (u *UserAccountHandler) UserGetOTPForPwChange(c *fiber.Ctx) error {
	userID, ok, err := gateway.GetUserIdFromContextFiber(c)
	if !ok {
		return err
	}

	resp, err := u.accountsClient.UserGetOTPForPwChange(c.Context(), &pb.UserGetOTPForPwChangeRequest{
		UserId: userID,
	})
	if err == nil {
		return c.Status(200).JSON(response.GetOTPResponse{
			Status:      mystatus.Success,
			Last4Digits: resp.Last4Digits,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (u *UserAccountHandler) UserVerifyOTPForPwChange(c *fiber.Ctx) error {
	var req request.UserVerifyOTPForPwChange
	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	userID, ok, err := gateway.GetUserIdFromContextFiber(c)
	if !ok {
		return err
	}

	resp, err := u.accountsClient.UserVerifyOTPForPwChange(c.Context(), &pb.UserVerifyOTPForPwChangeRequest{
		UserId: userID,
		Otp:    req.Otp,
	})
	if err == nil {
		return c.Status(200).JSON(response.UserVerifyOTPForPwChangeResponse{
			Status:    mystatus.Success,
			Msg:       "OTP verified",
			TempToken: resp.TempToken,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (u *UserAccountHandler) SignedUpUserSettingPw(c *fiber.Ctx) error {
	var req request.SettingNewPassword

	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	userID, ok, err := gateway.GetUserIdFromContextFiber(c)
	if !ok {
		return err
	}

	resp, err := u.accountsClient.SignedUpUserSettingPw(c.Context(), &pb.SignedUpUserSettingPwRequest{
		UserId:      userID,
		NewPassword: req.NewPassword,
	})
	if err == nil {
		return c.Status(200).JSON(response.UpdateToken{
			Status: mystatus.Success,
			Token:  resp.Token,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (u *UserAccountHandler) UserUpdateProfile(c *fiber.Ctx) error {
	var req request.UserUpdateProfile

	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	userID, ok, err := gateway.GetUserIdFromContextFiber(c)
	if !ok {
		return err
	}

	_, err = u.accountsClient.UserUpdateProfile(c.Context(), &pb.UserUpdateProfileRequest{
		UserId:    userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Address:   req.Address,
		Pincode:   req.Pincode,
	})
	if err == nil {
		return c.Status(200).JSON(response.SM{
			Status: mystatus.Success,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (u *UserAccountHandler) UserGetProfile(c *fiber.Ctx) error {
	userID, ok, err := gateway.GetUserIdFromContextFiber(c)
	if !ok {
		return err
	}

	resp, err := u.accountsClient.UserGetProfile(c.Context(), &pb.UserGetProfileRequest{
		UserId: userID,
	})
	if err == nil {
		return c.Status(200).JSON(response.UserGetProfileResponse{
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
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (u *UserAccountHandler) UserUpdatePasswordUsingOldPw(c *fiber.Ctx) error {
	var req request.UserUpdatePasswordUsingOldPw

	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	userID, ok, err := gateway.GetUserIdFromContextFiber(c)
	if !ok {
		return err
	}

	_, err = u.accountsClient.UserUpdatePasswordUsingOldPw(c.Context(), &pb.UserUpdatePasswordUsingOldPwRequest{
		UserId:      userID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err == nil {
		return c.Status(200).JSON(response.SM{
			Status: mystatus.Success,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (u *UserAccountHandler) UserSetNewPwAfterVerifyingOTP(c *fiber.Ctx) error {
	purpose, ok := c.Locals(tag.CtxPurpose).(string)
	if !ok {
		return c.Status(400).JSON(response.SRE{
			Status:       mystatus.Failed,
			Error:        "no '" + tag.CtxPurpose + "' in context",
			ResponseCode: respCode.NotEnoughPermissionsInToken,
		})
	}

	purposeExpiryTime, ok := c.Locals(tag.CtxPurposeExpiryTime).(time.Time)
	if !ok {
		return c.Status(400).JSON(response.SRE{
			Status:       mystatus.Failed,
			Error:        "no '" + tag.CtxPurposeExpiryTime + "' in context",
			ResponseCode: respCode.NotEnoughPermissionsInToken,
		})
	}

	if purpose == "" || purpose != tag.PwChange {
		return c.Status(400).JSON(response.SRE{
			Status:       mystatus.Failed,
			Error:        "no '" + tag.PwChange + "' purpose in token",
			ResponseCode: respCode.NotEnoughPermissionsInToken,
		})
	}

	if time.Now().After(purposeExpiryTime) {
		return c.Status(400).JSON(response.SRE{
			Status:       mystatus.Failed,
			Error:        "Purpose '" + tag.PwChange + "' expired in token",
			ResponseCode: respCode.NotEnoughPermissionsInToken,
		})
	}

	var req request.SettingNewPassword
	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	userID, ok, err := gateway.GetUserIdFromContextFiber(c)
	if !ok {
		return err
	}

	_, err = u.accountsClient.UserSetNewPwAfterVerifyingOTP(c.Context(), &pb.UserSetNewPwAfterVerifyingOTPRequest{
		UserId:      userID,
		NewPassword: req.NewPassword,
	})
	if err == nil {
		return c.Status(200).JSON(response.SM{
			Status: mystatus.Success,
		})
	} else {
		fmt.Println("error retuurned from grpc server in setting new password: ", err)
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}
