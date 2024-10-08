package gateway

import (
	statuscode "github.com/AbdulRahimOM/go-utils/statuscode"
	mystatus "github.com/AbdulRahimOM/gov-services-app/internal/std-response/my_status"
	respCode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/internal/std-response/std-response"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/status"
)

func HandleGrpcStatusFiber(c *fiber.Ctx, err error) error {
	gatewayLogger.WithField("method", "HandleGrpcStatusFiber")
	s, ok := status.FromError(err)
	if ok {
		gatewayLogger.Info("grpc error: ", s.Message(), " response code: ", s.Code())
		responseCode, errorMsg, _ := stdresponse.ParseGrpcStatus(s)
		return c.Status(statuscode.ConvertGrpcToHTTP(s.Code())).JSON(stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: responseCode,
			Error:        errorMsg,
		})
	} else {
		gatewayLogger.Error("ok is false while parsing grpc error. Error: ", err)
		return c.Status(500).JSON(stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respCode.GrpcCommunicationError,
			Error:        "Err communicating with accounts service: " + err.Error(),
		})
	}
}
