package acchandler

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gin-gonic/gin"
)

// AdminGetOffices
func (h *AdminAccountHandler) AdminGetOffices(c *gin.Context) {
	searchCriteria := pb.OfficeSearchCriteria{
		Name:    c.DefaultQuery("name", ""),
		Address: c.DefaultQuery("address", ""),
	}
	var ok bool
	searchCriteria.Id, ok = gateway.HandleGetQueryParamsInt32(c, "id")
	if !ok {
		return
	}
	searchCriteria.DeptID, ok = gateway.HandleGetQueryParamsInt32(c, "deptId")
	if !ok {
		return
	}
	searchCriteria.Rank, ok = gateway.HandleGetQueryParamsInt32(c, "rank")
	if !ok {
		return
	}
	searchCriteria.SuperiorOfficeID, ok = gateway.HandleGetQueryParamsInt32(c, "superiorOfficeId")
	if !ok {
		return
	}
	resp, err := h.accountsClient.AdminGetOffices(c, &searchCriteria)

	if err == nil {
		var offices []*commondto.Office
		for _, office := range resp.Office {
			offices = append(offices, &commondto.Office{
				ID:               office.Id,
				Name:             office.Name,
				DeptID:           office.DeptId,
				Rank:             office.Rank,
				Address:          office.Address,
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
