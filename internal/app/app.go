package app

import (
	"time"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
)

func New() *fiber.App {
	masterKey := os.Getenv("rynds_master_secret")

	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
		BodyLimit:    4 * 1024 * 1024, // 4MB limit
	})

	// Security Middlewares
	app.Use(fiberRecover.New()) // Anti-Panic (Prevents 502)
	app.Use(helmet.New())
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			// BYPASS LIMITER IF MASTER KEY MATCHES
			return c.Get("X-Master-Key") == masterKey
		},
		Max:        100,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("X-Real-IP", c.IP())
		},
	}))

	RegisterRoutes(app)
	// ... rest of the static routes
	app.Static("/assets", "./docs/assets")

	// Serve Data Files (JSON, etc)
	app.Static("/data", "./docs/data")

	// Serve Music Files (direct download)
	app.Static("/music/files", "./music")

	// Serve Documentation at /docs
	app.Static("/docs", "./docs/api")

	// Serve Landing Page at Root (Must be last to avoid overshadowing)
	app.Static("/", "./docs/landing")

	return app
}
