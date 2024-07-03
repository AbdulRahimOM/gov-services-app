package handler

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gin-gonic/gin"
)

// AdminGetOffices
func (h *AdminAccountHandler) AdminGetOffices(c *gin.Context) {
	resp, err := h.accountsClient.AdminGetOffices(c, nil)

	if err == nil {
		var offices []*commondto.Office
		for _, office := range resp.Office {
			offices = append(offices, &commondto.Office{
				ID:               office.Id,
				DeptID:           office.DeptId,
				HierarchyRank:    office.HierarchyRank,
				RegionName:       office.RegionName,
				HeadOfficerID:    office.HeadOfficerId,
				OfficeLocation:   office.OfficeLocation,
				SuperiorOfficeID: office.SuperiorOfficeId,
			})
		}
		c.JSON(200, response.AdminGetOffices{
			Status:  mystatus.Success,
			Offices: offices,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}
