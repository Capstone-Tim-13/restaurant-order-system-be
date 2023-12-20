package category

import (
	"capstone/features/category/dto"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Save(Newcategory *Category) (*Category, error)
	FindAll() ([]Category, error)
	FindById(id int) (*Category, error)
	FindByName(Name string) (*Category, error)
	Delete(id int) error
}

type Service interface {
	Create(ctx echo.Context, req dto.ReqCategoryCreate) (*dto.ResCategoryCreate, error)
	FindAll(ctx echo.Context) ([]Category, error)
	FindById(ctx echo.Context, id int) (*Category, error)
	Delete(ctx echo.Context, id int) error
}

type Handler interface {
	Create(ctx echo.Context) error
	Find(ctx echo.Context) error
	FindById(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
