package acchandler

import (
	"context"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/request"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	gateway "github.com/AbdulRahimOM/gov-services-app/internal/gateway/fiber"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gofiber/fiber/v2"
)

func (u *AdminAccountHandler) AdminGetProfile(c *fiber.Ctx) error {
	adminID, ok, err := gateway.GetAdminIdFromContextFiber(c)
	if !ok {
		return err
	}

	var resp *pb.AdminGetProfileResponse
	err = u.circuitBreaker.Run(func() error {
		var err error
		resp, err = u.accountsClient.AdminGetProfile(context.Background(), &pb.AdminGetProfileRequest{
			AdminId: adminID,
		})
		return err
	})

	if err == nil {
		return c.Status(200).JSON(response.AdminGetProfileResponse{
			Status: mystatus.Success,
			Profile: response.Profile{
				FirstName:   resp.FirstName,
				LastName:    resp.LastName,
				Username:    resp.Username,
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

func (u *AdminAccountHandler) AdminUpdatePasswordUsingOldPw(c *fiber.Ctx) error {
	var req request.AdminUpdatePasswordUsingOldPw

	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	adminID, ok, err := gateway.GetAdminIdFromContextFiber(c)
	if !ok {
		return err
	}

	_, err = u.accountsClient.AdminUpdatePasswordUsingOldPw(c.Context(), &pb.AdminUpdatePasswordUsingOldPwRequest{
		AdminId:     adminID,
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

func (u *AdminAccountHandler) AdminUpdateProfile(c *fiber.Ctx) error {
	var req request.AdminUpdateProfile

	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	adminID, ok, err := gateway.GetAdminIdFromContextFiber(c)
	if !ok {
		return err
	}

	_, err = u.accountsClient.AdminUpdateProfile(context.Background(), &pb.AdminUpdateProfileRequest{
		AdminId:     adminID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Address:     req.Address,
		Pincode:     req.Pincode,
		PhoneNumber: req.PhoneNumber,
	})
	if err == nil {
		return c.Status(200).JSON(response.SM{
			Status: mystatus.Success,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}
