package respcode

import "google.golang.org/grpc/codes"

// KSEB response codes
const (
	//Section code related
	KSEB_SectionCodeExists      = "KSEB-ERR-001"
	KSEB_SectionOfficeNotExists = "KSEB-ERR-002"
	KSEB_SectionOfficeNotValid  = "KSEB-ERR-003"

	//Consumer number related
	KSEB_ConsumerNumberInvalid           = "KSEB-ERR-004"
	KSEB_ConsumerNumberAlreadyRegistered = "KSEB-ERR-005"
)

// KSEB response codes
func init() {
	//Section code related
	errCodeMap[KSEB_SectionCodeExists] = codes.AlreadyExists       //KSEB-ERR-001
	errCodeMap[KSEB_SectionOfficeNotExists] = codes.NotFound       //KSEB-ERR-002
	errCodeMap[KSEB_SectionOfficeNotValid] = codes.InvalidArgument //KSEB-ERR-003

	//Consumer number related
	errCodeMap[KSEB_ConsumerNumberInvalid] = codes.NotFound //KSEB-ERR-004
}
