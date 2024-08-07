package acchandler

import (
	"context"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/request"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	"github.com/AbdulRahimOM/gov-services-app/internal/gateway"
	pb "github.com/AbdulRahimOM/gov-services-app/internal/pb/generated"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	"github.com/gin-gonic/gin"
)

// AdminGetProfile
func (u *AdminAccountHandler) AdminGetProfile(c *gin.Context) {
	adminID, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	resp, err := u.accountsClient.AdminGetProfile(context.Background(), &pb.AdminGetProfileRequest{
		AdminId: adminID,
	})
	if err == nil {
		c.JSON(200, response.AdminGetProfileResponse{
			Status: mystatus.Success,
			Profile: response.Profile{
				FirstName:   resp.FirstName,
				LastName:    resp.LastName,
				Username:    resp.Username,
				Email:       resp.Email,
				Address:     resp.Address,
				PhoneNumber: resp.PhoneNumber,
				Pincode:     resp.Pincode,
			},
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

// UpdatePassword
func (u *AdminAccountHandler) AdminUpdatePasswordUsingOldPw(c *gin.Context) {
	var req request.AdminUpdatePasswordUsingOldPw

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	adminID, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	_, err := u.accountsClient.AdminUpdatePasswordUsingOldPw(context.Background(), &pb.AdminUpdatePasswordUsingOldPwRequest{
		AdminId:     adminID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err == nil {
		c.JSON(200, response.SM{
			Status: mystatus.Success,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}

// AdminUpdateProfile
func (u *AdminAccountHandler) AdminUpdateProfile(c *gin.Context) {
	var req request.AdminUpdateProfile

	if ok := gateway.BindAndValidateRequest(c, &req); !ok {
		return
	}

	adminID, ok := gateway.GetAdminIdFromContext(c)
	if !ok {
		return
	}

	_, err := u.accountsClient.AdminUpdateProfile(context.Background(), &pb.AdminUpdateProfileRequest{
		AdminId:   adminID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Address:   req.Address,
		Pincode:   req.Pincode,
		PhoneNumber: req.PhoneNumber,
	})
	if err == nil {
		c.JSON(200, response.SM{
			Status: mystatus.Success,
		})
	} else {
		gateway.HandleGrpcStatus(c, err)
	}
}
