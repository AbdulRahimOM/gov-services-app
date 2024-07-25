package appointments

import (
	"context"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	gateway "github.com/AbdulRahimOM/gov-services-app/internal/gateway/fiber"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gofiber/fiber/v2"
)

type AppointmentHandler struct {
	appointmentsClient pb.AppointmentServiceClient
}

func NewAppointmentHandler(client pb.AppointmentServiceClient) *AppointmentHandler {
	return &AppointmentHandler{
		appointmentsClient: client,
	}
}

func (h *AppointmentHandler) AppointAttender(c *fiber.Ctx) error {
	var req requests.Appointee
	if err := gateway.BindAndValidateRequestFiber(c, &req); err != nil {
		return err
	}

	adminId, err := gateway.GetAdminIdFromContextFiber(c)
	if err != nil {
		return err
	}

	_, err = h.appointmentsClient.AppointAttender(context.Background(), &pb.AttenderAppointmentRequest{
		Appointer: &pb.Appointer{
			Id: adminId,
		},
		Appointee: &pb.Appointee{
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			Email:       req.Email,
			PhoneNumber: req.PhoneNumber,
		},
	})

	if err == nil {
		return c.Status(200).JSON(response.SM{
			Status: mystatus.Success,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (h *AppointmentHandler) CreateChildOffice(c *fiber.Ctx) error {
	var req requests.ProposedOffice
	if err := gateway.BindAndValidateRequestFiber(c, &req); err != nil {
		return err
	}

	adminId, err := gateway.GetAdminIdFromContextFiber(c)
	if err != nil {
		return err
	}

	resp, err := h.appointmentsClient.CreateChildOffice(context.Background(), &pb.CreateChildOfficeRequest{
		AdminID: adminId,
		ProposedChildOffice: &pb.ProposedChildOffice{
			Name:    req.Name,
			Address: req.Address,
		},
	})

	if err == nil {
		return c.Status(200).JSON(response.CreateChildOffice{
			Status:        mystatus.Success,
			ChildOfficeID: resp.ChildOfficeID,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (h *AppointmentHandler) AppointChildOfficeHead(c *fiber.Ctx) error {
	var req requests.AppointChildOfficeHead
	if err := gateway.BindAndValidateRequestFiber(c, &req); err != nil {
		return err
	}

	adminId, err := gateway.GetAdminIdFromContextFiber(c)
	if err != nil {
		return err
	}

	_, err = h.appointmentsClient.AppointChildOfficeHead(context.Background(), &pb.OfficeHeadAppointmentRequest{
		Appointer: &pb.Appointer{
			Id: adminId,
		},
		Appointee: &pb.Appointee{
			FirstName:   req.Appointee.FirstName,
			LastName:    req.Appointee.LastName,
			Email:       req.Appointee.Email,
			PhoneNumber: req.Appointee.PhoneNumber,
		},
		ChildOfficeID: req.ChildOfficeID,
	})

	if err == nil {
		return c.Status(200).JSON(response.SM{
			Status: mystatus.Success,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (h *AppointmentHandler) AppointChildOfficeDeputyHead(c *fiber.Ctx) error {
	var req requests.AppointChildOfficeDeputyHead
	if err := gateway.BindAndValidateRequestFiber(c, &req); err != nil {
		return err
	}

	adminId, err := gateway.GetAdminIdFromContextFiber(c)
	if err != nil {
		return err
	}

	_, err = h.appointmentsClient.AppointChildOfficeDeputyHead(context.Background(), &pb.OfficeHeadAppointmentRequest{
		Appointer: &pb.Appointer{
			Id: adminId,
		},
		Appointee: &pb.Appointee{
			FirstName:   req.Appointee.FirstName,
			LastName:    req.Appointee.LastName,
			Email:       req.Appointee.Email,
			PhoneNumber: req.Appointee.PhoneNumber,
		},
		ChildOfficeID: req.ChildOfficeID,
	})

	if err == nil {
		return c.Status(200).JSON(response.SM{
			Status: mystatus.Success,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}
