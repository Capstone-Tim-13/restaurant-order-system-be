package conversion

import (
	"capstone/features/feedback"
	"capstone/features/feedback/dto"
)

func CreateFeedbackResponse(res *feedback.Feedback) dto.CreateFeedbackResponse {
	return dto.CreateFeedbackResponse{
		ID:      res.ID,
		UserID:  res.UserID,
		OrderID: res.OrderID,
		Rating:  res.Rating,
		Review:  res.Review,
	}
}

func UpdateFeedbackResponse(res *feedback.Feedback) dto.UpdateFeedbackResponse {
	return dto.UpdateFeedbackResponse{
		Rating:  res.Rating,
		Review:  res.Review,
	}
}
