package appointments

import (
	"context"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	appointmentsClient pb.AppointmentServiceClient
}

func NewAppointmentHandler(client pb.AppointmentServiceClient) *AppointmentHandler {
	return &AppointmentHandler{
		appointmentsClient: client,
	}
}

func (h *AppointmentHandler) AppointAttender(c *gin.Context) {
	var req requests.Appointee
	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	adminId, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	_, err := h.appointmentsClient.AppointAttender(context.Background(), &pb.AttenderAppointmentRequest{
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
		c.JSON(200, response.SM{
			Status: mystatus.Success,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

func (h *AppointmentHandler) CreateChildOffice(c *gin.Context) {
	var req requests.ProposedOffice
	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	adminId, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	resp, err := h.appointmentsClient.CreateChildOffice(context.Background(), &pb.CreateChildOfficeRequest{
		AdminID: adminId,
		ProposedChildOffice: &pb.ProposedChildOffice{
			Name:    req.Name,
			Address: req.Address,
		},
	})

	if err == nil {
		c.JSON(200, response.CreateChildOffice{
			Status: mystatus.Success,
			ChildOfficeID: resp.ChildOfficeID,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

// AppointChildOfficeHead
func (h *AppointmentHandler) AppointChildOfficeHead(c *gin.Context) {
	var req requests.AppointChildOfficeHead
	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	adminId, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	_, err := h.appointmentsClient.AppointChildOfficeHead(context.Background(), &pb.OfficeHeadAppointmentRequest{
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
		c.JSON(200, response.SM{
			Status: mystatus.Success,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

// AppointChildOfficeDeputyHead
func (h *AppointmentHandler) AppointChildOfficeDeputyHead(c *gin.Context) {
	var req requests.AppointChildOfficeDeputyHead
	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	adminId, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	_, err := h.appointmentsClient.AppointChildOfficeDeputyHead(context.Background(), &pb.OfficeHeadAppointmentRequest{
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
		c.JSON(200, response.SM{
			Status: mystatus.Success,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}