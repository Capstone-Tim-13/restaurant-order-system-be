package handler

import (
	"capstone/features/menu"
	"capstone/features/menu/dto"
	"capstone/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type MenuHandlerImpl struct {
	MenuService menu.Service
}

func NewMenuHandler(MenuService menu.Service) menu.Handler {
	return &MenuHandlerImpl{MenuService: MenuService}
}

func (h *MenuHandlerImpl) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fileHeader, err := ctx.FormFile("image")
		if err != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, helpers.ErrorResponse("image not found"))
		}

		var req dto.ReqMenuCreate
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("error when parsing data"))
		}

		result, err := h.MenuService.Create(ctx, fileHeader, req)
		if err != nil {
			switch {
			case strings.Contains(err.Error(), "validation error"):
				return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
			case strings.Contains(err.Error(), "menu already exists"):
				return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("menu already exists"))
			default:
				return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("creating error"))
			}
		}

		return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("successfully created menu", result))
	}
}

func (h *MenuHandlerImpl) Update() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid menu ID"))
		}

		fileHeader, err := ctx.FormFile("image")
		if err != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, helpers.ErrorResponse("image not found"))
		}

		var req dto.ReqMenuUpdate
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("error when parsing data"))
		}

		result, err := h.MenuService.Update(ctx, id, fileHeader, req)
		if err != nil {
			if strings.Contains(err.Error(), "validation error") {
				return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
			}

			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("updating error"))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("successfully update menu", result))
	}
}

func (h *MenuHandlerImpl) FindAll() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		results, err := h.MenuService.FindAll(ctx)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("menus not found"))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("menus found", results))
	}
}

func (h *MenuHandlerImpl) FindById() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid menu ID"))
		}

		result, err := h.MenuService.FindById(ctx, id)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("menu not found"))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("menu found", result))
	}
}

func (h *MenuHandlerImpl) FindByName() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		name := ctx.Param("name")
		if name == "" {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("missing menu name"))
		}

		result, err := h.MenuService.FindByName(ctx, name)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("menu not found"))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("menu found", result))
	}
}

func (h *MenuHandlerImpl) FindByCategoryId() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		categoryID, err := strconv.Atoi(ctx.Param("categoryid"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid category ID"))
		}

		results, err := h.MenuService.FindByCategoryId(ctx, categoryID)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("menus not found"))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("menus found", results))
	}
}

func (h *MenuHandlerImpl) Delete() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid menu ID"))
		}

		err = h.MenuService.Delete(ctx, id)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("menu deleted successfully", nil))
	}
}
