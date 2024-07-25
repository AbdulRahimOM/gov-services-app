package acchandler

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	gateway "github.com/AbdulRahimOM/gov-services-app/internal/gateway/fiber"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gofiber/fiber/v2"
)

func (h *AdminAccountHandler) AdminGetOffices(c *fiber.Ctx) error {
	searchCriteria := pb.OfficeSearchCriteria{
		Name:    c.Query("name"),
		Address: c.Query("address"),
	}
	var err error
	searchCriteria.Id, err = gateway.HandleGetQueryParamsInt32Fiber(c, "id")
	if err != nil {
		return err
	}
	searchCriteria.DeptID, err = gateway.HandleGetQueryParamsInt32Fiber(c, "deptId")
	if err != nil {
		return err
	}
	searchCriteria.Rank, err = gateway.HandleGetQueryParamsInt32Fiber(c, "rank")
	if err != nil {
		return err
	}
	searchCriteria.SuperiorOfficeID, err = gateway.HandleGetQueryParamsInt32Fiber(c, "superiorOfficeId")
	if err != nil {
		return err
	}
	resp, err := h.accountsClient.AdminGetOffices(c.Context(), &searchCriteria)

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
		return c.Status(200).JSON(response.AdminGetOffices{
			Status:  mystatus.Success,
			Offices: offices,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}
