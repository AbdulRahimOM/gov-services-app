package acchandler

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/request"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	gateway "github.com/AbdulRahimOM/gov-services-app/internal/gateway/fiber"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gofiber/fiber/v2"
)

func (u *AdminAccountHandler) Ping(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "pong",
	})
}

func (u *AdminAccountHandler) AdminLoginViaPassword(c *fiber.Ctx) error {
	var req request.AdminLoginViaPassword

	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}
	var resp *pb.AdminLoginResponse
	err := u.circuitBreaker.Run(func() error {
		var err error
		resp, err = u.accountsClient.AdminLoginViaPassword(c.Context(), &pb.AdminLoginViaPasswordRequest{
			Username: req.Username,
			Password: req.Password,
		})
		return err
	})
	if err == nil {
		return c.Status(200).JSON(response.AdminLogin{
			Status: mystatus.Success,
			AdminData: response.AdminBasicData{
				Id:          resp.AdminDetails.Id,
				FirstName:   resp.AdminDetails.FirstName,
				LastName:    resp.AdminDetails.LastName,
				PhoneNumber: resp.AdminDetails.PhoneNumber,
				// DeptID:      resp.AdminDetails.DeptId,
				// RankID:      resp.AdminDetails.RankId,
				Designation: resp.AdminDetails.Designation,
			},
			Token: resp.Token,
		})
	}
	return gateway.HandleGrpcStatusFiber(c, err)
}
