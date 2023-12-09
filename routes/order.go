package routes

import (
	"capstone/config"
	"capstone/features/order/handler"
	"capstone/features/order/repository"
	"capstone/features/order/service"
	"os"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func OrderRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate, config config.DatabaseConfig) {
	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository, validate)
	orderHandler := handler.NewOrderHandler(orderService)

	GroupAdmin := e.Group("/admin")
	GroupAdmin.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY_ADMIN"))))
	GroupAdmin.GET("/order", orderHandler.FindAll())
	GroupAdmin.GET("/order/:id", orderHandler.FindById())
	GroupAdmin.DELETE("/order/delete/:id", orderHandler.Delete())
	GroupAdmin.PUT("/order/status/:id", orderHandler.UpdateStatus())

	GroupUser := e.Group("/user")
	GroupUser.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY_USER"))))
	GroupUser.POST("/create/order", orderHandler.Create())
	GroupUser.GET("/order", orderHandler.FindAll())
	GroupUser.GET("/order/:id", orderHandler.FindById())
	GroupUser.PUT("/order/update-item", orderHandler.UpdateOrderItem())
}
