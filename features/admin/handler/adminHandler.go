package handler

import (
	"capstone/features/admin"
	"capstone/features/admin/dto"
	"capstone/helpers"
	conversion "capstone/helpers/conversion/admin"
	"capstone/helpers/middlewares"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type AdminHandlerImpl struct {
	AdminService admin.Service
}

func NewAdminHandler(AdminService admin.Service) admin.Handler {
	return &AdminHandlerImpl{AdminService: AdminService}
}

func (h *AdminHandlerImpl) Register(ctx echo.Context) error {
	admin := dto.ReqAdminRegister{}
	err := ctx.Bind(&admin)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	result, err := h.AdminService.Register(ctx, admin)
	if err != nil {
		if strings.Contains(err.Error(), "validation error") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}

		if strings.Contains(err.Error(), "email already in use") {
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("email already in use"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("sign up error"))
	}

	response := conversion.AdminRegisterResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("successfully signed up account", response))

}

func (h *AdminHandlerImpl) Login(ctx echo.Context) error {
	admin := dto.ReqAdminLogin{}

	err := ctx.Bind(&admin)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
	}

	response, err := h.AdminService.Login(ctx, admin)
	if err != nil {
		if strings.Contains(err.Error(), "invalid validation") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}

		if strings.Contains(err.Error(), "invalid email and password") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid email and password"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Sign in error"))
	}

	result := conversion.AdminLoginResponse(response)
	token, err := middlewares.GenerateTokenAdmin(response.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("generate jwt token error"))
	}

	result.AccessToken = token

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Successfully sign in", result))
}

func (h *AdminHandlerImpl) Find(ctx echo.Context) error {
	response, err := h.AdminService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("admin not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("find admin failed"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Find Data Admin", response))
}

func (h *AdminHandlerImpl) UpdatePassword(ctx echo.Context) error {
	adminId := ctx.Param("id")
	admin, err := strconv.Atoi(adminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	request := dto.ReqAdminUpdate{}
	err = ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	_, err = h.AdminService.UpdatePassword(ctx, request, admin)
	if err != nil {
		if strings.Contains(err.Error(), "valodation error") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("admin not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("update password error"))
	}

	result, err := h.AdminService.FindById(ctx, admin)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success update password", result))
}

func (h *AdminHandlerImpl) Delete(ctx echo.Context) error {
	adminId := ctx.Param("id")
	admin, err := strconv.Atoi(adminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	err = h.AdminService.Delete(ctx, admin)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("admin not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete admin failed"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully deleted", nil))
}
