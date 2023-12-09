package handler

import (
	"capstone/features/order"
	"capstone/features/order/dto"
	"capstone/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type OrderHandlerImpl struct {
	OrderService order.Service
}

func NewOrderHandler(OrderService order.Service) order.Handler {
	return &OrderHandlerImpl{OrderService: OrderService}
}

func (h *OrderHandlerImpl) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req dto.CreateOrder
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("error when parsing data"))
		}

		result, err := h.OrderService.Create(ctx, req)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("creating error"))
		}

		return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("successfully created order", result))
	}
}

func (h *OrderHandlerImpl) FindAll() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		results, err := h.OrderService.FindAll(ctx)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("orders not found"))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("orders found", results))
	}
}

func (h *OrderHandlerImpl) FindById() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid order ID"))
		}

		result, err := h.OrderService.FindById(id)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("order not found"))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("order found", result))
	}
}

func (h *OrderHandlerImpl) Delete() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid order ID"))
		}

		err = h.OrderService.Delete(id)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("order deleted successfully", nil))
	}
}

func (h *OrderHandlerImpl) UpdateStatus() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid order ID"))
		}

		var req dto.ReqUpdateStatus
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("error when parsing data"))
		}

		err = h.OrderService.UpdateStatus(id, req.Status)
		if err != nil {
			if strings.Contains(err.Error(), "validation error") {
				return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
			}

			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("updating error"))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("successfully update status order", nil))
	}
}

func (h *OrderHandlerImpl) UpdateOrderItem() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req dto.ReqUpdateOrderItem
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("error when parsing data"))
		}

		result, err := h.OrderService.UpdateOrderItems(req)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("updating error"))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("successfully update status order", result))
	}
}
