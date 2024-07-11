package acchandler

import (
	response "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gin-gonic/gin"
)



// AdminGetAdmins
func (h *AdminAccountHandler) AdminGetAdmins(c *gin.Context) {

	adminID, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	searchCriteria := pb.AdminSearchCriteria{
		FirstName:   c.DefaultQuery("fName", ""),
		LastName:    c.DefaultQuery("lName", ""),
		Email:       c.DefaultQuery("email", ""),
		PhoneNumber: c.DefaultQuery("phoneNumber", ""),
		Designation: c.DefaultQuery("designation", ""),
	}

	searchCriteria.OfficeId, ok = gateway.HandleGetQueryParamsInt32(c, "officeId")
	if !ok {
		return
	}
	// searchCriteria.RankId, ok = gateway.HandleGetQueryParamsInt32(c, "rankId")
	// if !ok {
	// 	return
	// }
	// searchCriteria.PostId, ok = gateway.HandleGetQueryParamsInt32(c, "postId")
	// if !ok {
	// 	return
	// }

	resp, err := h.accountsClient.AdminGetAdmins(c, &pb.AdminGetAdminsRequest{
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
				OfficeId:   admin.OfficeId,
				Designation: admin.Designation,
			})
		}
		c.JSON(200, response.AdminGetAdminsResponse{
			Status: mystatus.Success,
			Admins: admins,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}
