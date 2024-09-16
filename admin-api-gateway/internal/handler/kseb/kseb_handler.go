package ksebhanlder

import (
	"context"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	gateway "github.com/AbdulRahimOM/gov-services-app/internal/gateway/fiber"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gofiber/fiber/v2"
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
func (kseb *KSEBHandler) KSEBRegisterSectionCode(c *fiber.Ctx) error {
	var req requests.KsebRegSectionCode
	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	adminId, ok, err := gateway.GetAdminIdFromContextFiber(c)
	if !ok {
		return err
	}

	_, err = kseb.accClient.RegisterSectionCode(context.Background(), &pb.RegisterSectionCodeRequest{
		AdminId:     adminId,
		SectionCode: req.SectionCode,
		OfficeId:    req.OfficeId,
	})

	if err == nil {
		return c.Status(200).JSON(response.SM{
			Status: mystatus.Success,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}

}

func (kseb *KSEBHandler) AdminGetComplaints(c *fiber.Ctx) error {
	status := c.Query("status", "all")       //all, opened, closed, not-opened
	attenderScope := c.Query("scope", "all") //all, me-only

	adminId, ok, err := gateway.GetAdminIdFromContextFiber(c)
	if !ok {
		return err
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
		return c.Status(200).JSON(response.GetKsebComplaints{
			Status:     mystatus.Success,
			Complaints: complaints,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}

}

func (kseb *KSEBHandler) AdminOpenComplaint(c *fiber.Ctx) error {
	complaintId, err := gateway.HandleGetUrlParamsInt32Fiber(c, "complaintId")
	if err != nil {
		return err
	}

	adminId, ok, err := gateway.GetAdminIdFromContextFiber(c)
	if !ok {
		return err
	}

	_, err = kseb.agencyAdminClient.OpenComplaint(context.Background(), &pb.OpenComplaintRequest{
		AdminId:     adminId,
		ComplaintId: complaintId,
	})

	if err == nil {
		return c.Status(200).JSON(response.SM{
			Status: mystatus.Success,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (kseb *KSEBHandler) AdminCloseComplaint(c *fiber.Ctx) error {
	var req requests.KsebCloseComplaint
	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	adminId, ok, err := gateway.GetAdminIdFromContextFiber(c)
	if !ok {
		return err
	}

	_, err = kseb.agencyAdminClient.CloseComplaint(context.Background(), &pb.CloseComplaintRequest{
		AdminId:     adminId,
		ComplaintId: req.ComplaintId,
		Remarks:     req.Remarks,
	})

	if err == nil {
		return c.Status(200).JSON(response.SM{
			Status: mystatus.Success,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}
