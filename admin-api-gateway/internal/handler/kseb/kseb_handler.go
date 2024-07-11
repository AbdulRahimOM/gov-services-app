package ksebhanlder

import (
	"context"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway"
	ksebpb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated/ksebpb"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gin-gonic/gin"
)

type KSEBHandler struct {
	ksebClient ksebpb.KSEBServiceClient
}

func NewAppointmentHandler(client ksebpb.KSEBServiceClient) *KSEBHandler {
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

	_, err := h.ksebClient.RegisterSectionCode(context.Background(), &ksebpb.RegisterSectionCodeRequest{
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
