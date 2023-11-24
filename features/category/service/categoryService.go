package service

import (
	"capstone/features/category"
	"capstone/features/category/dto"
	"capstone/helpers"
	conversion "capstone/helpers/conversion/category"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CategoryServiceImpl struct {
	CategoryRepository category.Repository
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository category.Repository, validate *validator.Validate) category.Service {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, Validate: validate}
}

func (s *CategoryServiceImpl) Create(ctx echo.Context, req dto.ReqCategoryCreate) (*category.Category, error) {
	// Check if the request is valid
	err := s.Validate.Struct(req)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	// Check Category already exists
	name, _ := s.CategoryRepository.FindByName(req.Name)
	if name != nil {
		return nil, fmt.Errorf("category already exists")
	}

	// Convert Request to Models
	category := conversion.CategoryCreateRequest(req)

	_ , err = s.CategoryRepository.Save(category)
	if err != nil {
		return nil, fmt.Errorf("error When to Register: %s", err.Error())
	}

	results, _ := s.CategoryRepository.FindByName(req.Name)

	return results, nil
}

func (s *CategoryServiceImpl) FindAll(ctx echo.Context) ([]category.Category, error) {
	category, err := s.CategoryRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("admin not found")
	}

	return category, nil
}

func (s *CategoryServiceImpl) FindById(ctx echo.Context, id int) (*category.Category, error) {
	category, _ := s.CategoryRepository.FindById(id)
	if category == nil {
		return nil, fmt.Errorf("category not found")
	}
	return category, nil
}

func (s *CategoryServiceImpl) Delete(ctx echo.Context, id int) error {
	// Check Category already exists
	category, _ := s.CategoryRepository.FindById(id)
	if category == nil{
		return fmt.Errorf("category Not Found")
	}

	err := s.CategoryRepository.Delete(id)
	if err != nil{
		return fmt.Errorf("error when deleting : %s", err)
	}

	return nil
}

