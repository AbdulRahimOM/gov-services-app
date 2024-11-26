package commondto

type UserConsumerNumber struct {
	Id             int32  `json:"id" `
	ConsumerNumber string `json:"consumer_number" `
	NickName       string `json:"nick_name"`
}

type KsebComplaintResponse struct {
	ID             int32  `json:"id"`
	UserID         int32  `json:"userId"`
	Type           string `json:"type"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	ConsumerNumber string `json:"consumerNumber"`
	AttenderID     int32  `json:"attenderId"`
	Status         string `json:"status"`
	CreatedAt      string `json:"createdAt"`
	Remarks        string `json:"remarks,omitempty"`
	ClosedAt       string `json:"closedAt,omitempty"`
}
