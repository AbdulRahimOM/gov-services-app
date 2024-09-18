package stdresponse

import (
	"encoding/json"
	"fmt"

	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewGetGrpcStatusForService(serviceName string, logger *logrus.Entry) func(respCode string, errMsg string) error {

	logger.WithField("service", serviceName)
	return func(respCode string, errMsg string) error {

		switch respcode.GetLogLevel(respCode) {
		case logrus.DebugLevel:
			logger.Debug(errMsg)
		case logrus.InfoLevel:
			logger.Info(errMsg)
		case logrus.WarnLevel:
			logger.Warn(errMsg)
		case logrus.ErrorLevel:
			logger.Error(errMsg)
		default:
			logger.Trace(errMsg, fmt.Sprintf("Not implemented log level handling for %v. This occured at GetGrpcStatusFunction. Need to implement", respcode.GetLogLevel(respCode)))
		}

		response := &ErrResponse{
			ResponseCode: respCode,
			Error:        errMsg,
		}
		data, err := json.Marshal(response)
		if err != nil {
			logger.Error("Failed to encode status message: ", err)
			return status.Error(respcode.GetGRPCCode(respCode), errMsg)
		}
		return status.Error(respcode.GetGRPCCode(response.ResponseCode), string(data))
	}
}
func GetGrpcStatus(respCode string, errMsg string) error {
	response := &ErrResponse{
		ResponseCode: respCode,
		Error:        errMsg,
	}

	data, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Failed to encode status message: %v", err)
		return status.Error(respcode.GetGRPCCode(respCode), errMsg)
	}
	return status.Error(respcode.GetGRPCCode(response.ResponseCode), string(data))
}


// ParseGrpcStatus parses the error message and returns the response code, error message and error if failed to parse
func ParseGrpcStatus(st *status.Status) (string, string, error) {
	response := &ErrResponse{}
	unmarshallErr := json.Unmarshal([]byte(st.Message()), response)
	if unmarshallErr == nil {
		return response.ResponseCode, response.Error, nil
	} else {
		if st.Code() == codes.Unimplemented {
			return respcode.GrpcUnimplementedHandler, st.Message(), nil
		} else {
			return respcode.UnknownErrorViaGrpc, st.Message(), unmarshallErr
		}
	}

}
