package dto

type CreateFeedbackRequest struct {
	OrderID uint `json:"order_id" form:"order_id" validate:"required,min=1"`
	UserID  uint `json:"user_id" form:"user_id" validate:"required,min=1"`
	Rating  int  `json:"rating" form:"rating" validate:"required,gte=1,lte=5"`
	Review  string `json:"review" form:"review" validate:"required,min=1"`
}

type UpdateFeedbackRequest struct {
	Rating  int  `json:"rating" form:"rating" validate:"required,min=1,max=5"`
	Review  string `json:"review" form:"review" validate:"required,min=1"`
}

