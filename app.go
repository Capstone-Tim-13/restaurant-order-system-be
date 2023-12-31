package main

import (
	"capstone/config"
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
	config := config.LoadDBConfig()
	DB := database.InitDB()
	cdn := database.CloudinaryInstance(config)
	snapClient := database.MidtransSnapClient(config)
	coreAPIClient := database.MidtransCoreAPIClient(config)

		// Routes
		app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome To Alta-Resto")
	})

	// Register routes with the app
	routes.AdminRoutes(app, DB, validate)
	routes.UserRoutes(app, DB, validate)
	routes.CategoryRoutes(app, DB, validate)
	routes.MenuRoutes(app, DB, validate, cdn, config)
	routes.OrderRoutes(app, DB, validate, config)
	routes.PaymentRoutes(app, DB, snapClient, coreAPIClient, validate, config)
	routes.FeedbackRoutes(app, DB, validate)

	// Middleware
	app.Use(middleware.CORS())
	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))
	middlewares.SetupCORS(app)

	// Start the server
	err := app.Start(":80")
	//err := app.Start(":8080")
	if err != nil {
		app.Logger.Fatal(err)
	}
}
