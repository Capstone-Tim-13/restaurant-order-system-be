package dto

type CreateFeedbackRequest struct {
	OrderID uint
	UserID  uint
	Rating  int
	Review  string 
}