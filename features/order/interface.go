package order

import (
	"capstone/features/order/dto"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Save(order *Order) (*Order, error)
	FindAll() ([]Order, error)
	FindById(id int) (*Order, error)
	FindOrderItemById(id int) (*OrderItem, error)
	Delete(id int) error
	FindMenu(menuID []int) (bool, []float32)
	Update(updateOrder *Order) (*Order, error)
	UpdateOrderItem(updateOrderItem *OrderItem) error
	CalculateTotalPrice(orderId int) (float32, error)
}

type Service interface {
	Create(ctx echo.Context, req dto.CreateOrder) (*dto.ResOrder, error)
	FindAll(ctx echo.Context) ([]dto.ResOrder, error)
	FindById(id int) (*dto.ResOrder, error)
	Delete(id int) error
	UpdateStatus(id int, status string) error
	UpdateOrderItems(updateOrderItem dto.ReqUpdateOrderItem) (*dto.ResOrder, error)
}

type Handler interface {
	Create() echo.HandlerFunc
	FindAll() echo.HandlerFunc
	FindById() echo.HandlerFunc
	Delete() echo.HandlerFunc
	UpdateStatus() echo.HandlerFunc
	UpdateOrderItem() echo.HandlerFunc
}
