package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func HandlePostLogin(c *fiber.Ctx) error {
	var loginReq LoginRequest
	if err := c.BodyParser(&loginReq); err != nil {
		return err
	}
	return c.JSON(loginReq.Email)
}
