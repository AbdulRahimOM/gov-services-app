package mystatus

const (
	Success = "success"
	Failed = "failed"
	ValidationError = "validation_error"	//includes request binding errors
	// InternalError = "internal_error"	//includes gRPC communication errors, db errors, other api errors, etc
	// Unauthorized = "unauthorized"	//login errors, auth errors, etc
)

