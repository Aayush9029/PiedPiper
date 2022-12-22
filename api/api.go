package api

import (
	"github.com/Aayush9029/PiedPiper/api/nodes"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")
	nodes.Routes(v1)
}
