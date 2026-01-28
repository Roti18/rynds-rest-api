package user

import "github.com/gofiber/fiber/v2"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	return c.JSON(GetAllUsers())
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"id": c.Params("id"),
	})
}

func (h *Handler) Create(c *fiber.Ctx) error {
	return c.Status(201).JSON(fiber.Map{
		"status": "created",
	})
}
