package handler

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

	searchCriteria:= pb.AdminSearchCriteria{
		FirstName: c.DefaultQuery("fName",""),
		LastName: c.DefaultQuery("lName",""),
		Email: c.DefaultQuery("email",""),
		PhoneNumber: c.DefaultQuery("phoneNumber",""),
		Designation: c.DefaultQuery("designation",""),
	}

	searchCriteria.DeptId, ok = gateway.HandleGetQueryParamsInt32(c, "deptId")
	if !ok {
		return
	}
	searchCriteria.RankId,ok = gateway.HandleGetQueryParamsInt32(c, "rankId")
	if !ok {
		return
	}

	resp, err := h.accountsClient.AdminGetAdmins(c, &pb.AdminGetAdminsRequest{
		AdminId: adminID,
		SearchCriteria: &searchCriteria,
	})
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
			DeptID:      admin.DeptId,
			RankID:      admin.RankId,
			Designation: admin.Designation,
		})
	}

	if err == nil {
		c.JSON(200, response.AdminGetAdminsResponse{
			Status: mystatus.Success,
			Admins: admins,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}
