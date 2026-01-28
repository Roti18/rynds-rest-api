package app

import "github.com/gofiber/fiber/v2"

func New() *fiber.App {
	app := fiber.New()

	RegisterRoutes(app)
	// Serve Shared Assets
	app.Static("/assets", "./docs/assets")

	// Serve Data Files (JSON, etc)
	app.Static("/data", "./docs/data")

	// Serve Music Files
	app.Static("/music", "./music")

	// Serve Documentation at /docs
	app.Static("/docs", "./docs/api")

	// Serve Landing Page at Root (Must be last to avoid overshadowing)
	app.Static("/", "./docs/landing")

	return app
}
