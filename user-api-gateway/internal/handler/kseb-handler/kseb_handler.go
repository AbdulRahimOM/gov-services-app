package ksebhandler

import (
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway"
	ksebpb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated/ksebpb"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/models/response"
	"github.com/gin-gonic/gin"
)

type KsebHandler struct {
	ksebClient ksebpb.KSEBUserServiceClient
}

func NewKsebHandler(client ksebpb.KSEBUserServiceClient) *KsebHandler {
	return &KsebHandler{
		ksebClient: client,
	}
}

// AddConsumerNumber
func (k *KsebHandler) AddConsumerNumber(c *gin.Context) {
	var req requests.UserAddConsumerNumber

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	userID, ok := gateway.GetUserIdFromContext(c)
	if !ok {
		return
	}

	_, err := k.ksebClient.AddConsumerNumber(c, &ksebpb.AddConsumerNumberRequest{
		UserId:         userID,
		ConsumerNumber: req.ConsumerNumber,
		NickName:       req.NickName,
	})
	if err == nil {
		c.JSON(200, response.SM{
			Status: mystatus.Success,
			Msg:    "Consumer number added successfully",
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

// GetUserConsumerNumbers
func (k *KsebHandler) GetUserConsumerNumbers(c *gin.Context) {
	userID, ok := gateway.GetUserIdFromContext(c)
	if !ok {
		return
	}

	resp, err := k.ksebClient.GetUserConsumerNumbers(c, &ksebpb.GetUserConsumerNumbersRequest{
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
		c.JSON(200, response.GetConnections{
			Status:      mystatus.Success,
			Connections: connections,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

// RaiseComplaint
func (k *KsebHandler) RaiseComplaint(c *gin.Context) {
	var req requests.KSEBComplaint

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	userID, ok := gateway.GetUserIdFromContext(c)
	if !ok {
		return
	}

	resp, err := k.ksebClient.RaiseComplaint(c, &ksebpb.RaiseComplaintRequest{
		UserId:    userID,
		Complaint: &ksebpb.Complaint{
			Type:        req.Type,
			Category:    req.Category,
			Title:       req.Title,
			Description: req.Description,
		},
	})
	if err == nil {
		c.JSON(200, response.KSEB_RaiseComplaint{
			Status: mystatus.Success,
			ComplaintDetails: response.ComplaintDetails{
				Id: resp.ComplaintId,
			},
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}