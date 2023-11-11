package routes

import (
	"capstone/features/admin/handler"
	"capstone/features/admin/repository"
	"capstone/features/admin/service"
	"os"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AdminRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	adminRepository := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepository, validate)
	adminHandler := handler.NewAdminHandler(adminService)

	Group := e.Group("admin")

	Group.POST("/register", adminHandler.Register)
	Group.POST("/login", adminHandler.Login)

	Group.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY_ADMIN"))))

	Group.GET("", adminHandler.Find)
	Group.PUT("/password/:id", adminHandler.UpdatePassword)
	Group.DELETE("/delete/:id", adminHandler.Delete)

}
