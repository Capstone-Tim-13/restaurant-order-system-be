package order

import (
	"capstone/features/order/dto"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Save(order *Order) (*Order, error)
	FindAll() ([]Order, error)
	FindById(id int) (*Order, error)
	Delete(id int) error
	FindMenu(menuID []int) (bool, []float32)
}

type Service interface {
	Create(ctx echo.Context, req dto.CreateOrder) (*Order, error)
	FindAll(ctx echo.Context) ([]Order, error)
	FindById(id int) (*Order, error)
	Delete(id int) error
}

type Handler interface {
	Create() echo.HandlerFunc
	FindAll() echo.HandlerFunc
	FindById() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
