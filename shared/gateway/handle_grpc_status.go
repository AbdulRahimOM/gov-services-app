package gateway

import (
	"log"

	statuscode "github.com/AbdulRahimOM/go-utils/statuscode"
	mystatus "github.com/AbdulRahimOM/gov-services-app/shared/std-response/my_status"
	respCode "github.com/AbdulRahimOM/gov-services-app/shared/std-response/response-code"
	stdresponse "github.com/AbdulRahimOM/gov-services-app/shared/std-response/std-response"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

func HandleGrpcStatus(c *gin.Context, err error) {
	s, ok := status.FromError(err)
	if ok {
		responseCode, errorMsg, _ := stdresponse.ParseGrpcStatus(s)
		c.JSON(statuscode.ConvertGrpcToHTTP(s.Code()), stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: responseCode,
			Error:        errorMsg,
		})
		return
	} else {
		log.Println("ok is false while parsing grpc error. Error: ", err)
		c.JSON(500, stdresponse.SRE{
			Status:       mystatus.Failed,
			ResponseCode: respCode.GrpcCommunicationError,
			Error:        "Err communicating with accounts service: " + err.Error(),
		})
		return
	}
}
