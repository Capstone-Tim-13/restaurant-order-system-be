package conversion

import (
	"capstone/features/feedback"
	"capstone/features/feedback/dto"
)

func FeedbackCreateRequest(req dto.CreateFeedbackRequest) *feedback.Feedback {
	return &feedback.Feedback{
		UserID:  req.UserID,
		OrderID: req.OrderID,
		Rating:  req.Rating,
		Review:  req.Review,
	}
}

func FeedbackUpdateRequest(req dto.UpdateFeedbackRequest) *feedback.Feedback {
	return &feedback.Feedback{
		Rating: req.Rating,
		Review: req.Review,
	}
}
