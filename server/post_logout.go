package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) Logout(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	cookie.Name = "authorization"
	cookie.HTTPOnly = true
	cookie.Value = "deleted"
	cookie.Expires = time.Now().Add(-3 * time.Second)
	cookie.SameSite = "lax"
	cookie.Domain = "localhost"

	c.Cookie(cookie)

	return c.SendStatus(fiber.StatusNoContent)
}
