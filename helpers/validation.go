package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func ValidationError(ctx echo.Context, err error) error {
	ValidationError, ok := err.(validator.ValidationErrors)
	if ok {
		messages := make([]string, 0)
		for _, e := range ValidationError {
			messages = append(messages, fmt.Sprintf("Validation error on field %s, tag %s", e.Field(), e.Tag()))
		}
		return fmt.Errorf("validation failed : %s", strings.Join(messages, ", "))
	}
	return nil
}