package gateway

import (
	"log"
	"net/http"

	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	respCode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"github.com/AbdulRahimOM/gov-services-app/internal/validation"
	"github.com/gofiber/fiber/v2"
)

func BindAndValidateRequestFiber(c *fiber.Ctx, req interface{}) error {
	if err := c.BodyParser(req); err != nil {
		log.Println("error in parsing request to req struct: ", err)
		return c.Status(http.StatusBadRequest).JSON(stdresponse.SRE{
			Status:       mystatus.ValidationError,
			ResponseCode: respCode.BindingError,
			Error:        err.Error(),
		})
	}

	if err := validation.ValidateRequestDetailed(req); err != nil {
		log.Println("error in validating request: ", err)
		return c.Status(http.StatusBadRequest).JSON(stdresponse.SMValidationErrors{
			Status:       mystatus.ValidationError,
			ResponseCode: respCode.ValidationError,
			Errors:       err,
		})
	}
	return nil
}
