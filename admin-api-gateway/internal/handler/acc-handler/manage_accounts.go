package acchandler

import (
	response "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	gateway "github.com/AbdulRahimOM/gov-services-app/internal/gateway/fiber"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gofiber/fiber/v2"
)

func (h *AdminAccountHandler) AdminGetAdmins(c *fiber.Ctx) error {

	adminID, ok,err := gateway.GetAdminIdFromContextFiber(c)
	if !ok {
		return err
	}

	searchCriteria := pb.AdminSearchCriteria{
		FirstName:   c.Query("fName"),
		LastName:    c.Query("lName"),
		Email:       c.Query("email"),
		PhoneNumber: c.Query("phoneNumber"),
		Designation: c.Query("designation"),
	}

	searchCriteria.OfficeId, ok, err = gateway.HandleGetQueryParamsInt32Fiber(c, "officeId")
	if !ok {
		return err
	}

	// searchCriteria.RankId, err = gateway.HandleGetQueryParamsInt32Fiber(c, "rankId")
	// if err != nil {
	// 	return err
	// }
	// searchCriteria.PostId, err = gateway.HandleGetQueryParamsInt32Fiber(c, "postId")
	// if err != nil {
	// 	return err
	// }

	resp, err := h.accountsClient.AdminGetAdmins(c.Context(), &pb.AdminGetAdminsRequest{
		AdminId:        adminID,
		SearchCriteria: &searchCriteria,
	})

	if err == nil {
		var admins []*commondto.Admin
		for _, admin := range resp.Admin {
			admins = append(admins, &commondto.Admin{
				ID:          admin.Id,
				FirstName:   admin.FirstName,
				LastName:    admin.LastName,
				Email:       admin.Email,
				Address:     admin.Address,
				Pincode:     admin.Pincode,
				PhoneNumber: admin.PhoneNumber,
				OfficeId:    admin.OfficeId,
				Designation: admin.Designation,
			})
		}
		return c.Status(200).JSON(response.AdminGetAdminsResponse{
			Status: mystatus.Success,
			Admins: admins,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}
