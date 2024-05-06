package handlers

import (
	"io"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Render(ctx *fiber.Ctx, w io.Writer, c templ.Component) error {
	ctx.Set("Content-type", "text/html")
	return c.Render(ctx.Context(), w)
}
