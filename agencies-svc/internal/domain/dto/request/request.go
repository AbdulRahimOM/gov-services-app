package request

type KsebComplaintSearchCriteria struct {
	Status        string `json:"status"`
	AttenderScope string `json:"attender_scope"`
}
