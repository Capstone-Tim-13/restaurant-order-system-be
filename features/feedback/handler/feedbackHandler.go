package handler

import (
	"capstone/features/feedback"
	"capstone/features/feedback/dto"
	"capstone/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FeedbackHandlerImpl struct {
	FeedbackService feedback.Service
}

func NewFeedbackHandler(feedbackService feedback.Service) feedback.Handler {
	return &FeedbackHandlerImpl{FeedbackService: feedbackService}
}

func (h *FeedbackHandlerImpl) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var createRequest dto.CreateFeedbackRequest
		if err := c.Bind(&createRequest); err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error when parsing data"))
		}

		result, err := h.FeedbackService.Create(c, createRequest)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("creating error"))
		}

		return c.JSON(http.StatusCreated, helpers.SuccessResponse("successfully created feedback", result))
	}
}

func (h *FeedbackHandlerImpl) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid ID"))
		}

		var updateRequest dto.UpdateFeedbackRequest
		if err := c.Bind(&updateRequest); err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("update error"))
		}

		result, err := h.FeedbackService.Update(c, updateRequest, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, helpers.SuccessResponse("successfully created feedback", result))
	}
}

func (h *FeedbackHandlerImpl) FindById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid ID"))
		}

		result, err := h.FeedbackService.FindById(c, id)
		if err != nil {
			return c.JSON(http.StatusNotFound, helpers.ErrorResponse("not found"))
		}

		return c.JSON(http.StatusOK, helpers.SuccessResponse("successfully created feedback", result))
	}
}

func (h *FeedbackHandlerImpl) FindAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		results, err := h.FeedbackService.FindAll(c)
		if err != nil {
			return c.JSON(http.StatusNotFound, helpers.ErrorResponse("not found"))
		}

		return c.JSON(http.StatusOK, helpers.SuccessResponse("successfully created feedback", results))
	}
}

func (h *FeedbackHandlerImpl) Delete() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid feedback ID"))
		}

		err = h.FeedbackService.Delete(ctx, id)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, helpers.SuccessResponse("feedback deleted successfully", nil))
	}
}
