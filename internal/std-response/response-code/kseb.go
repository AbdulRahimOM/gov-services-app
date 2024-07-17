package respcode

import "google.golang.org/grpc/codes"

// KSEB response codes
const (
	//Section code related
	KSEB_SectionCodeExists      = "KSEB-SCR-001"
	KSEB_SectionOfficeNotExists = "KSEB-SCR-002"
	KSEB_SectionOfficeNotValid  = "KSEB-SCR-003"

	//Consumer number related
	KSEB_ConsumerNumberInvalid           = "KSEB-CNR-001"
	KSEB_ConsumerNumberAlreadyRegistered = "KSEB-CNR-002"

	//Complaint related
	KSEB_ComplaintNotBelongsToUser     = "KSEB-CRE-001"
	KSEB_ComplaintAlreadyOpened        = "KSEB-CRE-002"
	KSEB_ComplaintNotOpened            = "KSEB-CRE-003"
	KSEB_ComplaintNotAccessibleToAdmin = "KSEB-CRE-004"
)

// KSEB response codes
func init() {
	//Section code related
	errCodeMap[KSEB_SectionCodeExists] = codes.AlreadyExists       //KSEB-SCR-001
	errCodeMap[KSEB_SectionOfficeNotExists] = codes.NotFound       //KSEB-SCR-002
	errCodeMap[KSEB_SectionOfficeNotValid] = codes.InvalidArgument //KSEB-SCR-003

	//Consumer number related
	errCodeMap[KSEB_ConsumerNumberInvalid] = codes.NotFound                //KSEB-CNR-004
	errCodeMap[KSEB_ConsumerNumberAlreadyRegistered] = codes.AlreadyExists //KSEB-CNR-005

	//Complaint related
	errCodeMap[KSEB_ComplaintNotBelongsToUser] = codes.PermissionDenied     //KSEB-CRE-001
	errCodeMap[KSEB_ComplaintAlreadyOpened] = codes.FailedPrecondition      //KSEB-CRE-002
	errCodeMap[KSEB_ComplaintNotOpened] = codes.FailedPrecondition          //KSEB-CRE-003
	errCodeMap[KSEB_ComplaintNotAccessibleToAdmin] = codes.PermissionDenied //KSEB-CRE-004
}
