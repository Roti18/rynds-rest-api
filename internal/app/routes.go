package app

import (
	"github.com/gofiber/fiber/v2"
	"rynds-api/internal/modules/user"
	"rynds-api/internal/modules/auth"
	"rynds-api/internal/modules/health"
)

func RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	user.Register(v1)
	auth.Register(v1)
	health.Register(v1)
}
