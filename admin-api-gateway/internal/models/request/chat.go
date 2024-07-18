package request

type KsebAdminChat struct {
	ComplaintId int32 `json:"complaint_id" binding:"required" validate:"min=1"`
	Message     string `json:"message" binding:"required" validate:"min=1,max=1000"`
}
