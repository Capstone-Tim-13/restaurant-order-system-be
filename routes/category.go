package routes

import (
	"capstone/features/category/handler"
	"capstone/features/category/repository"
	"capstone/features/category/service"
	"os"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CategoryRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository, validate)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	Group := e.Group("category")

	Group.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY_ADMIN"))))
	Group.POST("/create", categoryHandler.Create)
	Group.GET("", categoryHandler.Find)
	Group.GET("/:id", categoryHandler.FindById)
	Group.DELETE("/delete/:id", categoryHandler.Delete)

}
