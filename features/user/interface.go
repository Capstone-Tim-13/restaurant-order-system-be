package user

import (
	"capstone/features/user/dto"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Save(Newuser *User) (*User, error)
	UpdatePassword(Newuser *User, id int) (*User, error)
	FindAll() ([]User, error)
	FindByUsername(username string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindById(id int) (*User, error)
	Delete(id int) error
}

type Service interface {
	Register(ctx echo.Context, req dto.ReqUserRegister) (*User, error)
	Login(ctx echo.Context, req dto.ReqUserLogin) (*User, error)
	FindAll(ctx echo.Context) ([]User, error)
	FindById(ctx echo.Context, id int) (*User, error)
	UpdatePassword(ctx echo.Context, req dto.ReqUserUpdate, id int) (*User, error)
	Delete(ctx echo.Context, id int) error
}

type Handler interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
	Find(ctx echo.Context) error
	UpdatePassword(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
