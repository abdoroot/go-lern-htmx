package handlers

import (
	"go-htmx/views/home"

	"github.com/gofiber/fiber/v2"
)

func HandleGetIndex(c *fiber.Ctx) error {
	err := Render(c, c.Response().BodyWriter(), home.Index("string title good"))
	return err
}
