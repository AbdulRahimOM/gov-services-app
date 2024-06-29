package respcode

import (
	"google.golang.org/grpc/codes"
)

const (
	//General
	BindingError    = "GEN-ERR-001"
	ValidationError = "GEN-ERR-002"
	CorruptRequest  = "GEN-ERR-003"

	//Account Credentials
	MobileNotRegistered         = "ACC-ERR-001"
	MobileAlreadyRegistered     = "ACC-ERR-002"
	InvalidOTP                  = "ACC-ERR-003"
	Unauthenticated             = "ACC-ERR-004"
	Unauthorized                = "ACC-ERR-005"
	InvalidPassword             = "ACC-ERR-006"
	TokenExpired                = "ACC-ERR-007"
	NotEnoughPermissionsInToken = "ACC-ERR-008"

	//Internal
	DBError                = "INT-ERR-001"
	OtherInternalError     = "INT-ERR-002"
	GrpcCommunicationError = "INT-ERR-003"
	UnknownErrorViaGrpc    = "INT-ERR-004"

	//Potential bugs (will reach this code if the code is not handled properly)
	BugNoUserInContext       = "BUG-ERR-001" //Caution: If user is not set in context by the middleware
	GrpcUnimplementedHandler = "BUG-ERR-002" //Caution: Missing handler in the grpc server
)

var errCodeMap = map[string]codes.Code{}

func init() {
	//General
	errCodeMap[BindingError] = codes.InvalidArgument
	errCodeMap[ValidationError] = codes.InvalidArgument
	errCodeMap[CorruptRequest] = codes.InvalidArgument

	//Account Credentials
	errCodeMap[MobileNotRegistered] = codes.NotFound
	errCodeMap[MobileAlreadyRegistered] = codes.AlreadyExists
	errCodeMap[InvalidOTP] = codes.Unauthenticated
	errCodeMap[Unauthenticated] = codes.Unauthenticated
	errCodeMap[Unauthorized] = codes.PermissionDenied
	errCodeMap[InvalidPassword] = codes.PermissionDenied

	//Internal
	errCodeMap[DBError] = codes.Internal
	errCodeMap[OtherInternalError] = codes.Internal
	errCodeMap[GrpcCommunicationError] = codes.Internal
	errCodeMap[UnknownErrorViaGrpc] = codes.Internal

	//Potential bugs
	errCodeMap[BugNoUserInContext] = codes.Unimplemented
	errCodeMap[GrpcUnimplementedHandler] = codes.Unimplemented
}

// GetGRPCCode returns the grpc code for the given error code string of the application
func GetGRPCCode(errCode string) codes.Code {
	if grpcCode, ok := errCodeMap[errCode]; !ok {
		return codes.Unknown
	} else {
		return grpcCode
	}
}
