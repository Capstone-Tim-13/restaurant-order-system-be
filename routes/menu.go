package routes

import (
	"capstone/config"
	"capstone/features/menu/handler"
	"capstone/features/menu/repository"
	"capstone/features/menu/service"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func MenuRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate, cdn *cloudinary.Cloudinary, config config.DatabaseConfig) {
	menuRepository := repository.NewMenuRepository(db, cdn)
	menuService := service.NewMenuService(menuRepository, validate)
	menuHandler := handler.NewMenuHandler(menuService)

	GroupAdmin := e.Group("/admin")
	GroupAdmin.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY_ADMIN"))))
	GroupAdmin.POST("/create/menu", menuHandler.Create())
	GroupAdmin.PUT("/update/menu/:id", menuHandler.Update())
	GroupAdmin.DELETE("/delete/menu/:id", menuHandler.Delete())
	GroupAdmin.GET("/menu", menuHandler.FindAll())
	GroupAdmin.GET("/menu/:id", menuHandler.FindById())
	GroupAdmin.GET("/menu/name/:name", menuHandler.FindByName())
	GroupAdmin.GET("/menu/category/:categoryid", menuHandler.FindByCategoryId())
	GroupAdmin.PUT("/status/:id", menuHandler.UpdateStatus())

	GroupUser := e.Group("/user")
	GroupUser.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY_USER"))))
	GroupUser.GET("/menu", menuHandler.FindAll())
	GroupUser.GET("/menu/:id", menuHandler.FindById())
	GroupUser.GET("/menu/name/:name", menuHandler.FindByName())
	GroupUser.GET("/menu/category/:categoryid", menuHandler.FindByCategoryId())

}
