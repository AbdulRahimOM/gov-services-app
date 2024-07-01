package stdresponse

import (
	"encoding/json"
	"fmt"

	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

// // ParseGrpcStatus parses the error message and returns the response code, error message and error if failed to parse
// func ParseGrpcStatus(err error) (string, string, error) {
// 	fmt.Println("err:2: ", err)
// 	st, ok := status.FromError(err)
// 	fmt.Println("@@2 ok: ", ok)
// 	if !ok {
// 		return "", "", fmt.Errorf("#error is not a status error")
// 	}
// 	response := &ErrResponse{}

// 	if unmarshallErr := json.Unmarshal([]byte(st.Message()), response); unmarshallErr != nil {
// 		fmt.Println("Failed to unmarshall error message: ", unmarshallErr)
// 		return "", "", unmarshallErr
// 	}

// 	return response.ResponseCode, response.Error, nil
// }

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
