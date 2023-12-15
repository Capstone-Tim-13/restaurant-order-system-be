package feedback

import (
	"capstone/features/feedback/dto"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Save(feedback *Feedback) (*Feedback, error)
	Update(feedbacks *Feedback, id int) (*Feedback, error)
	FindById(id int) (*Feedback, error)
	FindAll() ([]Feedback, error)
	Delete(id int) error
}

type Service interface {
	Create(ctx echo.Context, req dto.CreateFeedbackRequest) (*dto.CreateFeedbackResponse, error)
	Update(ctx echo.Context, req dto.UpdateFeedbackRequest, id int) (*dto.UpdateFeedbackResponse, error)
	FindById(ctx echo.Context, id int) (*dto.CreateFeedbackResponse, error)
	FindAll(ctx echo.Context) ([]dto.CreateFeedbackResponse, error)
	Delete(ctx echo.Context, id int) error
}

type Handler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	FindById() echo.HandlerFunc
	FindAll() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
