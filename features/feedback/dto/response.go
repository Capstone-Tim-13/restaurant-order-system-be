package dto

type CreateFeedbackResponse struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	OrderID uint   `json:"order_id"`
	Rating  int    `json:"rating"`
	Review  string `json:"review"`
}

type UpdateFeedbackResponse struct {
	Rating  int    `json:"rating"`
	Review  string `json:"review"`
}