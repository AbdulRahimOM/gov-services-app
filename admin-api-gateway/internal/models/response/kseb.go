package response

import commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"

// GetComplaints
type GetKsebComplaints struct {
	Status     string                            `json:"status"`
	Complaints []commondto.KsebComplaintResponse `json:"complaints"`
}

// OpenComplaint
type OpenKsebComplaint struct {
	Status    string                          `json:"status"`
	Complaint commondto.KsebComplaintResponse `json:"complaint"`
}
