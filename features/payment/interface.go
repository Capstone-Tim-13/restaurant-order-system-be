package payment

import (
	"capstone/features/payment/dto"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Save(payment *Payment) (*Payment, error)
	FindAll() ([]Payment, error)
	FindById(id int) (*Payment, error)
	FindOrderById(orderId int) (*Order, error)
	Update(updatePayment *Payment) (*Payment, error)
	Delete(id int) error
	SnapRequest(paymentId string, total int64) (string, string)
	CheckTransaction(paymentId string) (string, error)
}

type Service interface {
	Create(ctx echo.Context, newPayment dto.CreatePayment) (*dto.ResCreatePayment, error)
	FindAll(ctx echo.Context) ([]dto.ResPayment, error)
	FindById(id int) (*dto.ResPayment, error)
	Delete(id int) error
	Notification(notificationPayload map[string]any) error
}

type Handler interface {
	Create() echo.HandlerFunc
	FindAll() echo.HandlerFunc
	FindById() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Notification() echo.HandlerFunc
}
