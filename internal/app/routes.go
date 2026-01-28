package app

import (
	"rynds-api/internal/modules/auth"
	"rynds-api/internal/modules/health"
	"rynds-api/internal/modules/music"
	"rynds-api/internal/modules/user"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	// Register routes directly without /v1 prefix
	user.Register(app)
	auth.Register(app)
	health.Register(app)
	music.Register(app)
}
