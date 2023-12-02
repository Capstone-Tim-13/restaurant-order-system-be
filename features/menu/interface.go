package menu

import (
	"capstone/features/menu/dto"
	"context"
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Save(menu *Menu) (*Menu, error)
	UploadImage(ctx context.Context, file multipart.File, name string) (string, error)
	Update(menu *Menu) (*Menu, error)
	FindAll() ([]Menu, error)
	FindById(id int) (*Menu, error)
	FindByName(name string) (*Menu, error)
	FindByCategoryId(categoryId int) ([]Menu, error)
	Delete(id int) error
}

type Service interface {
	Create(ctx echo.Context, fileHeader *multipart.FileHeader, req dto.ReqMenuCreate) (*dto.ResMenuCreate, error)
	Update(ctx echo.Context, id int, fileHeader *multipart.FileHeader, req dto.ReqMenuUpdate) (*dto.ResMenuUpdate, error)
	FindAll(ctx echo.Context) ([]dto.ResMenuCreate, error)
	FindById(ctx echo.Context, id int) (*dto.ResMenuCreate, error)
	FindByName(ctx echo.Context, name string) (*dto.ResMenuCreate, error)
	FindByCategoryId(ctx echo.Context, categoryId int) ([]dto.ResMenuCreate, error)
	Delete(ctx echo.Context, id int) error
}

type Handler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	FindAll() echo.HandlerFunc
	FindById() echo.HandlerFunc
	FindByName() echo.HandlerFunc
	FindByCategoryId() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
