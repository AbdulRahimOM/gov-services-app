package appointments

import (
	"context"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway/fiber"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/eapache/go-resiliency/breaker"
	"github.com/gofiber/fiber/v2"
)

type AppointmentHandler struct {
	appointmentsClient pb.AppointmentServiceClient
	circuitBreaker     *breaker.Breaker
}

func NewAppointmentHandler(client pb.AppointmentServiceClient, circuitBreaker *breaker.Breaker) *AppointmentHandler {
	return &AppointmentHandler{
		appointmentsClient: client,
		circuitBreaker:     circuitBreaker,
	}
}

func (h *AppointmentHandler) AppointAttender(c *fiber.Ctx) error {
	var req requests.Appointee
	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	adminId, ok, err := gateway.GetAdminIdFromContextFiber(c)
	if !ok {
		return err
	}

	err = h.circuitBreaker.Run(func() error {
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
		return err
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
	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	adminId, ok, err := gateway.GetAdminIdFromContextFiber(c)
	if !ok {
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
	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	adminId, ok, err := gateway.GetAdminIdFromContextFiber(c)
	if !ok {
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
	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	adminId, ok, err := gateway.GetAdminIdFromContextFiber(c)
	if !ok {
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
