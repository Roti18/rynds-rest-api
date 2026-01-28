package music

import "github.com/gofiber/fiber/v2"

func Register(r fiber.Router) {
	r.Get("/music", List)
}
