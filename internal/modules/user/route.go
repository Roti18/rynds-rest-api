package user

import "github.com/gofiber/fiber/v2"

func Register(r fiber.Router) {
	g := r.Group("/users")
	h := NewHandler()

	g.Get("/", h.GetAll)
	g.Get("/:id", h.GetByID)
	g.Post("/", h.Create)
}
