package routes

import (
	"capstone/features/user/handler"
	"capstone/features/user/repository"
	"capstone/features/user/service"
	"os"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserservice(userRepository, validate)
	userHandler := handler.NewUserHandler(userService)

	Group := e.Group("user")

	Group.POST("/register", userHandler.Register)
	Group.POST("/login", userHandler.Login)

	Group.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY_USER"))))

	Group.GET("", userHandler.Find)
	Group.PUT("/password/:id", userHandler.UpdatePassword)
	Group.DELETE("/delete/:id", userHandler.Delete)

}
