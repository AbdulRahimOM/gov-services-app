package handler

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/request"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gin-gonic/gin"
)

// AdminAddDept
func (h *AdminAccountHandler) AdminAddDept(c *gin.Context) {
	adminID, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	var req request.AdminAddDept
	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	resp, err := h.accountsClient.AdminAddDept(c, &pb.AdminAddDeptRequest{
		AdminId: adminID,
		NewDept: &pb.NewDept{
			Name:        req.DeptName,
			Description: req.DeptDescription,
		},
	})
	if err == nil {
		c.JSON(200, response.AdminAddDept{
			Status:      mystatus.Success,
			AddedDeptID: resp.NewDeptID,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}
