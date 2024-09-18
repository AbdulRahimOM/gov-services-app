package ksebhandler

import (
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway/fiber"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/response"
	"github.com/eapache/go-resiliency/breaker"
	"github.com/gofiber/fiber/v2"
)

type KsebHandler struct {
	agencyUserClient pb.KSEBAgencyUserServiceClient
	ksebChatClient   pb.KsebChatServiceClient
	circuitBreaker   *breaker.Breaker
}

func NewKsebHandler(
	ksebClient pb.KSEBAgencyUserServiceClient,
	chatClient pb.KsebChatServiceClient,
	circuitBreaker *breaker.Breaker,
) *KsebHandler {
	return &KsebHandler{
		agencyUserClient: ksebClient,
		ksebChatClient:   chatClient,
		circuitBreaker:   circuitBreaker,
	}
}

func (k *KsebHandler) AddConsumerNumber(c *fiber.Ctx) error {
	var req requests.UserAddConsumerNumber

	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	userID, ok, err := gateway.GetUserIdFromContextFiber(c)
	if !ok {
		return err
	}

	_, err = k.agencyUserClient.AddConsumerNumber(c.Context(), &pb.AddConsumerNumberRequest{
		UserId:         userID,
		ConsumerNumber: req.ConsumerNumber,
		NickName:       req.NickName,
	})
	if err == nil {
		return c.Status(200).JSON(response.SM{
			Status: mystatus.Success,
			Msg:    "Consumer number added successfully",
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (k *KsebHandler) GetUserConsumerNumbers(c *fiber.Ctx) error {
	userID, ok, err := gateway.GetUserIdFromContextFiber(c)
	if !ok {
		return err
	}

	resp, err := k.agencyUserClient.GetUserConsumerNumbers(c.Context(), &pb.GetUserConsumerNumbersRequest{
		UserId: userID,
	})
	if err == nil {
		connections := make([]response.Connection, len(resp.ConsumerNumbers))
		for i, consumerNumber := range resp.ConsumerNumbers {
			connections[i] = response.Connection{
				Id:             consumerNumber.Id,
				ConsumerNumber: consumerNumber.ConsumerNumber,
				NickName:       consumerNumber.NickName,
			}
		}
		return c.Status(200).JSON(response.GetConnections{
			Status:      mystatus.Success,
			Connections: connections,
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}

func (k *KsebHandler) RaiseComplaint(c *fiber.Ctx) error {
	var req requests.KSEBComplaint

	if ok, err := gateway.BindAndValidateRequestFiber(c, &req); !ok {
		return err
	}

	userID, ok, err := gateway.GetUserIdFromContextFiber(c)
	if !ok {
		return err
	}

	resp, err := k.agencyUserClient.RaiseComplaint(c.Context(), &pb.RaiseComplaintRequest{
		UserId: userID,
		Complaint: &pb.Complaint{
			Type:           req.Type,
			Category:       req.Category,
			Title:          req.Title,
			Description:    req.Description,
			ConsumerNumber: req.ConsumerNumber,
		},
	})
	if err == nil {
		return c.Status(200).JSON(response.KSEB_RaiseComplaint{
			Status: mystatus.Success,
			ComplaintDetails: response.ComplaintDetails{
				Id: resp.ComplaintId,
			},
		})
	} else {
		return gateway.HandleGrpcStatusFiber(c, err)
	}
}
