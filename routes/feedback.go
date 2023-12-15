package routes

import (
	"capstone/features/feedback/handler"
	"capstone/features/feedback/repository"
	"capstone/features/feedback/service"
	"os"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FeedbackRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	feedbackRepository := repository.NewFeedbackRepository(db)
	feedbackService := service.NewFeedbackService(feedbackRepository, validate)
	feedbackHandler := handler.NewFeedbackHandler(feedbackService)

	GroupAdmin := e.Group("/admin")
	GroupAdmin.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY_ADMIN"))))
	GroupAdmin.GET("/feedback", feedbackHandler.FindAll())
	GroupAdmin.GET("/feedback/:id", feedbackHandler.FindById())
	GroupAdmin.DELETE("/feedback/delete/:id", feedbackHandler.Delete())


	GroupUser := e.Group("/user")
	GroupUser.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY_USER"))))
	GroupUser.POST("/feedback/create", feedbackHandler.Create())
	GroupUser.GET("/feedback", feedbackHandler.FindAll())
}
