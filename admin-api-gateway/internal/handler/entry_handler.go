package handler

import (
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/request"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	"github.com/AbdulRahimOM/gov-services-app/shared/gateway"
	mystatus "github.com/AbdulRahimOM/gov-services-app/shared/std-response/my_status"
	pb "github.com/AbdulRahimOM/gov-services-app/shared/pb/generated"

	"github.com/gin-gonic/gin"
)

func (u *AccountEntryHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (u *AccountEntryHandler) AdminLoginViaPassword(c *gin.Context) {
	var req request.AdminLoginViaPassword

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}
	resp, err := u.accountsClient.AdminLoginViaPassword(c, &pb.AdminLoginViaPasswordRequest{
		Username: req.Username,
		Password:    req.Password,
	})
	if err == nil {
		c.JSON(200, response.AdminLogin{
			Status: mystatus.Success,
			AdminData: response.AdminBasicData{
				Id:          resp.AdminDetails.Id,
				FirstName:   resp.AdminDetails.FirstName,
				LastName:    resp.AdminDetails.LastName,
				PhoneNumber: resp.AdminDetails.PhoneNumber,
			},
			Token: resp.Token,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}
