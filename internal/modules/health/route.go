package health

import "github.com/gofiber/fiber/v2"

func Register(r fiber.Router) {
	r.Get("/health", Check)
}
