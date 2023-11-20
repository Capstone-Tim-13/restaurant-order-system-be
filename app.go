package main

import (
	"capstone/database"
	"capstone/helpers/middlewares"
	"capstone/routes"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	validate := validator.New()
	DB := database.InitDB()

	// Routes
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome To Alta-Resto")
	})

	// Register routes with the app
	routes.AdminRoutes(app, DB, validate)
	routes.UserRoutes(app, DB, validate)

	// Middleware
	app.Use(middlewares.SetupCORS())
	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	// Start the server
	err := app.Start(":8080")
	if err != nil {
		app.Logger.Fatal(err)
	}
}
