package menu

import (
	"capstone/features/menu/dto"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Save(menu *Menu) (*Menu, error)
	Update(menu *Menu, id int) (*Menu, error)
	FindAll() ([]*Menu, error)
	FindById(id int) (*Menu, error)
	FindByName(name string) (*Menu, error)
	FindByCategoryId(categoryId int) ([]Menu, error)
	Delete(id int) error
}

type Service interface {
	Create(ctx echo.Context, request dto.ReqMenuCreate) (*Menu, error)
	Update(ctx echo.Context, request dto.ReqMenuUpdate, id int) (*Menu, error)
	FindAll(ctx echo.Context) ([]Menu, error)
	FindById(ctx echo.Context, id int) (*Menu, error)
	FindByCategoryId(ctx echo.Context, categoryId int) (*Menu, error)
	Delete(ctx echo.Context, id int) error
}

type Handler interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	FindById(ctx echo.Context) error
	FindByName(ctx echo.Context) error
	FindByCategoryId(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
