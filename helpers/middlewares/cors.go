package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupCORS() echo.MiddlewareFunc {
	config := middleware.CORSConfig{
		AllowOrigins: []string{"https://altaresto-staging.vercel.app"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}
	return middleware.CORSWithConfig(config)
}