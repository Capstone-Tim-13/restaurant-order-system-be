package handler

import (
	"capstone/features/payment"
	"capstone/features/payment/dto"
	"capstone/helpers"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentHandlerImpl struct {
	PaymentService payment.Service
}

func NewPaymentHandler(PaymentService payment.Service) payment.Handler {
	return &PaymentHandlerImpl{PaymentService: PaymentService}
}

func (h *PaymentHandlerImpl) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req dto.CreatePayment
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("error when parsing data"))
		}

		result, err := h.PaymentService.Create(ctx, req)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("creating error"))
		}

		return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("successfully created transaction", result))
	}
}

func (h *PaymentHandlerImpl) FindAll() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		results, err := h.PaymentService.FindAll(ctx)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("payment not found"))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("payment found", results))
	}
}

func (h *PaymentHandlerImpl) FindById() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid payment ID"))
		}

		result, err := h.PaymentService.FindById(id)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("payment not found"))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("payment found", result))
	}
}

func (h *PaymentHandlerImpl) Delete() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid payment ID"))
		}

		err = h.PaymentService.Delete(id)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("payment deleted successfully", nil))
	}
}

func (h *PaymentHandlerImpl) Notification() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var notificationPayload map[string]any

		if err := json.NewDecoder(ctx.Request().Body).Decode(&notificationPayload); err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("error when parshing data"))
		}

		err := h.PaymentService.Notification(notificationPayload)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, echo.Map{"status": "ok"})
	}
}
