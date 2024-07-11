package response

type GetConnections struct {
	Status string `json:"status"`
	Connections []Connection `json:"connections"`
}

type Connection struct {
	Id int32 `json:"id"`
	ConsumerNumber string `json:"consumer_number"`
	NickName string `json:"name"`
}