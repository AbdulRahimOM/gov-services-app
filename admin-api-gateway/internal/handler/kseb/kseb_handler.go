package ksebhanlder

import (
	"context"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gin-gonic/gin"
)

type KSEBHandler struct {
	ksebClient pb.KSEBAdminAccServiceClient
}

func NewAppointmentHandler(client pb.KSEBAdminAccServiceClient) *KSEBHandler {
	return &KSEBHandler{
		ksebClient: client,
	}
}

// KSEBRegisterSectionCode
func (h *KSEBHandler) KSEBRegisterSectionCode(c *gin.Context) {
	var req requests.KsebRegSectionCode
	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	adminId, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	_, err := h.ksebClient.RegisterSectionCode(context.Background(), &pb.RegisterSectionCodeRequest{
		AdminId:     adminId,
		SectionCode: req.SectionCode,
		OfficeId:    req.OfficeId,
	})

	if err == nil {
		c.JSON(200, response.SM{
			Status: mystatus.Success,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}

}
