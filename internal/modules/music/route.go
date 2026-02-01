package music

import "github.com/gofiber/fiber/v2"

func Register(r fiber.Router) {
	r.Get("/music", List)
	r.Get("/music/:name", GetFile)
	r.Get("/music/:name/hls/*", Stream)
}
