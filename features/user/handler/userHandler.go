package handler

import (
	"capstone/features/user"
	"capstone/features/user/dto"
	"capstone/helpers"
	conversion "capstone/helpers/conversion/user"
	"capstone/helpers/middlewares"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandlerImpl struct {
	UserService user.Service
}

func NewUserHandler(UserService user.Service) user.Handler {
	return &UserHandlerImpl{UserService: UserService}
}

func (h *UserHandlerImpl) Register(ctx echo.Context) error {
	user := dto.ReqUserRegister{}
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	result, err := h.UserService.Register(ctx, user)
	if err != nil {
		if strings.Contains(err.Error(), "validation error") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}

		if strings.Contains(err.Error(), "email already in use") {
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("email already in use"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("sign up error"))
	}

	response := conversion.UserRegisterResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("successfully signed up account", response))

}

func (h *UserHandlerImpl) Login(ctx echo.Context) error {
	user := dto.ReqUserLogin{}

	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
	}

	response, err := h.UserService.Login(ctx, user)
	if err != nil {
		if strings.Contains(err.Error(), "invalid validation") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}

		if strings.Contains(err.Error(), "invalid email and password") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid email and password"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Sign in error"))
	}
	result := conversion.UserLoginResponse(response)
	token, err := middlewares.GenerateTokenUser(response.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("generate jwt token error"))
	}

	result.AccessToken = token

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Successfully sign in", result))
}

func (h *UserHandlerImpl) Find(ctx echo.Context) error {
	response, err := h.UserService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("user not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("find user failed"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Find Data User", response))
}

func (h *UserHandlerImpl) UpdatePassword(ctx echo.Context) error {
	userId := ctx.Param("id")
	user, err := strconv.Atoi(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	request := dto.ReqUserUpdate{}
	err = ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	_, err = h.UserService.UpdatePassword(ctx, request, user)
	if err != nil {
		if strings.Contains(err.Error(), "valodation error") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}
		if strings.Contains(err.Error(), "user not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("user not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("update password error"))
	}

	result, err := h.UserService.FindById(ctx, user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success update password", result))
}

func (h *UserHandlerImpl) Delete(ctx echo.Context) error{
	userId := ctx.Param("id")
	user, err := strconv.Atoi(userId)
	if err != nil{
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	err = h.UserService.Delete(ctx, user)
	if err != nil{
		if strings.Contains(err.Error(), "user not found"){
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("user not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete user failed"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully deleted", nil))
}