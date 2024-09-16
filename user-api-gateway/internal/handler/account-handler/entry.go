package acchandler

import (
	gateway "github.com/AbdulRahimOM/gov-services-app/internal/gateway/fiber"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/request"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/response"
	"github.com/gofiber/fiber/v2"
)

func (u *UserAccountHandler) Ping(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "pong",
	})
}

func (u *UserAccountHandler) RequestOTPForLogin(c *fiber.Ctx) error {
	var req request.UserLoginGetOTP

	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	_, err := u.accountsClient.UserLoginGetOTP(c.Context(), &pb.UserLoginGetOTPRequest{
		PhoneNumber: req.PhoneNumber,
	})
	if err == nil {
		return c.Status(200).JSON(response.GetOTPResponse{
			Status:      mystatus.Success,
			Last4Digits: req.PhoneNumber[len(req.PhoneNumber)-4:],
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (u *UserAccountHandler) UserLoginVerifyOTP(c *fiber.Ctx) error {
	var req request.UserLoginVerifyOTP

	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	resp, err := u.accountsClient.UserLoginVerifyOTP(c.Context(), &pb.UserLoginVerifyOTPRequest{
		PhoneNumber: req.PhoneNumber,
		Otp:         req.OTP,
	})

	if err == nil {
		return c.Status(200).JSON(response.UserLogin{
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
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (u *UserAccountHandler) RequestOTPForSignUp(c *fiber.Ctx) error {
	var req request.GetOTPForSignup

	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	_, err := u.accountsClient.UserSignUpGetOTP(c.Context(), &pb.UserSignUpGetOTPRequest{
		PhoneNumber: req.PhoneNumber,
	})
	if err == nil {
		return c.Status(200).JSON(response.GetOTPResponse{
			Status:      mystatus.Success,
			Last4Digits: req.PhoneNumber[len(req.PhoneNumber)-4:],
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (u *UserAccountHandler) SubmitOTPForSignUp(c *fiber.Ctx) error {
	var req request.UserSignpViaOTP

	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	resp, err := u.accountsClient.UserSignUpVerifyOTP(c.Context(), &pb.UserSignUpVerifyOTPRequest{
		PhoneNumber: req.PhoneNumber,
		Otp:         req.OTP,
	})
	if err == nil {
		return c.Status(200).JSON(response.UserSignUp{
			Status:  mystatus.Success,
			Message: "User signed up successfully",
			User: response.PreliminaryUserData{
				Id:          resp.UserDetails.Id,
				PhoneNumber: resp.UserDetails.PhoneNumber,
			},
			Token: resp.Token,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (u *UserAccountHandler) UserLoginViaPassword(c *fiber.Ctx) error {
	var req request.UserLoginViaPassword

	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}
	resp, err := u.accountsClient.UserLoginViaPassword(c.Context(), &pb.UserLoginViaPasswordRequest{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})
	if err == nil {
		return c.Status(200).JSON(response.UserLogin{
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
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}
