package service

import (
	"capstone/features/menu"
	"capstone/features/menu/dto"
	"capstone/helpers"
	conversion "capstone/helpers/conversion/Menu"
	"fmt"
	"mime/multipart"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type MenuServiceImpl struct {
	MenuRepository menu.Repository
	Validate       *validator.Validate
}

func NewMenuService(menuRepository menu.Repository, validate *validator.Validate) menu.Service {
	return &MenuServiceImpl{MenuRepository: menuRepository, Validate: validate}
}

func (s *MenuServiceImpl) Create(ctx echo.Context, fileHeader *multipart.FileHeader, req dto.ReqMenuCreate) (*dto.ResMenuCreate, error) {
	// Check if the request is valid
	err := s.Validate.Struct(req)
	if err != nil {
		fmt.Println("validation error", err.Error())
		return nil, helpers.ValidationError(ctx, err)
	}

	// Check Menu already exists
	menu, _ := s.MenuRepository.FindByName(req.Name)
	if menu != nil {
		return nil, fmt.Errorf("menu already exists")
	}

	// Open the file from the file header
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("error opening file header", err.Error())
		return nil, fmt.Errorf("error opening file: %s", err.Error())
	}
	defer file.Close()

	// Upload the image and get the URL
	urlImage, err := s.MenuRepository.UploadImage(ctx.Request().Context(), file, req.Name)
	if err != nil {
		fmt.Println("error upload", err.Error())
		return nil, fmt.Errorf("upload image failed: %s", err.Error())
	}

	// Update the request with the image URL
	req.Image = urlImage

	menus := conversion.MenuCreateRequest(req)

	// Save the new menu to the database
	result, err := s.MenuRepository.Save(menus)
	if err != nil {
		fmt.Println("error register")
		return nil, fmt.Errorf("error when registering menu: %s", err.Error())
	}

	response := conversion.MenuCreateResponse(result)

	return &response, nil
}

func (s *MenuServiceImpl) Update(ctx echo.Context, id int, fileHeader *multipart.FileHeader, req dto.ReqMenuUpdate) (*dto.ResMenuUpdate, error) {
	// Validate the request
	err := s.Validate.Struct(req)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	// Check if the menu exists
	existingMenu, err := s.MenuRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("error when checking menu existence: %s", err.Error())
	}
	if existingMenu == nil {
		return nil, fmt.Errorf("menu not found")
	}

	// Open the file from the file header if provided
	var file multipart.File
	if fileHeader != nil {
		file, err = fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("error opening file: %s", err.Error())
		}
		defer file.Close()

		// Upload the image and get the URL
		urlImage, err := s.MenuRepository.UploadImage(ctx.Request().Context(), file, req.Name)
		if err != nil {
			return nil, fmt.Errorf("upload image failed: %s", err.Error())
		}
		// Update the request with the new image URL
		req.Image = urlImage
	} else {
		// If no new file is provided, retain the existing image URL
		req.Image = existingMenu.Image
	}

	// Update the existing menu with the new data
	// existingMenu.Name = req.Name
	// existingMenu.CategoryID = req.CategoryID
	// existingMenu.Description = req.Description
	// existingMenu.Price = req.Price

	updateData := conversion.MenuUpdateRequest(req)
	updateData.ID = uint(id)
	// Save the updated menu to the database
	result, err := s.MenuRepository.Update(updateData)
	if err != nil {
		return nil, fmt.Errorf("error when updating menu: %s", err.Error())
	}

	response := conversion.MenuUpdateResponse(result)

	return &response, nil
}

func (s *MenuServiceImpl) FindAll(ctx echo.Context) ([]dto.ResMenuCreate, error) {
	menus, err := s.MenuRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("menu not found")
	}

	var responses []dto.ResMenuCreate

	for _, data := range menus {
		result := conversion.MenuCreateResponse(&data)
		responses = append(responses, result)
	}

	return responses, nil
}

func (s *MenuServiceImpl) FindById(ctx echo.Context, id int) (*dto.ResMenuCreate, error) {
	menus, err := s.MenuRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("menu not found")
	}

	response := conversion.MenuCreateResponse(menus)

	return &response, nil
}

func (s *MenuServiceImpl) FindByName(ctx echo.Context, name string) (*dto.ResMenuCreate, error) {
	menus, err := s.MenuRepository.FindByName(name)
	if err != nil {
		return nil, fmt.Errorf("menu not found")
	}

	response := conversion.MenuCreateResponse(menus)

	return &response, nil
}

func (s *MenuServiceImpl) FindByCategoryId(ctx echo.Context, categoryId int) ([]dto.ResMenuCreate, error) {
	menus, err := s.MenuRepository.FindByCategoryId(categoryId)
	if err != nil {
		return nil, fmt.Errorf("menu not found")
	}

	var responses []dto.ResMenuCreate

	for _, data := range menus {
		result := conversion.MenuCreateResponse(&data)
		responses = append(responses, result)
	}

	return responses, nil
}

func (s *MenuServiceImpl) Delete(ctx echo.Context, id int) error {
	// Check Admin already exists
	menus, _ := s.MenuRepository.FindById(id)
	if menus == nil {
		return fmt.Errorf("menu Not Found")
	}

	err := s.MenuRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting : %s", err)
	}

	return nil
}
