package service

import (
	"capstone/features/feedback"
	"capstone/features/feedback/dto"
	"capstone/helpers"
	conversion "capstone/helpers/conversion/feedback"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type FeedbackServiceImpl struct {
	FeedbackRepository feedback.Repository
	Validate           *validator.Validate
}

func NewFeedbackService(feedbackRepository feedback.Repository, validate *validator.Validate) feedback.Service {
	return &FeedbackServiceImpl{FeedbackRepository: feedbackRepository, Validate: validate}
}

func (s *FeedbackServiceImpl) Create(ctx echo.Context, req dto.CreateFeedbackRequest) (*dto.CreateFeedbackResponse, error) {
	// Check if the request is valid
	err := s.Validate.Struct(req)
	if err != nil {
		fmt.Println("validation error", err.Error())
		return nil, helpers.ValidationError(ctx, err)
	}

	//convert request to models
	feedback := conversion.FeedbackCreateRequest(req)

	result, err := s.FeedbackRepository.Save(feedback)
	if err != nil {
		return nil, fmt.Errorf("Error when create : %s", err.Error())
	}

	response := conversion.CreateFeedbackResponse(result)

	return &response, nil
}

func (s *FeedbackServiceImpl) Update(ctx echo.Context, req dto.UpdateFeedbackRequest, id int) (*dto.UpdateFeedbackResponse, error) {
	// Validate the request
	err := s.Validate.Struct(req)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	// Check if the feedback exists
	existingFeedback, _ := s.FeedbackRepository.FindById(id)
	if existingFeedback == nil {
		return nil, fmt.Errorf("Feedback not found")
	}

	// Convert request to models
	feedback := conversion.FeedbackUpdateRequest(req)

	result, err := s.FeedbackRepository.Update(feedback, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating : %s", err.Error())
	}

	response := conversion.UpdateFeedbackResponse(result)
	
	return &response, nil
}

func (s *FeedbackServiceImpl) FindById(ctx echo.Context, id int) (*dto.CreateFeedbackResponse, error) {
	feedback, err := s.FeedbackRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("feedback not found")
	}

	response := conversion.CreateFeedbackResponse(feedback)

	return &response, nil
}

func (s *FeedbackServiceImpl) FindAll(ctx echo.Context) ([]dto.CreateFeedbackResponse, error) {
	feedback, err := s.FeedbackRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("feedback not found")
	}

	var responses []dto.CreateFeedbackResponse

	for _, data := range feedback {
		result := conversion.CreateFeedbackResponse(&data)
		responses = append(responses, result)
	}

	return responses, nil
}

func (s *FeedbackServiceImpl) Delete(ctx echo.Context, id int) error {
	feedback, _ := s.FeedbackRepository.FindById(id)
	if feedback == nil {
		return fmt.Errorf("feedback Not Found")
	}

	err := s.FeedbackRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting : %s", err)
	}

	return nil
}