package auth

import "github.com/gofiber/fiber/v2"

func Register(r fiber.Router) {
	r.Post("/auth/login", Login)
}
