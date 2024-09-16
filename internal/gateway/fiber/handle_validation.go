package gateway

import (
	"net/http"

	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	respCode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"github.com/AbdulRahimOM/gov-services-app/internal/validation"
	"github.com/gofiber/fiber/v2"
)

func BindAndValidateRequestFiber(c *fiber.Ctx, req interface{}) (bool, error) {
	gatewayLogger.WithField("method", "BindAndValidateRequestFiber")
	if err := c.BodyParser(req); err != nil {
		gatewayLogger.Info("error in parsing request to req struct")
		return false, c.Status(http.StatusBadRequest).JSON(stdresponse.SRE{
			Status:       mystatus.ValidationError,
			ResponseCode: respCode.BindingError,
			Error:        err.Error(),
		})
	}

	if err := validation.ValidateRequestDetailed(req); err != nil {
		gatewayLogger.Info("error in validating request. Error: ", err)

		return false, c.Status(http.StatusBadRequest).JSON(stdresponse.SMValidationErrors{
			Status:       mystatus.ValidationError,
			ResponseCode: respCode.ValidationError,
			Errors:       err,
		})
	}
	return true, nil
}
