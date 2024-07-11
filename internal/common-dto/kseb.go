package commondto

type UserConsumerNumber struct {
	Id 		   int32  `json:"id" `
	ConsumerNumber string `json:"consumer_number" `
	NickName       string `json:"nick_name"`
}