package admin

import (
	"capstone/features/admin/dto"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Save(Newadmin *Admin) (*Admin, error)
	UpdatePassword(Newadmin *Admin, id int) (*Admin, error)
	FindAll() ([]Admin, error)
	FindByUsername(username string) (*Admin, error)
	FindByEmail(email string) (*Admin, error)
	FindById(id int) (*Admin, error)
	Delete(id int) error
}

type Service interface {
	Register(ctx echo.Context, req dto.ReqAdminRegister) (*Admin, error)
	Login(ctx echo.Context, req dto.ReqAdminLogin) (*Admin, error)
	FindAll(ctx echo.Context) ([]Admin, error)
	FindById(ctx echo.Context, id int) (*Admin, error)
	UpdatePassword(ctx echo.Context, req dto.ReqAdminUpdate, id int) (*Admin, error)
	Delete(ctx echo.Context, id int) error
}

type Handler interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
	Find(ctx echo.Context) error
	UpdatePassword(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
