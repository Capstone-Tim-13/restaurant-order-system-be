package routes

import (
	"capstone/config"
	"capstone/features/payment/handler"
	"capstone/features/payment/repository"
	"capstone/features/payment/service"
	"os"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

func PaymentRoutes(e *echo.Echo, db *gorm.DB, snapClient snap.Client, coreAPIClient coreapi.Client, validate *validator.Validate, config config.DatabaseConfig) {
	paymentRepository := repository.NewPaymentRepository(db, snapClient, coreAPIClient)
	paymentService := service.NewPaymentService(paymentRepository, validate)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	e.POST("/notification", paymentHandler.Notification())

	GroupAdmin := e.Group("/admin")
	GroupAdmin.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY_ADMIN"))))
	GroupAdmin.GET("/payment", paymentHandler.FindAll())
	GroupAdmin.GET("/payment/:id", paymentHandler.FindById())
	GroupAdmin.DELETE("/payment/:id", paymentHandler.Delete())

	GroupUser := e.Group("/user")
	GroupUser.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY_USER"))))
	GroupUser.POST("/create/payment", paymentHandler.Create())
}
