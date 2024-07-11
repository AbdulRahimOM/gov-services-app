package request

import requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"

type AppointAttender struct {
	Appointer Appointer `json:"appointer" binding:"required"`
	Appointee Appointee `json:"appointee" binding:"required"`
}
type Appointer struct {
	Id int32 `json:"id" binding:"required" validate:"min=1,number"`
}

type Appointee requests.Appointee