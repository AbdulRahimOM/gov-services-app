package acchandler

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/request"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	gateway "github.com/AbdulRahimOM/gov-services-app/internal/gateway/fiber"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gofiber/fiber/v2"
)

//	func (u *AdminAccountHandler) Ping(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	}
func (u *AdminAccountHandler) Ping(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "pong",
	})
}

// func (u *AdminAccountHandler) AdminLoginViaPassword(c *gin.Context) {
// 	var req request.AdminLoginViaPassword

// 	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
// 		return
// 	}
// 	resp, err := u.accountsClient.AdminLoginViaPassword(c, &pb.AdminLoginViaPasswordRequest{
// 		Username: req.Username,
// 		Password: req.Password,
// 	})
// 	if err == nil {
// 		c.JSON(200, response.AdminLogin{
// 			Status: mystatus.Success,
// 			AdminData: response.AdminBasicData{
// 				Id:          resp.AdminDetails.Id,
// 				FirstName:   resp.AdminDetails.FirstName,
// 				LastName:    resp.AdminDetails.LastName,
// 				PhoneNumber: resp.AdminDetails.PhoneNumber,
// 				// DeptID:      resp.AdminDetails.DeptId,
// 				// RankID:      resp.AdminDetails.RankId,
// 				Designation: resp.AdminDetails.Designation,

//				},
//				Token: resp.Token,
//			})
//		} else {
//			gateway.HandleGrpcStatus(c, err)
//		}
//	}
func (u *AdminAccountHandler) AdminLoginViaPassword(c *fiber.Ctx) error {
	var req request.AdminLoginViaPassword

	if err := gateway.BindAndValidateRequestFiber(c, &req); err != nil {
		return err
	}
	resp, err := u.accountsClient.AdminLoginViaPassword(c.Context(), &pb.AdminLoginViaPasswordRequest{
		Username: req.Username,
		Password: req.Password,
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
	gateway.HandleGrpcStatusFiber(c, err)
	return nil
}
