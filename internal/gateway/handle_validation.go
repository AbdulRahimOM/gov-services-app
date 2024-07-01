package gateway

import (
	"log"
	"net/http"

	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	respCode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"github.com/AbdulRahimOM/gov-services-app/internal/validation"
	"github.com/gin-gonic/gin"
)

func BindAndValidateRequest(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		log.Println("error in binding request: ", err)
		c.JSON(http.StatusBadRequest, stdresponse.SRE{
			Status:       mystatus.ValidationError,
			ResponseCode: respCode.BindingError,
			Error:        err.Error(),
		})
		return false
	}

	if err := validation.ValidateRequestDetailed(req); err != nil {
		log.Println("error in validating request: ", err)
		c.JSON(http.StatusBadRequest, stdresponse.SMValidationErrors{
			Status:       mystatus.ValidationError,
			ResponseCode: respCode.ValidationError,
			Errors:       err,
		})
		return false
	}
	return true
}
