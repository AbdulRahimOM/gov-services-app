package handler

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/request"
	response "github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gin-gonic/gin"
)

// // AdminGetDepts
// func (h *AdminAccountHandler) AdminGetDepts(c *gin.Context) {
// 	resp, err := h.accountsClient.AdminGetDepts(c, nil)

// 	if err == nil {
// 		var depts []*commondto.Department
// 		for _, dept := range resp.Dept {
// 			depts = append(depts, &commondto.Department{
// 				ID:          dept.Id,
// 				Name:        dept.Name,
// 				Description: dept.Description,
// 			})
// 		}
// 		c.JSON(200, response.AdminGetDepts{
// 			Status: mystatus.Success,
// 			Depts:  depts,
// 		})
// 	} else {
// 		gateway.HandleGrpcStatus(c, err)
// 	}
// }

// AdminGetAddSubAdminPage
func (h *AdminAccountHandler) AdminAddSubAdmin(c *gin.Context) {
	adminID, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	var req request.AdminAddSubAdmin
	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	resp, err := h.accountsClient.AdminAddSubAdmin(c, &pb.AdminAddSubAdminRequest{
		AdminId: adminID,
		NewSubAdmin: &pb.NewSubAdmin{
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			Email:       req.Email,
			PhoneNumber: req.PhoneNumber,
			DeptId:      req.DeptID,
			Designation: req.Designation,
			RankId:      req.RankID,
			OfficeId:    req.OfficeID,
		},
	})
	if err == nil {
		c.JSON(200, response.AdminAddSubAdmin{
			Status:       mystatus.Success,
			AddedAdminID: resp.NewSubAdminID,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

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

	searchCriteria.DeptId, ok = gateway.HandleGetQueryParamsInt32(c, "deptId")
	if !ok {
		return
	}
	searchCriteria.RankId, ok = gateway.HandleGetQueryParamsInt32(c, "rankId")
	if !ok {
		return
	}

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
				DeptID:      admin.DeptId,
				RankID:      admin.RankId,
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
