package handler

import (
	"capstone/features/category"
	"capstone/features/category/dto"
	"capstone/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type CategoryHandlerImpl struct {
	CategoryService category.Service
}

func NewCategoryHandler(CategoryService category.Service) category.Handler {
	return &CategoryHandlerImpl{CategoryService: CategoryService}
}

func (h *CategoryHandlerImpl) Create(ctx echo.Context) error {
	category := dto.ReqCategoryCreate{}

	err := ctx.Bind(&category)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
	}

	response, err := h.CategoryService.Create(ctx, category)
	if err != nil {
		if strings.Contains(err.Error(), "invalid validation") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}

		if strings.Contains(err.Error(), "already exists") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("already exists"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Create error"))
	}

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Successfully create", response))

}

func (h *CategoryHandlerImpl) Find(ctx echo.Context) error {
	response, err := h.CategoryService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "category not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("category not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("find category failed"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Find Data Category", response))
}

func (h *CategoryHandlerImpl) FindById(ctx echo.Context) error {
	categoryId := ctx.Param("id")
	categoryIdInt, err := strconv.Atoi(categoryId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("find category failed"))
	}

	response, err := h.CategoryService.FindById(ctx, categoryIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("category not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("find category failed"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Find Data Category", response))
}

func (h *CategoryHandlerImpl) Delete(ctx echo.Context) error {
	categoryId := ctx.Param("id")
	category, err := strconv.Atoi(categoryId)
	if err != nil{
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	err = h.CategoryService.Delete(ctx, category)
	if err != nil{
		if strings.Contains(err.Error(), "category not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("category not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete category failed"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully deleted", nil))
}
