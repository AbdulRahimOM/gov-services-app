package ksebhanlder

import (
	"context"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gin-gonic/gin"
)

type KSEBHandler struct {
	accClient         pb.KSEBAdminAccServiceClient
	agencyAdminClient pb.KSEBAgencyAdminServiceClient
	ksebChatClient    pb.KsebChatServiceClient
}

func NewKsebHandler(accClient pb.KSEBAdminAccServiceClient, agencyAdminClient pb.KSEBAgencyAdminServiceClient, ksebChatClient pb.KsebChatServiceClient) *KSEBHandler {
	return &KSEBHandler{
		accClient:         accClient,
		agencyAdminClient: agencyAdminClient,
		ksebChatClient:    ksebChatClient,
	}
}

// KSEBRegisterSectionCode
func (kseb *KSEBHandler) KSEBRegisterSectionCode(c *gin.Context) {
	var req requests.KsebRegSectionCode
	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	adminId, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	_, err := kseb.accClient.RegisterSectionCode(context.Background(), &pb.RegisterSectionCodeRequest{
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

// KSEBGetComplaints
func (kseb *KSEBHandler) AdminGetComplaints(c *gin.Context) {
	status := c.DefaultQuery("status", "all")       //all, opened, closed, not-opened
	attenderScope := c.DefaultQuery("scope", "all") //all, me-only

	adminId, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	resp, err := kseb.agencyAdminClient.GetComplaints(context.Background(), &pb.GetComplaintsRequest{
		AdminId: adminId,
		SearchCriteria: &pb.KsebComplaintSearchCriteria{
			Status:        status,
			AttenderScope: attenderScope,
		},
	})
	if err == nil {
		complaints := make([]commondto.KsebComplaintResponse, len(resp.Complaints))
		for i, complaint := range resp.Complaints {
			complaints[i] = commondto.KsebComplaintResponse{
				ID:             complaint.ID,
				UserID:         complaint.UserID,
				Type:           complaint.Type,
				Title:          complaint.Title,
				Description:    complaint.Description,
				ConsumerNumber: complaint.ConsumerNumber,
				AttenderID:     complaint.AttenderID,
				Status:         complaint.Status,
				CreatedAt:      complaint.CreatedAt,
				Remarks:        complaint.Remarks,
				ClosedAt:       complaint.ClosedAt,
			}
		}
		c.JSON(200, response.GetKsebComplaints{
			Status:     mystatus.Success,
			Complaints: complaints,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

// AdminOpenComplaint
func (kseb *KSEBHandler) AdminOpenComplaint(c *gin.Context) {
	complaintId, ok := gateway.HandleGetUrlParamsInt32(c, "complaintId")
	if !ok {
		return
	}

	adminId, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	_, err := kseb.agencyAdminClient.OpenComplaint(context.Background(), &pb.OpenComplaintRequest{
		AdminId:     adminId,
		ComplaintId: complaintId,
	})

	if err == nil {
		c.JSON(200, response.SM{
			Status: mystatus.Success,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

// AdminCloseComplaint
func (kseb *KSEBHandler) AdminCloseComplaint(c *gin.Context) {
	var req requests.KsebCloseComplaint
	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	adminId, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	_, err := kseb.agencyAdminClient.CloseComplaint(context.Background(), &pb.CloseComplaintRequest{
		AdminId:     adminId,
		ComplaintId: req.ComplaintId,
		Remarks:     req.Remarks,
	})

	if err == nil {
		c.JSON(200, response.SM{
			Status: mystatus.Success,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}
