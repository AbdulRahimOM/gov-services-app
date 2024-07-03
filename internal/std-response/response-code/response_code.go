package respcode

import (
	"log"

	"google.golang.org/grpc/codes"
)

const (
	//General
	BindingError       = "GEN-ERR-001"
	ValidationError    = "GEN-ERR-002"
	CorruptRequest     = "GEN-ERR-003"
	InvalidQueryParams = "GEN-ERR-004"

	//Account Credentials
	PhoneNumberNotRegistered     = "ACC-ERR-001"
	PhoneNumberAlreadyRegistered = "ACC-ERR-002"
	InvalidOTP                   = "ACC-ERR-003"
	Unauthenticated              = "ACC-ERR-004"
	Unauthorized                 = "ACC-ERR-005"
	InvalidPassword              = "ACC-ERR-006"
	TokenExpired                 = "ACC-ERR-007"
	NotEnoughPermissionsInToken  = "ACC-ERR-008"

	//Data
	AlreadyExists = "DATA-ERR-001"

	//Internal
	DBError                = "INT-ERR-001"
	OtherInternalError     = "INT-ERR-002"
	GrpcCommunicationError = "INT-ERR-003"
	UnknownErrorViaGrpc    = "INT-ERR-004"

	//Potential bugs (will reach this code if the code is not handled properly)
	BugNoUserInContext       = "BUG-ERR-001" //Caution: If user is not set in context by the middleware
	GrpcUnimplementedHandler = "BUG-ERR-002" //Caution: Missing handler in the grpc server
	BugNoAdminInContext      = "BUG-ERR-003" //Caution: If admin is not set in context by the middleware
	FailureToGenerate        = "BUG-ERR-004" //Caution: If failed to generate something
)

var errCodeMap = map[string]codes.Code{}

func init() {
	//General
	errCodeMap[BindingError] = codes.InvalidArgument    //GEN-ERR-001
	errCodeMap[ValidationError] = codes.InvalidArgument //GEN-ERR-002
	errCodeMap[CorruptRequest] = codes.InvalidArgument  //GEN-ERR-003

	//Account Credentials
	errCodeMap[PhoneNumberNotRegistered] = codes.NotFound            //ACC-ERR-001
	errCodeMap[PhoneNumberAlreadyRegistered] = codes.AlreadyExists   //ACC-ERR-002
	errCodeMap[InvalidOTP] = codes.Unauthenticated                   //ACC-ERR-003
	errCodeMap[Unauthenticated] = codes.Unauthenticated              //ACC-ERR-004
	errCodeMap[Unauthorized] = codes.PermissionDenied                //ACC-ERR-005
	errCodeMap[InvalidPassword] = codes.PermissionDenied             //ACC-ERR-006
	errCodeMap[TokenExpired] = codes.PermissionDenied                //ACC-ERR-007
	errCodeMap[NotEnoughPermissionsInToken] = codes.PermissionDenied //ACC-ERR-008

	//Internal
	errCodeMap[DBError] = codes.Internal                //INT-ERR-001
	errCodeMap[OtherInternalError] = codes.Internal     //INT-ERR-002
	errCodeMap[GrpcCommunicationError] = codes.Internal //INT-ERR-003
	errCodeMap[UnknownErrorViaGrpc] = codes.Internal    //INT-ERR-004

	//Potential bugs
	errCodeMap[BugNoUserInContext] = codes.Unimplemented       //BUG-ERR-001
	errCodeMap[GrpcUnimplementedHandler] = codes.Unimplemented //BUG-ERR-002
	errCodeMap[BugNoAdminInContext] = codes.Unimplemented      //BUG-ERR-003
}

// GetGRPCCode returns the grpc code for the given error code string of the application
func GetGRPCCode(errCode string) codes.Code {
	if grpcCode, ok := errCodeMap[errCode]; !ok {
		log.Printf("Unimplemented/Unknown error code: %s", errCode)
		return codes.Unknown
	} else {
		return grpcCode
	}
}
