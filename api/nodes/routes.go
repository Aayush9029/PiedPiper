package nodes

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(route fiber.Router) {
	route.Get("/nodes/:key", Get)
	route.Post("/nodes/:key", AddUpdate)
}
