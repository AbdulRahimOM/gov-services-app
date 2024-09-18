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
	var ok bool
	if searchCriteria.Id, ok, err = gateway.HandleGetQueryParamsInt32Fiber(c, "id"); !ok {
		return err
	}
	if searchCriteria.DeptID, ok, err = gateway.HandleGetQueryParamsInt32Fiber(c, "deptId"); !ok {
		return err
	}
	if searchCriteria.Rank, ok, err = gateway.HandleGetQueryParamsInt32Fiber(c, "rank"); !ok {
		return err
	}
	if searchCriteria.SuperiorOfficeID, ok, err = gateway.HandleGetQueryParamsInt32Fiber(c, "superiorOfficeId"); !ok {
		return err
	}

	var resp *pb.AdminGetOfficesResponse
	err = h.circuitBreaker.Run(func() error {
		var err error
		resp, err = h.accountsClient.AdminGetOffices(c.Context(), &searchCriteria)

		return err
	})

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
